module api-srv

go 1.16

require (
	cinema-srv v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	film-srv v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/tools v0.1.9 // indirect
	place-srv v0.0.0-00010101000000-000000000000
	user-srv v0.0.0-00010101000000-000000000000
)

replace (
	cinema-srv => ../cinema-srv
	config => ../../config
	film-srv => ../film-srv
	place-srv => ../place-srv
	user-srv => ../user-srv
	utils => ../../utils
)
