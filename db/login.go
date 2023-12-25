package db

type Permission int64

const (
	None Permission = iota
	ReadOnly
	Write
)

type Permissions struct {
	AccessAdmin Permission `json:"access_admin"`
}

type Login struct {
	User       string      `gorm:"primaryKey" json:"user"`
	Password   string      `json:"password"`
	Permission Permissions `json:"permission" gorm:"embedded"`
}

func FetchLogin(user string) (Login, error) {
	var login Login
	rs := db.Where("user = ?", user).Find(&login)
	return login, rs.Error
}

func AddLogin(newLogin *Login) error {
	rs := db.Create(newLogin)
	return rs.Error
}

func DeleteLogin(user string) (Login, error) {
	var login Login
	rs := db.Where("user = ?", user).Delete(&login)
	return login, rs.Error
}
