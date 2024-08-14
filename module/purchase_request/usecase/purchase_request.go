package usecase

import (
	"prisca-orchestrator/domain"
	"time"
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

func (pu *PurchaseRequestUseCase) PurchaseRequestGet() (*domain.PurchaseRequestGetResponse, error) {
	return pu.purchaseRequestRepo.PurchaseRequestGet()
}

func (pu *PurchaseRequestUseCase) PurchaseRequestPost(req domain.PurchaseRequestPostRequest) (*domain.PurchaseRequestPostResponse, error) {
	return pu.purchaseRequestRepo.PurchaseRequestPost(req)
}
