package service

import "crud_chi_with_text/interfaces"

type TextFileService struct {
	Repository interfaces.FileRepository
}

func NewTextFileService(repo interfaces.FileRepository) *TextFileService {
	return &TextFileService{Repository: repo}
}

func (t *TextFileService) CreateData(data string) {
	t.Repository.Create(data)
}

func (t *TextFileService) ReadAllData() []string {
	return t.Repository.ReadAll()
}

func (t *TextFileService) FindByID(targetID int) string {
	return t.Repository.FindById(targetID)
}

func (t *TextFileService) UpdateData(index int, newData string) {
	t.Repository.Update(index, newData)
}

func (t *TextFileService) DeleteData(index int) {
	t.Repository.Delete(index)
}
