# Example

[Go Generated] (<https://protobuf.dev/reference/go/go-generated/>)

## Plugin install

## Command

```shell
# add schema version
sh schema.sh $msg
```

```shell
# CDI wire generation
sh wire-build.sh
```

```shell
# start demo server A
export GRPC_SERVER_PORT=50051
sh start-server.sh
```

```shell
# start demo server B
export GRPC_SERVER_PORT=50052
sh start-server.sh
```

```shell
# start client call
sh start-client.sh
```

## Tech Stack

- gRPC
- Go
- Gorm
- Zerolog
- Wire
- Viper
- Yaml
- PostgreSQL

## Features

|Feature | Remark | Status |
|--------|--------|--------|
|Authentication| JWT OAuth2 | OK |
|Configuration | Yaml, Viper | OK |
|Gorm|PostgreSQL|OK|
|Interceptor| Unray | OK |
|Metrics| OTEL | OK |
|Tracing| OTEL | OK |
|Validation| protovalidate-go | OK |
|Error Handling| Errno naming, Translation| OK |
|Error Model| Google Richer Error model| OK |
|Load balancing|client side| - |
|Health check| client & server side import health check package |OK|
|Retry|client side |-|
|Wait-for-Ready|client side | - |
|Deadline|timeout, client side, context.WithTimeout|-|
|Service Config| client side |-|
|Profiling|channelz|-|
|Logging|zerolog| OK |
|Streaming|-|-|
|Cache|-|-|
|Messaging|-|-|
|Distributed|-|-|
|Testing|testify|OK|
|Flow Control| Rate Limit, SHOULD be handled by Gateway |-|
|API Versioning| PB having backward compatibility on model |-|
|DB Migration|go-migrate on schema| OK |
|Repository| as design with entity |OK|
|CDI| wire | OK |
|AuthZ| casbin | -|

## Reference

[Error model](https://google.aip.dev/193#error_model)

[Error proto](https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto)

## Validation

[protovalidate](https://github.com/bufbuild/protovalidate/)

[protovalidate-go](https://github.com/bufbuild/protovalidate-go)

Rules are here. <https://buf.build/bufbuild/protovalidate/docs/main:buf.validate>

## Health

health/v1

## Debugging

### Logs

GRPC_GO_LOG_VERBOSITY_LEVEL=99
GRPC_GO_LOG_SEVERITY_LEVEL=info

### Channelz

[Channelz](https://grpc.io/blog/a-short-introduction-to-channelz)

[grpc-zpages](https://github.com/grpc/grpc-experiments/tree/master/gdebug)

## Tools

[revive](https://revive.run/docs)

[pre-commit](https://pre-commit.com/hooks.html)

[golangci-lint](https://golangci-lint.run/)

Golangci-lint aggregates dozens of tools with hundreds of checks. Revive is one of such tools

[gofumpt](https://github.com/mvdan/gofumpt)

[buf.build](https://buf.build/docs/)

## Git hooks

```shell
git config core.hooksPath .git-hooks 
```

## Development Env

### Go Requirements

```shell
go mod tidy
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

```

### Buf build

Buf CLI is a great drop-in replacement for protoc

```shell
# install CLI
brew install bufbuild/buf/buf

# update dependencies
buf dep update

# lint on proto files
buf lint

# generate code
buf generate
```

[plugins](https://buf.build/plugins)

### OpenAPI

[Open API Spec](https://github.com/getkin/kin-openapi)
