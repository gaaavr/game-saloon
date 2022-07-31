package handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"saloon/pkg/service"
)

// Структура для работы сервера с хранилищем данных
type Handler struct {
	services *service.Service
}

// Функция конструктор для Handler
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
	// Эндпоинты для бармена (показать список напитков и создать напиток)
	barman := router.Group("/barman")
	{
		barman.Get("/list")
		barman.Post("/create")
	}

	// Эндпоинты для клиента бара (показать данные клиента, показать список доступных напитков, купить напиток)
	visitor := router.Group("/visitor")
	{
		visitor.Get("/me", h.getData)
		visitor.Get("/list", h.getDrinks)
		visitor.Post("/buy", h.buyDrink)
	}
	return router
}
