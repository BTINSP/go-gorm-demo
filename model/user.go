package modal


type User struct {
	Id int 				`gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Username string		`gorm:"column:username;type:varchar"`
	Password string		`gorm:"column:password;type:varchar"`
	role int			`gorm:"column:role;type:int"`
}