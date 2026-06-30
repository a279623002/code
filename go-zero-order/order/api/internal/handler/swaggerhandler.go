package handler

import (
	_ "embed"
	"net/http"
)

//go:embed docs/swagger.json
var swaggerJSON []byte

//go:embed docs/swagger.html
var swaggerHTML []byte

// SwaggerHandler 返回 swagger.json
func SwaggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerJSON)
	}
}

// SwaggerUIHandler 返回 swagger-ui 页面
func SwaggerUIHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(swaggerHTML)
	}
}
