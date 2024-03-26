package auth

type Token struct {
	UserID    string
	Value     string
	ExpiredAt int64
}
