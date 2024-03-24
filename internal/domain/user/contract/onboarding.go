package user

type QRIS struct {
	ID       string `json:"id" validate:"required"`
	DecodeQR string `json:"decode_qr" validate:"required"`
}

type Merchant struct {
	ID       string  `json:"id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Address  string  `json:"address" validate:"required"`
	Category string  `json:"category" validate:"required"`
	QRIS     []*QRIS `json:"qris" validate:"required,dive,required"`
}

type CIF struct {
	ID          string      `json:"id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	PhoneNumber string      `json:"phone_number" validate:"required"`
	Merchant    []*Merchant `json:"merchant" validate:"required,dive,required"`
}

type OnboardingRequest struct {
	CIF CIF `json:"cif" validate:"required"`
}

type OnboardingResponse struct {
}
