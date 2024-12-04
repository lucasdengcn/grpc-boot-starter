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

## Features

|Feature | Remark | Status |
|--------|--------|--------|
|Authentication| - |-|
|Configuration | Yaml, Viper | OK |
|Gorm|PostgreSQL|OK|
|Interceptor|-|-|
|Metrics|-|-|
|Tracing|-|-|
|Validation|-|-|
|Error Handling|-|-|
|Error Model| Google Richer Error model| OK |
|Load balancing|client side| OK |
|Health check|GRPC|OK|
|Retry|-|-|
|Service Config|-|-|
|Profiling|-|-|
|Logging|zerolog| OK |
|Streaming|-|-|
|Cache|-|-|
|Messaging|-|-|
|Distributed|-|-|
|Testing|testify|OK|
|Flow Control|-|-|
|API Versioning|-|-|
|DB Migration|schema migration| OK |
|Repository|-|OK|
|CDI| wire | OK |

## Reference

[Error model](https://google.aip.dev/193#error_model)

[Error proto](https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto)

## Health

health/v1
