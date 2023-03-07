package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo"
)

// createList - Создает список
func (h *Handler) createList(c *gin.Context) {
	// Получаем id пользователя
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Читаем тело запроса
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Обращаемся к уровню сервиса для создания записи
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем 200 код и json с id списка
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// getAllListsResponse - структура ответа для запроса на получение всех записей
type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// getAllLists - возвращает все списки, принадлежащие пользователю
func (h *Handler) getAllLists(c *gin.Context) {
	// Получаем id пользователя
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Передаем запрос на сервис
	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем ответ
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})

}

// getListById - Получает список по id
func (h *Handler) getListById(c *gin.Context) {
	// Получаем id пользователя
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

// updateList - Обновляет список по id
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// deleteList - Удаляет список по id
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.DeleteList(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
