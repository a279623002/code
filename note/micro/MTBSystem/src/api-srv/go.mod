module api-srv

go 1.16

require (
	config v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/tools v0.1.9 // indirect
	user-srv v0.0.0-00010101000000-000000000000
)

replace (
	config => ../../config
	user-srv => ../user-srv
)
