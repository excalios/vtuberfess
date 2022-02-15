package message

import "testing"

func TestMessages(t *testing.T) {
	t.Run("Extracting Messages", func(t *testing.T) {

		content := "nec feugiat in fermentum posuere"
		message := message{
			id:               "1234858592",
			createdTimestamp: "1392078023603",
			messageCreate: messageCreate{
				senderId: "3805104374",
				messageData: messageData{
					text: content,
				},
			},
		}

		got := extractMessage(message)

		if got != content {
			t.Errorf("got %q want %q", got, content)
		}
	})
}
