package entities

type Mock struct {
	ID        string `sql:"ID"`
	Code      string `sql:"CODE"`
	NameTH    string `sql:"NAME_TH"`
	NameEN    string `sql:"NAME_EN"`
	Version   string `sql:"VERSION"`
	CreatedAt int64  `sql:"CREATED_AT"`
	UpdatedAt int64  `sql:"UPDATED_AT"`
}

func (ent *Mock) TableName() string {
	return "MOCK"
}
