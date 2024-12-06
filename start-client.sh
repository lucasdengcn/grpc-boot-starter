# for debugging

#export GRPC_GO_LOG_VERBOSITY_LEVEL=99
#export GRPC_GO_LOG_SEVERITY_LEVEL=info

unset GRPC_GO_LOG_VERBOSITY_LEVEL
unset GRPC_GO_LOG_SEVERITY_LEVEL

go run clients/main.go