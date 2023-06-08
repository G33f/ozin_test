package api

import (
	shortUrlHandler "ShortURL/internal/handler"
	"ShortURL/internal/logging"
	"ShortURL/internal/model"
	"ShortURL/internal/postgres"
	shortUrlRepo "ShortURL/internal/repo"
	shortUrlUseCase "ShortURL/internal/usecase"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	connectionType = "tcp"
	storageType    = "storage_type"
	usage          = "This program works in two modes:\n" +
		"Save data to PostgresSQL \"PostgreSQL\" or in memory \"InMemory\".\n\n" +
		"If you do not specify the parameter, the program will work in " +
		"\"InMemory\" mode."
)

type API struct {
	server   *grpc.Server
	listener *net.Listener

	inMemory bool

	log *logging.Logger
}

func (api *API) Init() (err error) {
	ctx := context.Background()
	api.checkFlag()
	address := fmt.Sprintf("%s:%s",
		viper.GetString("APIServer.host"),
		viper.GetString("APIServer.port"))
	lis, err := net.Listen(connectionType, address)
	if err != nil {
		api.log.Error("failed to listen: %v", err)
		return err
	}
	api.server = grpc.NewServer()
	api.listener = &lis
	var repo shortUrlRepo.Repo
	if api.inMemory {
		repo = shortUrlRepo.NewInMemoryRepo(api.log)
	} else {
		connDB, err := postgres.NewClient(ctx, api.log)
		if err != nil {
			api.log.Error(err)
			return err
		}
		repo = shortUrlRepo.NewPostgresRepo(connDB)
	}
	useCase := shortUrlUseCase.NewUseCase(repo, api.log, api.inMemory)
	handler := shortUrlHandler.NewHandler(useCase, api.log)
	model.RegisterShortURLServer(api.server, handler)
	reflection.Register(api.server)
	return nil
}

func (api *API) Start() error {
	defer api.server.GracefulStop()
	api.log.Info("api is ready")
	if err := api.server.Serve(*api.listener); err != nil {
		api.log.Errorf("failed to serve: %v", err)
		return err
	}
	return nil
}

func (api *API) checkFlag() {
	api.inMemory = true
	flag.Func(storageType, usage, func(flag string) error {
		switch flag {
		case "PostgreSQL":
			api.inMemory = false
		case "InMemory":
			api.inMemory = true
		default:
			return errors.New("flag error")
		}
		return nil
	})
	flag.Parse()
}

func NewApi(log *logging.Logger) *API {
	return &API{
		log: log,
	}
}
