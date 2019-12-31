# Vault Generator #

## Introduction ##

A dead simple, wildly insecure application for turning arbitrary configuration data into an encrypted ansible vault. Written mainly in order to experiment with Go and modern React features.

The simple interface allows users to edit code via the Ace editor. The specified encryption key and text body are encrypted by a small microservice written in Go.

## Installation ##

TODO

## Table of Contents ##

TODO

## Software Overview ##

|Software|Version|
|--------|-------|
|Client|---------|
|[React](https://reactjs.org/docs/getting-started.html)|^16.12.0|
|[Jest](https://jestjs.io/docs/en/getting-started)|^24.9.0|
|[Axios](https://www.npmjs.com/package/axios)|^0.19.0|
|[Webpack](https://webpack.js.org/concepts/)|^4.41.5|
|[React Ace](https://github.com/securingsincity/react-ace/tree/master/docs)|^8.0.0|
|Server|---------|
[Go](https://golang.org/doc/)|1.13.5|
[Chi Router](https://github.com/go-chi/chi)|4.0.2|

## Development ##

Development is facilitated by Docker images tailored to provide a consistent, test-driven environment across platforms.

### Requirements ###

|Software|Version|
|--------|-------|
|[Docker](https://docs.docker.com/)|^19.03|
|[Docker Compose](https://docs.docker.com/compose/)|^1.25.0|

## Setup  ##

1. Start by cloning the repository:

```bash
$ git clone git@github.com:BeautifulTovarisch/vault-generator.git
```

2. Use `docker-compose` to build the development images:

```bash
$ cd vault-generator
$ docker-compose build
```

3. Run the client/server containers:

```bash
$ docker-compose up client
```

and/or

```bash
$ docker-compose up server
```

Source code is mounted inside the Docker containers. Detected changes automatically runs tests and restart the process, outputting the results to the terminal.

## Production Images ##

The Dockerfiles for the client and server contains multiple build-stages in order to reduce container size and minimize complexity.

Client:

```bash
$ docker build --rm -t <your image tag> client/
```

Server:
```bash
$ docker build --rm -t <your image tag> server/
```

The above commands will run instructions through the final build stage specified in each respective Dockerfile.

For more information on build stages, consult the Docker documentation on [multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/).

## Cleanup ##

You can teardown your environment with the following:
```bash
$ docker-compose down -v --rmi all
```

> **Note**: This command will completely remove the development images, containers, and volumes from your machine.
