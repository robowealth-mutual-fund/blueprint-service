package todo

import "database/sql"

type Todo struct {
	Id        string         `sql:"ID"`
	TaskName  string         `sql:"TASK_NAME"`
	Status    sql.NullString `sql:"STATUS"`
	CreatedAt sql.NullInt64  `sql:"CREATED_AT"`
	UpdatedAt sql.NullInt64  `sql:"UPDATED_AT"`
}

func (*Todo) TableName() string {
	return "TODOS"
}
