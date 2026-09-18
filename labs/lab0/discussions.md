## Written Questions (Short Answers)

These questions are meant to test your understanding of the above resources. It may be helpful to complete
these questions and understand the answers before diving into the coding section of this lab.

You should write your answers in your `discussions.md` file. Keep answers brief.

1) What is the difference between an unbuffered channel and a buffered channel?

Unbuffered channels will block more sends until a receiver receives the current message whereas a buffered channel is able to store values allowing multiple items in the channel, so long as the buffer is not full.

2) Which is the default in Go? Unbuffered or buffered?

Unbuffered.

3) What does the following code do?

```go
func FunWithChannels() {
    ch := make(chan string)

    ch <- "hello world!"

    message := <-ch
    fmt.Println(message)
}
```

The code makes an unbuffered channel, sends the message "hello world" to the channel, and stops there, as the channel is waiting on a receiver for the message before continuing.

4) In the function signature of `MergeChannels` in `merge_channels.go`:
```go
// Make sure to read the comment block in merge_channels.go
// T is a generic type in this method signature: https://go.dev/tour/generics/2
func MergeChannels[T any](a <-chan T, b <-chan T, out chan<- T) {
```

What is the difference between `<-chan T`, `chan<- T` and `chan T`?

`chan T` refers to channels that are both able to send and receive, whereas `<-chan T` can only receive from and `chan<- T` can only send from.

5) What happens when you read from a closed channel? What about a `nil` channel?

You are still able to read from a closed channel, if it has a buffer then it will send the remaining values from the buffer, otherwise it will return values of zero if empty. If from a nil channel it will block the operation.

6) When does the following loop terminate?
```go
func FunReadingChannels(ch chan string) {
    for item := range ch {
        fmt.Println(item)
    }
}
```

The loop terminates when the channel ch is closed and its buffer is empty.

7) How can you determine if a `context.Context` is done or canceled?

Since the Done() method returns a closed chennel when the context is done or cancelled, you can put it in a select statement to determine that when, and then you can use the Err() method to determine if it was either canceled or done.

8) What does the following code (most likely) print in the most recent versions of Go (e.g., Go 1.23)? Why is that?
```go
for i := 1; i <= 3; i++ {
    go func() {
        time.Sleep(time.Duration(i) * time.Second)
        fmt.Printf("%d\n", i)
    }()
}
fmt.Println("all done!")
```

Most likely prints "all done!" because inside the for loop they create three separate goroutines, its most likely that the main function finishes printing "All done!" and terminates before the goroutines finish.

9) What concurrency utility might you use to "fix" **question 8**?

Adding a WaitGroup would make the main goroutine block until the rest of the goroutines finish.

10) What is the difference between a mutex (as in `sync.Mutex`) and a semaphore (as in `semaphore.Weighted`)?

Both mutex and semaphore are used to block access to a shared resource, the main difference is that mutex is intended to only allow one at a time whereas semaphore can manage accesses to more than 1.

11) What does the following code print?
```go
type Bar struct{}
type Foo struct {
	items  []string
	str    string
	num    int
	barPtr *Bar
	bar    Bar
}

func FunWithStructs() {
	var foo Foo
	fmt.Println(foo.items)
	fmt.Println(len(foo.items))
	fmt.Println(foo.items == nil)
	fmt.Println(foo.str)
	fmt.Println(foo.num)
	fmt.Println(foo.barPtr)
	fmt.Println(foo.bar)
}
```

Because the values in foo are not initialized, it would print:

```
[]
0
true

0
<nil>
{}
```

12) What does `struct{}` in the type `chan struct{}` mean? Why might you use it?

The value struct{} is a completely empty struct that occupies zero bytes of memory, you might use it for a channel when you want to communicate a signal that does not require any data.