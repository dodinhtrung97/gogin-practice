package entity

type Animal struct {
	ID       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	IsFeline bool   `json:"isFeline" gorm:"type:tinyint"`
	Breed    string `json:"breed" gorm:"type:varchar(100)"`
}
