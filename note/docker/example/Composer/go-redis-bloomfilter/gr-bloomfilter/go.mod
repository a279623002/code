module gr-bloomfilter

go 1.16

replace git.chemm.com/backend/lib => ../../../../../assets/go-package/src/github.com/wcjs/lib

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/go-ini/ini v1.67.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.6 // indirect
)
