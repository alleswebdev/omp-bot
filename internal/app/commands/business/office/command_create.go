package office

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *OfficeCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var parsedData CreateRequest

	err := json.Unmarshal([]byte(args), &parsedData)

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error reading json data for type Office from "+
			"input string %v - %v", args, err)

		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`Wrong json struct. Please send like this: {"name":"name", "description":"description"}`,
		))
		return
	}

	id, err := c.officeService.Create(business.Office{
		Name:        parsedData.Name,
		Description: parsedData.Description,
	})

	if err != nil {
		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Error while creating an entity:%s", err),
		))
		return
	}

	c.SendMsg(tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Entity was added, id:%d", id)))
}
