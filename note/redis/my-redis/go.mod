module my-redis

go 1.14

require (
	git.chemm.com/backend/lib v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
)

replace git.chemm.com/backend/lib => ../../../assets/go-package/src/github.com/wcjs/lib
