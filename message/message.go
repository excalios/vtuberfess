package message

func extractMessage(msg message) []string {
	return []string{msg.messageCreate.messageData.text}
}

type message struct {
	id               string
	property         string
	createdTimestamp string
	messageCreate    messageCreate
}

type messageCreate struct {
	senderId    string
	messageData messageData
}

type messageData struct {
	text       string
	attachment []byte
}
