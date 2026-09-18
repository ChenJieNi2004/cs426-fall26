package lab0_test

import (
	"context"
	"testing"

	"cs426.cloud/lab0"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func chanToSlice[T any](ch chan T) []T {
	vals := make([]T, 0)
	for item := range ch {
		vals = append(vals, item)
	}
	return vals
}

type mergeFunc = func(chan string, chan string, chan string)

func runMergeTest(t *testing.T, merge mergeFunc) {
	t.Run("empty channels", func(t *testing.T) {
		a := make(chan string)
		b := make(chan string)
		out := make(chan string)
		close(a)
		close(b)

		merge(a, b, out)
		// If your lab0 hangs here, make sure you are closing your channels!
		require.Empty(t, chanToSlice(out))
	})

	// Please write your own tests
}

func TestMergeChannels(t *testing.T) {
	runMergeTest(t, func(a, b, out chan string) {
		lab0.MergeChannels(a, b, out)
	})
}

func TestMergeOrCancel(t *testing.T) {
	runMergeTest(t, func(a, b, out chan string) {
		_ = lab0.MergeChannelsOrCancel(context.Background(), a, b, out)
	})

	t.Run("already canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		a := make(chan string, 1)
		b := make(chan string, 1)
		out := make(chan string, 10)

		eg, _ := errgroup.WithContext(context.Background())
		eg.Go(func() error {
			return lab0.MergeChannelsOrCancel(ctx, a, b, out)
		})
		err := eg.Wait()
		a <- "a"
		b <- "b"

		require.Error(t, err)
		require.Equal(t, []string{}, chanToSlice(out))
	})

	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		a := make(chan string)
		b := make(chan string)
		out := make(chan string, 10)

		eg, _ := errgroup.WithContext(context.Background())
		eg.Go(func() error {
			return lab0.MergeChannelsOrCancel(ctx, a, b, out)
		})
		a <- "a"
		b <- "b"
		cancel()

		err := eg.Wait()
		require.Error(t, err)
		require.Equal(t, []string{"a", "b"}, chanToSlice(out))
	})
}

type channelFetcher struct {
	ch chan string
}

func newChannelFetcher(ch chan string) *channelFetcher {
	return &channelFetcher{ch: ch}
}

func (f *channelFetcher) Fetch() (string, bool) {
	v, ok := <-f.ch
	return v, ok
}

func TestMergeFetches(t *testing.T) {
	runMergeTest(t, func(a, b, out chan string) {
		lab0.MergeFetches(newChannelFetcher(a), newChannelFetcher(b), out)
	})
}

func TestMergeFetchesAdditional(t *testing.T) {
	t.Run("empty a channel", func(t *testing.T) {
		a := make(chan string)
		b := make(chan string, 2)
		out := make(chan string, 2)

		b <- "1"
		b <- "2"

		close(a)
		close(b)

		lab0.MergeFetches(newChannelFetcher(a), newChannelFetcher(b), out)
		require.ElementsMatch(t, []string{"1", "2"}, chanToSlice(out))
	})

	t.Run("a has less elements than b", func(t *testing.T) {
		a := make(chan string, 1)
		b := make(chan string, 4)
		out := make(chan string, 5)

		a <- "a"
		b <- "1"
		b <- "2"
		b <- "1"
		b <- "2"
		close(a)
		close(b)

		lab0.MergeFetches(newChannelFetcher(a), newChannelFetcher(b), out)
		require.ElementsMatch(t, []string{"a", "1", "2", "1", "2"}, chanToSlice(out))
	})
}
