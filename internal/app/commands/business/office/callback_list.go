package office

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *OfficeCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var parsedData CallbackListData

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	entities, err := c.officeService.List(parsedData.Cursor, parsedData.Limit)

	if err != nil {
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			fmt.Sprintf("list error: %s", err),
		)

		serializedData, err := json.Marshal(CallbackListData{
			Cursor: BaseOffset,
			Limit:  BaseLimit,
		})

		if err != nil {
			log.Printf("OfficeCommander.CallbackList: "+
				"error reading json data for type CallbackListData from "+
				"input string %v - %v", callbackPath.CallbackData, err)
			return
		}

		callbackPath.CallbackData = string(serializedData)

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("First page", callbackPath.String()),
			))

		c.SendMsg(msg)
		return
	}

	outputMsgText := fmt.Sprintf("Entity list page %d: \n\n", (parsedData.Cursor/parsedData.Limit)+1)

	for _, e := range entities {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		outputMsgText,
	)

	serializedData, err := json.Marshal(CallbackListData{
		Cursor: parsedData.Cursor + BaseLimit,
		Limit:  BaseLimit,
	})

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error marshal json data for type CallbackListData %v", err)
		return
	}

	callbackPath.CallbackData = string(serializedData)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.SendMsg(msg)
}
