package domain

type PurchaseRequestGetResponse struct {
	Status string `json:"status"`
}

type PurchaseRequestPostRequest struct {
	Name string `json:"name"`
}

type PurchaseRequestPostResponse struct {
	Message string `json:"message"`
}

type PurchaseRequestUseCase interface {
	PurchaseRequestGet() (*PurchaseRequestGetResponse, error)
	PurchaseRequestPost(PurchaseRequestPostRequest) (*PurchaseRequestPostResponse, error)
}

type PurchaseRequestRepository interface {
	PurchaseRequestGet() (*PurchaseRequestGetResponse, error)
	PurchaseRequestPost(PurchaseRequestPostRequest) (*PurchaseRequestPostResponse, error)
}
