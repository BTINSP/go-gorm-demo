package modal


type User struct {
	Id int 				`gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Username string		`gorm:"column:username;type:varchar" json:"username"`
	Password string		`gorm:"column:password;type:varchar" json:"password"`
	Role int			`gorm:"column:role;type:int" json:"role"`
}