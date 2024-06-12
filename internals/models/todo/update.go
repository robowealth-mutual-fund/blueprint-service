package todo

type UpdateRequest struct {
	Id       string
	TaskName string
	Status   string
}

type UpdateResponse struct {
	Id string
}
