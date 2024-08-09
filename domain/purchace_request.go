package domain

import (
	"github.com/AldiOktavianto/go-domain/module/proc/client"
)

type PurchaseRequestUseCase interface {
	PurchaseRequestGet() (*client.GetPrResponse, error)
}

type PurchaseRequestRepository interface {
	PurchaseRequestGet() (*client.GetPrResponse, error)
}
