package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/


import (
	"fmt"	
) 

// Payment интерфейс, определяющий метод Pay(), который должен быть реализован всеми платежными системами.
type Payment interface {
	Pay() error
}

// cardPayment структура для реализации платежа картой.
type cardPayment struct {
	cardNumber, cvv string // Номер карты и CVV.
}

// Pay реализует метод Pay() для платежа картой.
func (p *cardPayment) Pay() error {
	// Здесь должна быть реализация процесса оплаты картой.
	// Например, проверка данных карты, отправка запроса на платежную систему и т.д.
	fmt.Println("Оплата картой...")
	return nil
}

// paypalPayment структура для реализации платежа через PayPal.
type paypalPayment struct {
	account string // Аккаунт пользователя PayPal.
}

// Pay реализует метод Pay() для платежа через PayPal.
func (p *paypalPayment) Pay() error {
	// Здесь должна быть реализация процесса оплаты через PayPal.
	// Например, авторизация пользователя, отправка запроса на платежную систему и т.д.
	fmt.Println("Оплата через PayPal...")
	return nil
}

// qiwiPayment структура для реализации платежа через Qiwi.
type qiwiPayment struct {
	account string // Аккаунт пользователя Qiwi.
}

// Pay реализует метод Pay() для платежа через Qiwi.
func (q *qiwiPayment) Pay() error {
	// Здесь должна быть реализация процесса оплаты через Qiwi.
	// Например, авторизация пользователя, отправка запроса на платежную систему и т.д.
	fmt.Println("Оплата через Qiwi...")
	return nil
}

// NewCardPayment создает новый экземпляр платежа картой.
func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

// NewPayPalPayment создает новый экземпляр платежа через PayPal.
func NewPayPalPayment(account string) Payment {
	return &paypalPayment{account: account}
}

// NewQiwiPaymant создает новый экземпляр платежа через Qiwi.
// Обратите внимание на опечатку в названии функции, должно быть NewQiwiPayment.
func NewQiwiPayment(account string) Payment {
	return &qiwiPayment{account: account}
}

// processOrder обрабатывает заказ, используя предоставленный платежный метод.
func processOrder(orderNumber string, p Payment) {
	// Здесь должна быть реализация процесса обработки заказа.
	// Например, проверка платежа, подтверждение заказа и т.д.
	err := p.Pay()
	if err!= nil {
		fmt.Printf("Ошибка при оплате: %v\n", err)
		return
	}
	fmt.Printf("Заказ %s успешно оплачен\n", orderNumber)
}
