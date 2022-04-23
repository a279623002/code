module go-mysql

go 1.14

replace git.chemm.com/backend/lib => ../../../assets/go-package/src/github.com/wcjs/lib

require (
	git.chemm.com/backend/lib v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.5
)
