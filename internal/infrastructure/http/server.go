package http

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	return &Server{engine: r}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
func (s *Server) Engine() *gin.Engine {
	return s.engine
}
