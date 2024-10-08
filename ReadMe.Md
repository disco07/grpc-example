# example-go-grpc-gateway

## Prerequisites
```shell
# Installe le plugin protoc pour générer du code Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Installe le plugin protoc pour générer les services gRPC en Go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Installe le plugin pour générer le code de la passerelle gRPC
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# (Optionnel) Installe le plugin pour générer la spécification OpenAPI
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

export PATH="$PATH:$(go env GOPATH)/bin"

```

## Tested with:

```shell
grpcurl -d '{"a": 1} {"a": 5} {"a": 3} {"a": 6} {"a": 2} {"a": 15}' -plaintext localhost:50051 calculator.CalculatorService/Max
```

```shell
curl -X POST \
  http://localhost:8080/v1/calculator/add \
  -d '{"a": 5, "b": 3}' \
  -H "Content-Type: application/json"
```
