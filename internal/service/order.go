package service

import (
	"context"
	v1 "order-manager/api/v1"
	"order-manager/internal/biz"
)

//service DTO和DO互转
type OrderService struct {
	orderUseCase *biz.OrderUseCase
}

//注入usecase
func NewOrderService(orderUseCase *biz.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: orderUseCase,
	}
}

func (o *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderReq) (*v1.GetOrderReply, error) {
	x, err := o.orderUseCase.Get(ctx, req.OrderId)
	return &v1.GetOrderReply{
		OrderId:     x.Id,
		OrderAmount: x.Amount,
	}, err
}
