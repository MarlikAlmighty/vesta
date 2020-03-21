package bot

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
)

// Config for bot for starting
type Configuration struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	BotToken string `json:"bot_token,omitempty"`
	WebHook  string `json:"web_hook,omitempty"`
}

func Run(cfg *Configuration) error {

	// Start botAPI with token
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	// Set Webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.WebHook + cfg.BotToken))
	if err != nil {
		return err
	}

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(cfg.Host+":"+cfg.Port, nil)

	log.Printf("Starting, bot on: http://%s\n", cfg.Host+":"+cfg.Port)

	for update := range updates {

		var f = false

		if update.Message != nil {

			// Read only toxic trolls Sergey and Milenin on half an hour
			if update.Message.From.ID == 1055865722 || update.Message.From.UserName == "Chickenfresh" {

				tm := int64(update.Message.Date + 1800)

				if api, err := bot.RestrictChatMember(tgbotapi.RestrictChatMemberConfig{
					ChatMemberConfig: tgbotapi.ChatMemberConfig{
						ChatID: update.Message.Chat.ID,
						UserID: update.Message.From.ID,
					},
					CanSendMessages:       &f,
					CanSendMediaMessages:  &f,
					CanSendOtherMessages:  &f,
					CanAddWebPagePreviews: &f,
					UntilDate:             tm,
				}); err != nil {
					log.Printf("Err restrict user: %v\n", api.Result)
				}
			}

			// Delete message if check words vesta from Shatunov
			if checkWords(update.Message.Text) && update.Message.From.ID == 181588695 {
				if api, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
					ChatID:    update.Message.Chat.ID,
					MessageID: update.Message.MessageID,
				}); err != nil {
					log.Fatalf("Error: %v\n", api.Description)
				}
			}
		}

		// Delete message if edit message and check words vesta from Shatunov
		if update.EditedMessage != nil {
			if checkWords(update.EditedMessage.Text) && update.EditedMessage.From.ID == 181588695 {
				if api, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
					ChatID:    update.EditedMessage.Chat.ID,
					MessageID: update.EditedMessage.MessageID,
				}); err != nil {
					log.Fatalf("Error: %v\n", api.Description)
				}

			}
		}

	}
	return nil
}
