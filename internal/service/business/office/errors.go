package office

import "errors"

var (
	ErrorNotFound  = errors.New("entity with id %d not found")
	ErrorEmptyList = errors.New("entity list is empty")
	ErrorOutRange  = errors.New("the requested page is out of range")
)
