package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
/*
Строитель позволяет разбить создание объекта на отдельные этапы и инициализировать только необходимые поля объекта
Можно было бы сделать 1 конструктор, но в него передавалось бы слишком много полей
можно создать разные представления объектов. Заметьте, что Строитель не позволяет посторонним объектам иметь доступ
к конструируемому объекту, пока тот не будет полностью готов. Это защищает клиентский код от получения незаконченных
«битых» объектов
+  Позволяет использовать один и тот же код для создания различных продуктов.
+  Изолирует сложный код сборки продукта от его основной бизнес-логики.
+  Позволяет создавать продукты пошагово.
- Усложняет код программы




*/

import (
	"fmt"
)

// Структура для хранения информации о пользователе
type User struct {
	Name     string
	Email    string
	Password string
}

// Интерфейс Builder определяет методы для установки свойств пользователя
type UserBuilder interface {
	SetName(name string) UserBuilder
	SetEmail(email string) UserBuilder
	SetPassword(password string) UserBuilder
	Build() *User
}

// Конкретный класс Builder реализует интерфейс UserBuilder
type ConcreteUserBuilder struct{}

func (b *ConcreteUserBuilder) SetName(name string) UserBuilder {
	b.Name = name
	return b
}

func (b *ConcreteUserBuilder) SetEmail(email string) UserBuilder {
	b.Email = email
	return b
}

func (b *ConcreteUserBuilder) SetPassword(password string) UserBuilder {
	b.Password = password
	return b
}

func (b *ConcreteUserBuilder) Build() *User {
	return &User{Name: b.Name, Email: b.Email, Password: b.Password}
}

//  func main() {
// 	user := NewUser().SetName("John Doe").SetEmail("john.doe@example.com").SetPassword("password123").Build()

// 	fmt.Printf("User Name: %s\n", user.Name)
// 	fmt.Printf("User Email: %s\n", user.Email)
// 	fmt.Printf("User Password: %s\n", user.Password)
// }

// // Функция NewUser возвращает экземпляр ConcreteUserBuilder
// func NewUser() *ConcreteUserBuilder {
// 	return &ConcreteUserBuilder{}
// }
