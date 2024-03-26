package client

type Status int8

var (
	CLIENT_STATUS_UNDEFINED Status = 0
	CLIENT_STATUS_ACTIVE    Status = 1
	CLIENT_STATUS_INACTIVE  Status = 2
)

type Create struct {
	ClientID     string
	ClientSecret string
	Status       Status
	ExpiredAt    int64
}

type Edit struct {
	ID           string
	ClientID     string
	ClientSecret string
	Status       Status
	ExpiredAt    int64
}
