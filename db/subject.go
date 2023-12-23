package db

type Subject struct {
	Name string `gorm:"primaryKey" json:"name" form:"name"`
}

func SubjectList() ([]Subject, error) {
	var subjects []Subject
	rs := db.Find(&subjects)
	return subjects, rs.Error
}

func FetchSubjectByName(name string) (Subject, error) {
	var subject Subject
	rs := db.Where("name = ?", name).Find(&subject)
	return subject, rs.Error
}

func AddSubject(newSubject *Subject) (Subject, error) {
	rs := db.Create(newSubject)
	return *newSubject, rs.Error
}

func UpdateSubject(newSubject *Subject, name string) (Subject, error) {
	rs := db.Where("name = ?", name).Updates(&newSubject)
	return *newSubject, rs.Error
}

func DeleteSubject(name string) (Subject, error) {
	var subject Subject
	rs := db.Where("name = ?", name).Delete(&subject)
	return subject, rs.Error
}
