package client

type DetailRequest struct {
	ID string `json:"id" validate:"required"`
}

type DetailResponse struct {
	Client ClientResponse `json:"client"`
}
