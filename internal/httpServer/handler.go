package httpServer

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"task/internal/service/delivery/http"
	"task/internal/service/repository"
	"task/internal/service/usecase"
	"task/pkg/storage"
)

func (s *Server) MapHandlers(app *fiber.App) error {

	rdb, err := storage.InitRedis(s.cfg)
	if err != nil {
		log.Println("ERROR ::: ", err)
		log.Fatal("no connect redis")
	}

	serviceRepo := repository.NewRedisRepository(s.cfg, rdb)

	serviceUC := usecase.NewServiceUsecase(s.cfg, serviceRepo)

	serviceR := http.NewServiceHandler(serviceUC)

	http.MapRoutes(app, serviceR)

	return nil
}
