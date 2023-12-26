package db

type Grade struct {
	Name string `gorm:"primaryKey" json:"name" form:"name"`
}

func ItemsBySubjectAndGradeMap(subjectName string, query string) (map[string][]Item, error) {
	var items []Item
	rs := db.
		Joins("JOIN subjects ON items.subject_name = subjects.name").
		Where("subjects.name = ?", subjectName).
		Where("items.title LIKE ? OR items.grade_name LIKE ? OR items.author LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Order("items.title asc").
		Find(&items)

	if rs.Error != nil {
		return nil, rs.Error
	}

	// Create a map to organize items by grade
	itemsByGrade := make(map[string][]Item)

	// Group items by grade
	for _, item := range items {
		itemsByGrade[item.GradeName] = append(itemsByGrade[item.GradeName], item)
	}

	return itemsByGrade, nil
}

func FetchGradeByName(name string) (Grade, error) {
	var grade Grade
	rs := db.Where("name = ?", name).Find(&grade)
	return grade, rs.Error
}

func AddGrade(newGrade *Grade) (Grade, error) {
	rs := db.Create(newGrade)
	return *newGrade, rs.Error
}

func UpdateGrade(newGrade *Grade, name string) (Grade, error) {
	rs := db.Where("name = ?", name).Updates(&newGrade)
	return *newGrade, rs.Error
}

func DeleteGrade(name string) (Grade, error) {
	var grade Grade
	rs := db.Where("name = ?", name).Delete(&grade)
	return grade, rs.Error
}
