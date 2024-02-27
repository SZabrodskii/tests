package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
)

func NewHTTPServer(lifecycle fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	server := &http.Server{Addr: ":8070", Handler: mux}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listen, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			log.Info("HTTP Server starts on port", zap.String("addr", server.Addr))
			go server.Serve(listen)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (eh *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(w, r.Body)
	if err != nil {
		eh.log.Info("Failed to handler request", zap.Error(err))
	}
}

type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (hh *HelloHandler) Pattern() string {
	return "/hello"
}

func (hh *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		hh.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		hh.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

type Route interface {
	http.Handler
	Pattern() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			AsRoute(NewEchoHandler),
			AsRoute(NewHelloHandler),
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`)),
			zap.NewExample),
		fx.Invoke(func(server *http.Server) {}),
	).Run()
}
