package main

import (
	"fmt"
	"github.com/slack-go/slack"
)

func notifyToChannel(bookmark MyBookmark, mentionUser string) (string, error) {

	message := fmt.Sprintf("<@%s> <%s|%s> をブックマークしました", mentionUser, bookmark.URL, bookmark.Title)

	msgOption := slack.MsgOptionCompose(
		slack.MsgOptionText(message, false),
		slack.MsgOptionIconEmoji(fmt.Sprintf("cat-%s", bookmark.Category.ID)),
		slack.MsgOptionUsername(bookmark.Category.Name),
		slack.MsgOptionAsUser(false),
	)

	api := slack.New(config.BotToken)
	_, timestamp, err := api.PostMessage(
		config.ChannelID,
		msgOption,
	)
	return timestamp, err
}
