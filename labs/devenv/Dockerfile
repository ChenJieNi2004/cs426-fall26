# Local development environment (2026)
FROM ubuntu:24.04

ARG USER=nonroot
ARG GO_VERSION=1.27.1
ENV HOME=/home/$USER
ENV DEBIAN_FRONTEND=noninteractive
ENV PATH="/usr/local/go/bin:${HOME}/go/bin:${PATH}"

# Install system packages.
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash \
    build-essential \
    ca-certificates \
    curl \
    docker-buildx \
    docker.io \
    git \
    protobuf-compiler \
    sudo \
    vim \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Install Go for either amd64 or arm64 hosts (including Apple Silicon via Docker).
RUN ARCH="$(dpkg --print-architecture)" \
    && case "$ARCH" in \
         amd64|arm64) GOARCH="$ARCH" ;; \
         *) echo "Unsupported architecture: $ARCH" >&2; exit 1 ;; \
       esac \
    && wget -q "https://go.dev/dl/go${GO_VERSION}.linux-${GOARCH}.tar.gz" -O /tmp/go.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && rm /tmp/go.tar.gz

# Install the current stable kubectl release for the image architecture.
RUN ARCH="$(dpkg --print-architecture)" \
    && KUBECTL_VERSION="$(curl -fsSL https://dl.k8s.io/release/stable.txt)" \
    && curl -fsSL "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/${ARCH}/kubectl" -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

# Create an unprivileged development user with passwordless sudo.
RUN useradd --create-home --groups sudo --shell /bin/bash "${USER}" \
    && echo "${USER} ALL=(ALL) NOPASSWD:ALL" > "/etc/sudoers.d/${USER}" \
    && chmod 0440 "/etc/sudoers.d/${USER}"

USER ${USER}
WORKDIR /home/${USER}

# Go development tools (Sep 2026).
RUN go install golang.org/x/tools/cmd/goimports@v0.50.0 \
    && go install golang.org/x/tools/gopls@v0.23.0 \
    && go install honnef.co/go/tools/cmd/staticcheck@v0.7.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.2 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11

# godoc was split into a deprecated module; retain it only for compatibility
# with existing course workflows that invoke the `godoc` command.
RUN go install golang.org/x/tools/cmd/godoc@v0.1.0-deprecated

# Build: docker build -t ghcr.io/yryang/cs426-fall26-devenv:latest .
# Start: docker run -it -v myhomedir:/home/nonroot ghcr.io/yryang/cs426-fall26-devenv:latest /bin/bash
CMD ["/bin/bash"]
