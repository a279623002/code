module client

go 1.14

require (
	github.com/micro/go-micro/v2 v2.9.1
	user-srv v0.0.0-00010101000000-000000000000
)

replace user-srv => ../user-srv
