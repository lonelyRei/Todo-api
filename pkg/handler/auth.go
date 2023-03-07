package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo"
)

// signUp - хендлер регистрации (/auth/sign-up)
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	// Получаем структуру из тела запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Передаем на уровень сервиза структуру, заносим пользователя в бд
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Возвращаем 200 код и id сгенерированное для нового пользователя
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// signInUser - Структуры тела запроса для авторизации
type signInUser struct {
	UserName string `json:"user_name" binding:"required"` // Ник пользователя
	Password string `json:"password" binding:"required"`  // Пароль
}

// signIn - хендлер авторизации (/auth/sign-in)
func (h *Handler) singIn(c *gin.Context) {
	var input signInUser

	// Получаем тело запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Генерируем токен
	token, err := h.services.Authorization.GenerateToken(input.UserName, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	// Возвращаем 200 код и токен
	c.JSON(http.StatusOK, map[string]interface{}{"token": token, "user_name": input.UserName})
}
