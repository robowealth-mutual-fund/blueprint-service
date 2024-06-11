package todo

type DeleteRequest struct {
	Id string `json:"id"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
