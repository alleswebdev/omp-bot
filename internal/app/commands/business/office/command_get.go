package office

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *OfficeCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong args. Id must be a non-zero integer",
		))

		return
	}

	entity, err := c.officeService.Describe(idx)

	if err != nil {
		log.Printf("fail to get entity with id %d: %v", idx, err)

		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("fail to get entity with id %d", idx),
		))
	}

	c.SendMsg(tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		entity.String()))
}
