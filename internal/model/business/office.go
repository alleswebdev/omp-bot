package business

import "fmt"

type Office struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (o *Office) String() string {
	return fmt.Sprintf("id:%d, Office name:%s Description:%s", o.Id, o.Name, o.Description)
}
