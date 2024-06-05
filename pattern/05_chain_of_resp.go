package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

*/


import (
	"fmt"
)

// Service интерфейс для выполнения операций над данными и установки следующего сервиса в цепочке обработки.
type Service interface {
	Execute(d *Data)
	SetNext(service Service)
}

// Data структура данных с флагами для получения и обновления источника данных.
type Data struct {
	GetSource    bool // Флаг, указывающий, что данные уже были получены от устройства.
	UpdateSource bool // Флаг, указывающий, что данные должны быть обновлены.
}

// Dev устройство, которое может получать и передавать данные через следующий сервис в цепочке.
type Dev struct {
	Name string       // Имя устройства.
	Next Service      // Ссылка на следующий сервис в цепочке обработки.
}

// Execute метод для выполнения операции получения данных. Если данные уже были получены, они передаются следующему сервису.
func (device *Dev) Execute(d *Data) {
	if d.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", device.Name)
		device.Next.Execute(d)
	} else {
		fmt.Printf("Get data from device [%s]\n", device.Name)
		d.GetSource = true
		device.Next.Execute(d)
	}
}

// SetNext устанавливает следующий сервис в цепочке обработки.
func (device *Dev) SetNext(service Service) {
	device.Next = service
}

// UpdateDataService сервис для обновления данных.
type UpdateDataService struct {
	Name string       // Имя сервиса обновления данных.
	Next Service      // Ссылка на следующий сервис в цепочке обработки.
}

// Execute метод для выполнения операции обновления данных. Если данные уже были обновлены, они передаются следующему сервису.
func (upd *UpdateDataService) Execute(d *Data) {
	if d.UpdateSource {
		fmt.Printf("Data from device [%s] already update.\n", upd.Name)
		upd.Next.Execute(d)
	} else {
		fmt.Printf("Update data from device [%s]\n", upd.Name)
		d.UpdateSource = true
		upd.Next.Execute(d)
	}
}

// SetNext устанавливает следующий сервис в цепочке обработки.
func (upd *UpdateDataService) SetNext(service Service) {
	upd.Next = service
}

// SaveDataService сервис для сохранения данных.
type SaveDataService struct {
	Next Service // Ссылка на следующий сервис в цепочке обработки.
}

// Execute метод для выполнения операции сохранения данных, если данные были обновлены.
func (save *SaveDataService) Execute(d *Data) {
	if!d.UpdateSource {
		fmt.Println("Data not update")
	} else {
		fmt.Println("Data save")
	}
}

// SetNext устанавливает следующий сервис в цепочке обработки.
func (save *SaveDataService) SetNext(service Service) {
	save.Next = service
}

// NewDevice создает новый экземпляр устройства.
func NewDevice(name string) *Dev {
	return &Dev{
		Name: name,
	}
}

// NewUpdateSvc создает новый экземпляр сервиса обновления данных.
func NewUpdateSvc(name string) *UpdateDataService {
	return &UpdateDataService{
		Name: name,
	}
}

// NewSaveDataService создает новый экземпляр сервиса сохранения данных.
func NewSaveDataService() *SaveDataService {
	return &SaveDataService{}
}

// NewData создает новый экземпляр данных.
func NewData() *Data {
	return &Data{}
}






