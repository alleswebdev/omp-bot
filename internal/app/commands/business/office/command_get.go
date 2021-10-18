package office

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *OfficeCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Printf("wrong args %#v, err: %s", args, err)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong args. Id must be a non-zero integer",
		)
		_, err = c.bot.Send(msg)

		if err != nil {
			log.Printf("OfficeCommander.Delete: error sending reply message to chat - %v", err)
		}

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)
	entity, err := c.officeService.Describe(idx)

	if err != nil {
		log.Printf("fail to get entity with id %d: %v", idx, err)
		msg.Text = err.Error()
	} else {
		msg.Text = entity.String()
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Get: error sending reply message to chat - %v", err)
	}
}
