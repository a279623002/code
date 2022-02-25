module api-srv

go 1.14

replace domain/apid => ../domain/apid

replace user-srv => ../user-srv

require (
	domain/apid v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1
	user-srv v0.0.0-00010101000000-000000000000
)
