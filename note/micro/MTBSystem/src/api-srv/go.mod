module api-srv

go 1.16

require (
	github.com/micro/micro/v3 v3.9.0
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
)

replace domain/apid => ../domain/apid

replace user-srv => ../user-srv
