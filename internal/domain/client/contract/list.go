package client

type ListRequest struct {
}

type ListResponse struct {
	Client []ClientResponse `json:"client"`
}
