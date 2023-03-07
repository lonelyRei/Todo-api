package todo

// User - Структура пользователя
type User struct {
	Id       int    `json:"-" db:"id"`                    // Идентификатор пользователя
	Name     string `json:"name" binding:"required"`      // Имя пользователя
	UserName string `json:"user_name" binding:"required"` // Ник пользователя
	Password string `json:"password" binding:"required"`  // Пароль
}
