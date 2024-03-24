package otp

type ChannelType int8

var (
	OTP_CHANNEL_TYPE_UNDEFINED ChannelType = 0
	OTP_CHANNEL_TYPE_SMS       ChannelType = 1
	OTP_CHANNEL_TYPE_EMAIL     ChannelType = 2
)

type CreateOTP struct {
	Receiver    string
	ChannelType ChannelType
}
