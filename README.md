# Example

[Go Generated] (<https://protobuf.dev/reference/go/go-generated/>)

## Plugin install

## Command

```shell
# generate from proto files
sh pb-gen.sh
```

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
sh server-start.sh
```

```shell
# start demo server B
export GRPC_SERVER_PORT=50052
sh server-start.sh
```

```shell
# start client call
sh client-call.sh
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
|Validation| manually on demand |-|
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
|Bidirection|-|-|
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

## Git hooks

```shell
git config core.hooksPath .git-hooks 
```
