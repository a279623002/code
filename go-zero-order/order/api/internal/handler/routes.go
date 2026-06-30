package handler

import (
	"net/http"

	"go-zero-service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order",
				Handler: CreateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/order/:id",
				Handler: GetOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/order/list",
				Handler: ListOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/order/:id/status",
				Handler: UpdateOrderStatusHandler(serverCtx),
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
