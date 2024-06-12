package todo

type GetRequest struct {
	Id string
}

type GetResponse struct {
	Id        string
	TaskName  string
	Status    string
	CreatedAt int64
	UpdatedAt int64
}
