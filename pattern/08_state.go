package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/


import "fmt"

// State интерфейс, определяющий поведение для разных состояний устройства.
type State interface {
	RespondLockButton()
}

// MacBook структура, представляющая ноутбук Mac, который может находиться в одном из двух состояний: заблокированном или разблокированном.
type MacBook struct {
	locked       State // Состояние, когда ноутбук заблокирован.
	unlocked     State // Состояние, когда ноутбук разблокирован.
	currentState State // Текущее состояние ноутбука.
}

// LockedState структура, представляющая состояние ноутбука, когда он заблокирован.
type LockedState struct{}

// RespondLockButton реализует метод интерфейса State для состояния LockedState.
func (s *LockedState) RespondLockButton() {
	fmt.Println("Trynna to unlock") // Вывод сообщения о попытке разблокировки.
}

// UnlockedState структура, представляющая состояние ноутбука, когда он разблокирован.
type UnlockedState struct{}

// RespondLockButton реализует метод интерфейса State для состояния UnlockedState.
func (s *UnlockedState) RespondLockButton() {
	fmt.Println("Trynna to lock") // Вывод сообщения о попытке блокировки.
}
