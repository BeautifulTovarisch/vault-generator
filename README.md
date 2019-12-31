# Vault Generator #

## Introduction ##

A wildly insecure, simple application for turning arbitrary configuration into an encrypted ansible vault. Written mainly in order to experiment with Go and modern React features.

The simple interface allows users to edit code via the Ace editor. The specified encryption key and text body are encrypted by a small microservice written in Go.

## Development ##

### Table of Contents ###

TODO

#### Toolchain ####

|Software|Version|
|--------|-------|
|[React](https://reactjs.org/docs/getting-started.html)|^16.12.0|
|[Jest](https://jestjs.io/docs/en/getting-started)|^24.9.0|
|[Axios](https://www.npmjs.com/package/axios)|^0.19.0|
|[Webpack](https://webpack.js.org/concepts/)|^4.41.5|
|[React Ace](https://github.com/securingsincity/react-ace/tree/master/docs)|^8.0.0|
|--------|-------|
[Go](https://golang.org/doc/)|1.13.5|
[Chi Router](https://github.com/go-chi/chi)|4.0.2|


#### Requirements ####

|Software|Version|
|--------|-------|
|[Docker](https://docs.docker.com/)|^19.03|
|[Docker Compose](https://docs.docker.com/compose/)|^1.25.0|
