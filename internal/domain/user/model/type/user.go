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
