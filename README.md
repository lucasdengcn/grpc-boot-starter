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
|Validation|-|-|
|Error Handling| Errno naming, Translation| OK |
|Error Model| Google Richer Error model| OK |
|Load balancing|client side| OK |
|Health check|GRPC|OK|
|Retry|-|-|
|Service Config|-|-|
|Profiling|channelz|-|
|Logging|zerolog| OK |
|Streaming|-|-|
|Cache|-|-|
|Messaging|-|-|
|Distributed|-|-|
|Testing|testify|OK|
|Flow Control|-|-|
|API Versioning| PB having backward compatibility on model |-|
|DB Migration|go-migrate on schema| OK |
|Repository|-|OK|
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
