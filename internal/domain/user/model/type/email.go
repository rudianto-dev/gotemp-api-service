package user

type CreateEmail struct {
	Email  string
	UserID string
}

type UpdateEmail struct {
	ID     string
	Email  string
	UserID string
}
