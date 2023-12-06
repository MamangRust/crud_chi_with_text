package interfaces

type FileRepository interface {
	Create(data string)
	ReadAll() []string
	FindById(targetID int) string
	Update(index int, newData string)
	Delete(index int)
}
