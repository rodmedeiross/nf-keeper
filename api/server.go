package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)


type Server struct {
	*fiber.App
	connectionStr string
}

func NewServer(connectionStr string, engine *html.Engine) *Server{
	return &Server{
		App: fiber.New(fiber.Config{
			Views: engine,
		}),
		connectionStr: connectionStr,
	}
}

func (s *Server) Start() error {
	return s.App.Listen(s.connectionStr)
}

func(s *Server) InitializeRoutes() {
	s.App.Get("/", IndexHandler)
	s.App.Get("/nfs/", NfHandler)
	s.App.Get("/nfs/:operation", ModalHandler)
}

