package data

import (
	"context"
	"order-manager/internal/biz"
)

//模拟持久化对象
type orderPo struct {
	Id     int64
	Amount int64
}

//实现repo，进行数据库操作
type orderRepo struct {
}

func NewOrderRepo() biz.OrderRepo {
	return &orderRepo{}
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*biz.Order, error) {
	//模拟数据库操作
	po := orderPo{
		Id:     1,
		Amount: 100,
	}
	return &biz.Order{
		Id:     po.Id,
		Amount: po.Amount,
	}, nil
}
