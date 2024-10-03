package model

type Book struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Title  string `gorm:"type:varchar(255)"`
	Author string `gorm:"type:varchar(255)"`
}

func (Book) TableName() string {
	return "public.Books"
}
