package usecase

import (
	"prisca-orchestrator/domain"
	"time"

	"github.com/AldiOktavianto/go-domain/module/proc/client"
)

type PurchaseRequestUseCase struct {
	purchaseRequestRepo domain.PurchaseRequestRepository
	contextTimeout      time.Duration
}

func NewPurchaseRequestUseCase(r domain.PurchaseRequestRepository, timeout time.Duration) domain.PurchaseRequestUseCase {
	return &PurchaseRequestUseCase{
		purchaseRequestRepo: r,
		contextTimeout:      timeout,
	}
}

func (pu *PurchaseRequestUseCase) PurchaseRequestGet() (*client.GetPrResponse, error) {
	return pu.purchaseRequestRepo.PurchaseRequestGet()
}
