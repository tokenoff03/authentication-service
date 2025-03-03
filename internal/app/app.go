package app

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/tokenoff03/authentication-service/internal/config"
	"github.com/tokenoff03/authentication-service/pkg/access_v1"
	"github.com/tokenoff03/authentication-service/pkg/auth_v1"

	"github.com/tokenoff03/lib_ad1lek/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server

	path string
}

func NewApp(ctx context.Context, path string) (*App, error) {
	a := &App{path: path}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.path)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)
	auth_v1.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))
	access_v1.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.AccessImpl(ctx))
	return nil

}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is runnning on %s", a.serviceProvider.GRCPConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.grpcConfig.Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
