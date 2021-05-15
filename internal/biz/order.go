package biz

import "context"

//DO
type Order struct {
	Id     int64
	Amount int64
}

//repo接口
type OrderRepo interface {
	GetOrder(ctx context.Context, id int64) (*Order, error)
}

//领域对象服务，编排多个领域对象完成业务逻辑
type OrderUseCase struct {
	repo OrderRepo
}

//注入repo
func NewOrderUseCase(repo OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
	}
}

func (oc *OrderUseCase) Get(ctx context.Context, id int64) (*Order, error) {
	return oc.repo.GetOrder(ctx, id)
}
