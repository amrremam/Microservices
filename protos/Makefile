.PHONY: protos

protos:
	mkdir -p protos/currency
	protoc -I protos/ protos/currency.proto \
	--go_out=protos/currency \
	--go-grpc_out=protos/currency
