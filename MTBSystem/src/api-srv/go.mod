module api-srv

go 1.16

require (
	cinema-srv v0.0.0-00010101000000-000000000000
	cms-srv v0.0.0-00010101000000-000000000000
	comment-srv v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	film-srv v0.0.0-00010101000000-000000000000
	github.com/bytedance/sonic v1.8.8 // indirect
	github.com/gin-gonic/gin v1.9.0
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-playground/validator/v10 v10.13.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.1
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/tools v0.8.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	order-srv v0.0.0-00010101000000-000000000000
	place-srv v0.0.0-00010101000000-000000000000
	user-srv v0.0.0-00010101000000-000000000000
)

replace (
	cinema-srv => ../cinema-srv
	cms-srv => ../cms-srv
	comment-srv => ../comment-srv
	config => ../../config
	film-srv => ../film-srv
	order-srv => ../order-srv
	place-srv => ../place-srv
	user-srv => ../user-srv
	utils => ../../utils
)
