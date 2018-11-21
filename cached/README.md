# Cached Service

This is the Cached service

Generated with

```
micro new github.com/guobin8205/services/cached --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.cached
- Type: srv
- Alias: cached

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./cached-srv
```

Build a docker image
```
make docker
```