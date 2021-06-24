package entity

type Household struct {
	ID       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	AnimalId uint64 `json:"-"`
	Address  string `json:"address"`
}
