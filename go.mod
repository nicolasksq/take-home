module server

// +heroku goVersion 1.16
// +heroku install ./cmd/pkg/
go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/stretchr/testify v1.4.0
)
