package todo

import (
	todoRepo "github.com/robowealth-mutual-fund/blueprint-service/internals/repositories/todo"
)

type Service struct {
	repository todoRepo.Interface
}

func New(repository todoRepo.Interface) Interface {
	return &Service{
		repository: repository,
	}
}
