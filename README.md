Mock HTTP Server
================

This project provides an ultra-simple HTTP server written in Go
that responds to all requests with a 200 OK -- and is useful as
a mock endpoint for functional testing.

The project uses a special 2-phase build process that statically
compiles Go binaries and injects them into an empty Docker image,
shrinking the resulting image from hundreds of MB to < 5 MB.

## Running this Container

```console
$ docker run -p 8080:8080 deis/mock-http-server
```

## Building from Source

To build the image, run `make build`.

The build and runtime environments are split into two parts:

### The build environment

Based on [deis/go](https://registry.hub.docker.com/u/deis/go/),
this image installs a Go development environment and compiles a binary.

### The runtime environment

This image pulls in the standalone binary compiled in the build environment
and injects it into a minimal standalone container inherited `FROM scratch`.
