module database-example

go 1.23.0

toolchain go1.24.5

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.0
	// gRPC dependencies
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0
	// gRPC dependencies
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/handlers v1.5.2
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/crypto v0.40.0
)
