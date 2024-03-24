package user

type CreatePhoneNumber struct {
	PhoneNumber string
	UserID      string
}

type UpdatePhoneNumber struct {
	ID          string
	PhoneNumber string
	UserID      string
}
