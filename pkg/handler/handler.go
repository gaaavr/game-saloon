package handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon/pkg/service"
)

// Handler - структура для работы сервера с хранилищем данных
type Handler struct {
	services *service.Service
}

// NewHandler - функция конструктор для Handler
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Routing - метод для регистрации всех хендлеров в мультиплексоре
func (h *Handler) Routing() *routing.Router {
	router := routing.New()
	// Эндпоинты для регистрации и авторизации
	auth := router.Group("/auth")
	{
		auth.Post("/register", h.register)
		auth.Post("/login", h.login)
	}
	api := router.Group("/api", h.userIdentity)
	{
		// Эндпоинты для бармена (показать список напитков и создать напиток)
		barman := api.Group("/barman")
		{
			barman.Get("/list", h.getDrinks)
			barman.Post("/create", h.createDrink)
		}

		// Эндпоинты для клиента бара (показать данные клиента, показать список доступных напитков, купить напиток)
		visitor := api.Group("/visitor")
		{
			visitor.Get("/me", h.getData)
			visitor.Get("/list", h.getDrinks)
			visitor.Post("/buy", h.buyDrink)
		}
	}

	return router
}
