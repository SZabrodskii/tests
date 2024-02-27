package main

//
//import (
//	"context"
//	"fmt"
//	"go.uber.org/fx"
//	"go.uber.org/zap"
//	"io"
//	"net"
//	"net/http"
//)
//
//// NewHTTPServer builds an HTTP server that will begin serving requests
//// when the Fx application starts.
//
//func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
//	srv := &http.Server{Addr: ":8095", Handler: mux}
//	lc.Append(fx.Hook{
//		OnStart: func(ctx context.Context) error {
//			ln, err := net.Listen("tcp", srv.Addr)
//			if err != nil {
//				return err
//			}
//			log.Info("Starting HTTP Server", zap.String("addr", srv.Addr))
//			go srv.Serve(ln)
//			return nil
//		},
//		OnStop: func(ctx context.Context) error {
//			return srv.Shutdown(ctx)
//		},
//	})
//	return srv
//}
//
//// EchoHandler is http.Handler that copies its request body
//// back to the response
//
//type EchoHandler struct {
//	log *zap.Logger
//}
//
//// NewEchoHandler builds a new EchoHandler
//
//func NewEchoHandler(log *zap.Logger) *EchoHandler {
//	return &EchoHandler{
//		log: log,
//	}
//}
//
//// ServeHTTP handles an HTTP request to the /echo endpoint.
//
//func (eh *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if _, err := io.Copy(w, r.Body); err != nil {
//		eh.log.Warn("Failed to handle request:", zap.Error(err))
//	}
//}
//
//func AsRoute(f any) any {
//	return fx.Annotate(
//		f,
//		fx.As(new(Route)),
//		fx.ResultTags(`group:"routes"`),
//	)
//
//}
//
//// NewServeMux builds a ServeMux that will route requests
//// to the given Handler.
//// NewServeMux builds a ServeMux that will route requests
//// to the given Route.
//
//func NewServeMux(routes []Route) *http.ServeMux {
//	mux := http.NewServeMux()
//	for _, route := range routes {
//		mux.Handle(route.Pattern(), route)
//	}
//
//	return mux
//}
//
//// Route is a http.Handler that knows the mux pattern
//// under which it will be registered.
//// Pattern reports the path at which this is registered.
//
//type Route interface {
//	http.Handler
//	Pattern() string
//}
//
//func (*EchoHandler) Pattern() string {
//	return "/echo"
//}
//
//// HelloHandler is an HTTP handler that
//// prints a greeting to the user.
//
//type HelloHandler struct {
//	log *zap.Logger
//}
//
//// NewHelloHandler builds a new HelloHandler
//
//func NewHalloHandler(log *zap.Logger) *HelloHandler {
//	return &HelloHandler{
//		log: log,
//	}
//}
//
//func (hh *HelloHandler) Pattern() string {
//	return "/hello"
//}
//
//func (hh *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		hh.log.Error("Failed to read request", zap.Error(err))
//		http.Error(w, "Internal server error", http.StatusInternalServerError)
//		return
//	}
//
//	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
//		hh.log.Error("Failed to write response", zap.Error(err))
//		http.Error(w, "Internal server error", http.StatusInternalServerError)
//		return
//	}
//}
//
//var httpModule = fx.Module("server",
//	fx.Provide(
//		NewHTTPServer,
//		fx.Annotate(
//			NewServeMux,
//			fx.ParamTags(`group:"routes"`),
//		)),
//)
//
//var ehModule = fx.Provide(
//	AsRoute(NewEchoHandler),
//)
//
//var hhModule = fx.Provide(
//	AsRoute(NewHalloHandler),
//)
//
//func main() {
//	fx.New(
//		httpModule,
//		ehModule,
//		hhModule,
//		fx.Provide(zap.NewExample),
//		fx.Invoke(func(server *http.Server) {}),
//	).Run()
//}

// ================================ Not in main code =======================================

//	fx.New(
//		fx.Provide(
//			NewHTTPServer,
//			fx.Annotate(
//				NewServeMux,
//				fx.ParamTags(`group:"routes"`)),
//			AsRoute(NewHalloHandler),
//			AsRoute(NewEchoHandler),
//			zap.NewExample),
//		fx.Invoke(func(server *http.Server) {}),
//	).Run()
//
//}
