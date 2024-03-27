package alert

type Channel int8

var (
	ALERT_CHANNEL_UNDEFINED        Channel = 0
	ALERT_CHANNEL_TELEGRAM         Channel = 1
	ALERT_CHANNEL_ONE_NOTIFICATION Channel = 2
)

type Create struct {
	Message  string
	Sender   string
	Receiver string
	Channel  Channel
}
