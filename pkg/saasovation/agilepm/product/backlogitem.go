package product

import (
	"iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"
)

type BacklogItemID string

type BacklogItem struct {
	backlogItemID BacklogItemID
	ordering      int
	productID     ID
	tenantID      tenant.ID
}

func (b BacklogItem) BacklogItemID() BacklogItemID {
	return b.backlogItemID
}

func (b BacklogItem) Ordering() int {
	return b.ordering
}

func (b BacklogItem) ProductID() ID {
	return b.productID
}

func (b BacklogItem) TenantID() tenant.ID {
	return b.tenantID
}