package todo

import "errors"

// TodoList - Структура, описывающая список записей
type TodoList struct {
	Id          int    `json:"id" db:"id"`                          // Идентификатор
	Title       string `json:"title" db:"title" binding:"required"` // Заголовок
	Description string `json:"description" db:"description"`        // Описание
}

// UserList - Структуру, необходимая для связи многие ко многим,
// связывает пользователя и список зааписей
type UserList struct {
	Id     int // Идентификатор
	UserId int // Идентификатор пользователя
	ListId int // Идентификатор списка записей
}

// TodoItem - Структура, описывающая конкретную запись в списке
type TodoItem struct {
	Id          int    `json:"id"`           // Идентификатор
	Title       string `json:"title"`        // Заголовок
	Description string `json:"description""` // Описание
	Done        bool   `json:"done"`         // Флаг завершенности
}

// ListsItem - Структура, обеспечивающая связь многие ко многим,
// свззывает список записей и конкретную запись
type ListsItem struct {
	Id     int // Идентификатор
	ListId int // Идентификатор списка
	ItemId int // Идентификатор записи
}

// UpdateListInput - Структура для запроса на изменение списка
type UpdateListInput struct {
	Title       *string `json:"title"`       // Делаем указатель, чтобы в случае пустого параметра был nil
	Description *string `json:"description"` // Аналогино
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
