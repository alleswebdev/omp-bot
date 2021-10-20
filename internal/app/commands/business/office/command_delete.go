package office

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OfficeCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong args. Id must be a non-zero integer",
		))

		return
	}

	result, err := c.officeService.Remove(idx)

	if err != nil {
		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Error while deleting an entity:%s", err),
		))
		return
	}

	if !result {
		c.SendMsg(tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Error while deleting an entity: false result",
		))
		return
	}

	c.SendMsg(tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Entity id:%d deleted", idx),
	))
}
