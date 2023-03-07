package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"todo/pkg/service"
)

// Handler - Структура хендлера http-запросов
type Handler struct {
	services *service.Service // Ссылка на сервис
}

// NewHandler - Конструктор хендлера
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes - Связывает маршруты с соответствующими обработчиками
func (h *Handler) InitRoutes() *gin.Engine {
	// Создаем экземпляр роутера
	router := gin.New()

	// Устанавливаем cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Группа для теста
	test := router.Group("/test")
	{
		test.GET("/", h.testHandler)
	}

	// Группа для авторизации / регистрации
	auth := router.Group("/auth")
	{
		// Роут для регистрации
		auth.POST("/sign-up", h.signUp)

		// Роут для авторизации
		auth.POST("/sign-in", h.singIn)
	}

	// Группа для API
	api := router.Group("/api", h.userIdentity)
	{
		// Группа для работы со списками
		lists := api.Group("/lists")
		{
			// Создание списка
			lists.POST("/", h.createList)

			// Получение всех списков, доступных пользователю
			lists.GET("/", h.getAllLists)

			// Получение списка по id
			lists.GET("/:id", h.getListById)

			// Обновление списка по id
			lists.PUT("/:id", h.updateList)

			// Удаление списка по id
			lists.DELETE("/:id", h.deleteList)

			// Группа для работы с записями в списке
			items := lists.Group(":id/items")
			{
				// Создание записи
				items.POST("/", h.createItem)

				// Получение всех записей
				items.GET("/", h.getAllItems)

				// Получение записи по id
				items.GET("/:item_id", h.getItemById)

				// Обновление записи по id
				items.PUT("/:item_id", h.updateItem)

				// Удаление записи по id
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	// Возвращаем настроенный экземпляр
	return router
}
