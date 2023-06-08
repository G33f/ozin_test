package api

import (
	"ShortURL/internal/logging"
	shortenerHandler "ShortURL/internal/shortener/handler"
	"ShortURL/internal/shortener/model"
	shortenerRepo "ShortURL/internal/shortener/repo"
	shortenerUseCase "ShortURL/internal/shortener/usecase"
	"ShortURL/internal/storage"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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
	repo     *pgxpool.Pool

	inMemory bool

	log *logging.Logger
}

func (api *API) Init() {
	var err error
	ctx := context.Background()
	api.checkFlag()
	if !api.inMemory {
		api.repo, err = storage.NewStorage(ctx, api.log)
		if err != nil {
			api.log.Fatal(err)
		}
	}
	address := fmt.Sprintf("%s:%s",
		viper.GetString("APIServer.host"),
		viper.GetString("APIServer.port"))
	lis, err := net.Listen(connectionType, address)
	if err != nil {
		api.log.Fatalf("failed to listen: %v", err)
	}
	api.server = grpc.NewServer()
	api.listener = &lis
	repo := shortenerRepo.NewRepo(api.repo, api.log)
	useCase := shortenerUseCase.NewUseCase(repo, api.log, api.inMemory)
	handler := shortenerHandler.NewHandler(useCase, api.log)
	model.RegisterShortURLServer(api.server, handler)
}

func (api *API) Start() {
	reflection.Register(api.server)
	api.log.Info("api is ready")
	if err := api.server.Serve(*api.listener); err != nil {
		api.log.Fatalf("failed to serve: %v", err)
	}
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
			return fmt.Errorf("flag error")
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
