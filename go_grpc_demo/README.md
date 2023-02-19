
### add tool.go


```
package tools

import (
        _ "github.com/envoyproxy/protoc-gen-validate"
        _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
        _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
        _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
        _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
```

```
go install \    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
