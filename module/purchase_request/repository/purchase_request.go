package repository

import (
	"prisca-orchestrator/domain"

	"github.com/AldiOktavianto/go-domain/module/proc/client"
)

type PurchaseRequestRepository struct {
	prClient client.PrClient
}

func NewPurchaseRequestRepository(prClient client.PrClient) domain.PurchaseRequestRepository {
	return &PurchaseRequestRepository{
		prClient: prClient,
	}
}

func (pr *PurchaseRequestRepository) PurchaseRequestGet() (*domain.PurchaseRequestGetResponse, error) {
	resp, err := pr.prClient.GetPr()

	if err != nil {
		return nil, err
	}

	result := &domain.PurchaseRequestGetResponse{
		Status: resp.Status,
	}

	return result, nil
}

func (pr *PurchaseRequestRepository) PurchaseRequestPost(req domain.PurchaseRequestPostRequest) (*domain.PurchaseRequestPostResponse, error) {
	newPr := client.PostPrRequest{
		Name: req.Name,
	}

	respPost, err := pr.prClient.PostPr(newPr)

	if err != nil {
		return nil, err
	}

	result := &domain.PurchaseRequestPostResponse{
		Message: respPost.Message,
	}

	return result, nil
}
