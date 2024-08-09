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

func (pr *PurchaseRequestRepository) PurchaseRequestGet() (*client.GetPrResponse, error) {
	resp, err := pr.prClient.GetPr()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
