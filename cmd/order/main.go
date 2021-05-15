package main

import (
	"context"
	"fmt"
	"net"
	v1 "order-manager/api/v1"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return signalHandle(ctx)
	})
	group.Go(func() error {
		return NewGrpcServer(ctx)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("shutdown:", err)
	}

}

//信号处理
func signalHandle(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case sig := <-sigs:
		return errors.Errorf("get os signal: %v", sig)
	}

}

//grpc服务
func NewGrpcServer(ctx context.Context) error {
	server := grpc.NewServer()

	orderService := initApp()

	v1.RegisterOrderServer(server, orderService)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		server.Stop()
	}()
	return server.Serve(lis)
}
