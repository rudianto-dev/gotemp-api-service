package client

type DeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteResponse struct {
	Client ClientResponse `json:"client"`
}
