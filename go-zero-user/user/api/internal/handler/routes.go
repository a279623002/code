package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"go-zero-user/user/api/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/:id",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/:id/orders",
				Handler: GetUserOrdersHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/:id/orders",
				Handler: PlaceOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger",
				Handler: SwaggerUIHandler(),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger/swagger.json",
				Handler: SwaggerHandler(),
			},
		},
	)
}
