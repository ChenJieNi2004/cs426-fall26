# Development Environment Setup

To ensure a consistent development environment and to give you first-hand experience with containers, we provide a container specification as the basis for the course development environment.

The current image uses Go 1.27.1 and includes Git, Docker CLI/buildx, kubectl, Protocol Buffers, gopls, goimports, Staticcheck, and the Go protobuf/gRPC generators.

## Prerequisites

* A Linux or macOS machine.
* Docker installed and running.
  * On macOS, install and start Docker Desktop.
  * On Linux, install and start Docker Engine.

Verify that Docker is working before continuing:

```bash
docker version
```

The command should show both **Client** and **Server** sections.

## Run the Development Container

```bash
docker run -it -v myhomedir:/home/nonroot ghcr.io/yryang/cs426-fall26-devenv:latest /bin/bash
```

Notes:

* `-it` starts the container interactively and attaches your terminal.
* `-v myhomedir:/home/nonroot` mounts the Docker volume `myhomedir` at `/home/nonroot` in the container.
* The named volume persists across container restarts and image updates.

Once the container is running, you can work inside it as you would in a regular terminal. The image is preconfigured with the tools and dependencies used by the labs.

You can verify the main tools with:

```bash
go version
git --version
kubectl version --client
docker --version
protoc --version
gopls version
staticcheck -version
```

## IMPORTANT: Back Up Your Work

Your `myhomedir` Docker volume is persistent, but it should not be your only copy of important work. Learn how to copy files between the container and host, and periodically keep a backup outside the container/volume. See the Docker documentation for `docker cp` and volumes.

## Optional

### Run in VS Code Dev Containers

See `.devcontainer.json`.

### Build Your Own Image (Below is description of how we build the docker image above)

You should use your own github id to replace my id (yryang) below.

Build the Ubuntu-based image (using the Dockerfile in this directory):

```bash
docker build -t ghcr.io/yryang/cs426-fall26-devenv:latest .
```

Or build the smaller Alpine-based image (using the Dockerfile.alpine in this directory)::

```bash
docker build -f Dockerfile.alpine -t ghcr.io/yryang/cs426-fall26-devenv:latest .
```

The first build may take a while; later builds should be faster because Docker caches unchanged layers.

After build the image, it is stored in Docker’s local image store, and you can check it using
```bash
docker images
```

We then publish the built image to GitHub Container Registry (GHCR). Note that you may not need to publish.

```bash
gh auth token | docker login ghcr.io -u yryang --password-stdin
docker push ghcr.io/yryang/cs426-fall26-devenv:latest
```

We made the GHCR package **Public** so that others can pull it without authenticating to GitHub.

Start locally built image with:

```bash
docker run -it -v myhomedir:/home/nonroot ghcr.io/yryang/cs426-fall26-devenv:latest /bin/bash
```

The Dockerfiles are provided so you can add tools or dependencies as needed.

### Using Docker from Inside the Development Container

The image includes the Docker CLI, but it does not run a Docker daemon itself. If a lab needs to control the host Docker daemon from inside the development container on Linux, start it with the Docker socket mounted:

```bash
docker run -it \
  -v myhomedir:/home/nonroot \
  -v /var/run/docker.sock:/var/run/docker.sock \
  ghcr.io/yryang/cs426-fall26-devenv:latest /bin/bash
```

Only mount the Docker socket when you need it: access to the host Docker socket effectively gives the container administrative control over Docker on the host.
