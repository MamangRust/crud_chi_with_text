package interfaces

type FileService interface {
	CreateData(data string)
	ReadAllData() []string
	FindByID(targetID int) string
	UpdateData(index int, newData string)
	DeleteData(index int)
}
