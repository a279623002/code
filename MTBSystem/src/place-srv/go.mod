module place-srv

go 1.16

require (
	config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.4
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.27.1
)

replace config => ../../config
