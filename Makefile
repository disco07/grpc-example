protoc:
	cd proto && protoc --go_out=../generated --go_opt=paths=source_relative \
	--go-grpc_out=../generated --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../generated --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
  	--openapiv2_out ../api \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt merge_file_name=openapi.json \
    --openapiv2_opt allow_merge=true \
    --openapiv2_opt disable_default_errors=true \
    --openapiv2_opt include_package_in_tags=true \
	./**/*.proto

