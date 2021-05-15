//+build wireinject

package main

import (
	"order-manager/internal/biz"
	"order-manager/internal/data"
	"order-manager/internal/service"

	"github.com/google/wire"
)

func initApp() *service.OrderService {
	panic(wire.Build(data.NewOrderRepo, biz.NewOrderUseCase, service.NewOrderService))
}
