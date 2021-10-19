package office

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OfficeCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`/help__business__office - help
		/list__business__office - list products
		/get__business__office id - get entity by id
		/delete__business__office id - remove entity by id
		/create__business__office {"name":"name", "description":"description"} - create new entity by json string
		/edit__business__office {"id":1, "name":"name", "description":"description"} - edit entity by json string
	`,
	)

	c.SendMsg(msg)
}
