package db

import "time"

type Item struct {
	ID uint `gorm:"primaryKey" json:"id" form:"id"`

	Title   string `json:"title" form:"title"`
	Summary string `json:"summary" form:"summary"`
	Comment string `json:"comment" form:"comment"`

	SubjectName string  `json:"subject_name" form:"subject_name"`
	Subject     Subject `gorm:"foreignKey:SubjectName" json:"subject" form:"subject"`

	GradeName string `json:"grade_name" form:"grade_name"`
	Grade     Grade  `gorm:"foreignKey:GradeName" json:"grade" form:"grade"`

	CreatedAt time.Time `json:"created_at" form:"created_at"`
	FileName  string    `json:"file_name" form:"file_name"`

	Author string `json:"author" form:"author"`
}

func FetchItems() ([]Item, error) {
	var books []Item
	rs := db.Find(&books)
	return books, rs.Error
}

func FetchItemById(id uint) (Item, error) {
	var book Item
	rs := db.Where("id = ?", id).Find(&book)
	return book, rs.Error
}

func AddItem(newItem *Item) error {
	// Check if the subject exists, if not, add it
	existingSubject, err := FindOrCreateSubject(newItem.Subject.Name)
	if err != nil {
		return err
	}
	newItem.GradeName = existingSubject.Name

	// Check if the grade exists, if not, add it
	existingGrade, err := FindOrCreateGrade(newItem.Grade.Name)
	if err != nil {
		return err
	}
	newItem.GradeName = existingGrade.Name

	rs := db.Create(newItem)
	return rs.Error
}

func FindOrCreateSubject(subjectName string) (Subject, error) {
	var subject Subject
	rs := db.Where(Subject{Name: subjectName}).FirstOrCreate(&subject)
	return subject, rs.Error
}

func FindOrCreateGrade(gradeName string) (Grade, error) {
	var grade Grade
	rs := db.Where(Grade{Name: gradeName}).FirstOrCreate(&grade)
	return grade, rs.Error
}

func UpdateItem(newItem *Item, id int) (Item, error) {
	rs := db.Where("id = ?", id).Updates(&newItem)
	return *newItem, rs.Error
}

func DeleteItem(id uint) (Item, error) {
	var book Item
	rs := db.Where("id = ?", id).Delete(&book)
	return book, rs.Error
}
