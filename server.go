package saloon

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	httpServer *fasthttp.Server
}

func (s *Server) Run(port string, handler fasthttp.RequestHandler) error {
	s.httpServer = &fasthttp.Server{
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		Handler:            handler,
		MaxRequestBodySize: 10 * 1024 * 1024,
	}
	fmt.Println("server is running")
	return s.httpServer.ListenAndServe(":" + port)
}

func (s *Server) Shutdown() error {
	return s.httpServer.Shutdown()
}
