# Broker-Mjz Service

This is the Broker-Mjz service

Generated with

```
micro new github.com/guobin8205/services/broker-mjz --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.broker-mjz
- Type: srv
- Alias: broker-mjz

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
./broker-mjz-srv
```

Build a docker image
```
make docker
```