package todo

type CreateRequest struct {
	TaskName string
	Status   string
}

type CreateResponse struct {
	TaskName  string
	Status    string
	CreatedAt int64
	UpdatedAt int64
}
