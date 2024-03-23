package user

type Status int8

var (
	USER_STATUS_UNDEFINED     Status = 0
	USER_STATUS_PREREGISTERED Status = 1
	USER_STATUS_ACTIVE        Status = 2
	USER_STATUS_SUSPEND       Status = 3
)

type Create struct {
	Name   string
	Status Status
}

type Edit struct {
	ID   string
	Name string
}

type CreatePhoneNumber struct {
	PhoneNumber string
	UserID      string
}

type UpdatePhoneNumber struct {
	ID          string
	PhoneNumber string
	UserID      string
}

type CreateEmail struct {
	Email  string
	UserID string
}

type UpdateEmail struct {
	ID     string
	Email  string
	UserID string
}

type CreateDeviceInstance struct {
	DeviceID   string
	InstanceID string
	UserID     string
}

type UpdateDeviceInstance struct {
	ID         string
	DeviceID   string
	InstanceID string
	UserID     string
}
