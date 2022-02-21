module api-srv

go 1.16

require (
	domain/apid v0.0.0-00010101000000-000000000000
	github.com/micro/micro/v3 v3.9.0
)

replace domain/apid => ../domain/apid

replace user-srv => ../user-srv
