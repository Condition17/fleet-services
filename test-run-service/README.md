# TestRunService Service

This is the TestRunService service

Generated with

```
micro new --namespace=go.micro --type=service test-run-service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.test-run-service
- Type: service
- Alias: test-run-service

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./test-run-service-service
```

Build a docker image
```
make docker
```