//go:build wireinject
// +build wireinject

package wire

import (
	"nunu-rcc-stack/internal/handler"
	"nunu-rcc-stack/internal/repository"
	"nunu-rcc-stack/internal/server"
	"nunu-rcc-stack/internal/service"
	"nunu-rcc-stack/pkg/app"
	"nunu-rcc-stack/pkg/jwt"
	"nunu-rcc-stack/pkg/log"
	"nunu-rcc-stack/pkg/server/http"
	"nunu-rcc-stack/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
