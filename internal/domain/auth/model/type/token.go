package auth

type Token struct {
	Value     string
	ExpiredAt int64
}
