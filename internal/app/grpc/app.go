package grpcApp

import (
	"context"
	productsAPI "eshop-products-ms/internal/grpc/server"
	productService "eshop-products-ms/internal/services/product"
	userService "eshop-products-ms/internal/services/user"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func New(log *slog.Logger, productsServ productService.Product, usersServ userService.User, port int) *App {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))
			if e, ok := p.(error); ok {
				sentry.CaptureException(e)
				return status.Errorf(codes.Internal, "internal error")
			}
			genericError := fmt.Errorf("unable to get error, panic: %v", p)
			sentry.CaptureException(genericError)
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	grpcAppServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
		//productsAPI.CorsLikeInterceptor,
	))

	productsAPI.Register(grpcAppServer, productsServ, usersServ)

	return &App{log: log, gRPCServer: grpcAppServer, port: port}
}

func (a *App) Run() error {
	const op = "grpcApp.App.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	a.log.Info("grpc server is listening")

	err = a.gRPCServer.Serve(l)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
