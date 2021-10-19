package office

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const BaseLimit = 2
const BaseOffset = 0

func (c *OfficeCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Entity list page 1: \n\n"

	entities, err := c.officeService.List(BaseOffset, BaseLimit)

	for _, e := range entities {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, err := json.Marshal(CallbackListData{
		Cursor: BaseOffset + BaseLimit,
		Limit:  BaseLimit,
	})

	if err != nil {
		log.Printf("OfficeCommander.List: "+
			"error marshal json data for type CallbackListData %v", err)
		return
	}

	callbackPath := path.CallbackPath{
		Domain:       "business",
		Subdomain:    "office",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.SendMsg(msg)
}
