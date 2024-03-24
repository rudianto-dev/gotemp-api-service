package user

type CreateCIF struct {
	UserID      string
	ReferenceID string
}

type UpdateCIF struct {
	ID          string
	UserID      string
	ReferenceID string
}
