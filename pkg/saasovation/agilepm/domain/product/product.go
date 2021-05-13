package product

import (
	"iddd_samples_go/pkg/saasovation/agilepm/domain"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/discussion"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/product/backlogitem"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/team"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"
)

type Product struct {
	domain.Entity

	tenantID               tenant.ID
	productID              ProductID
	productOwnerID         team.ProductOwnerID
	name                   string
	description            string
	discussion             Discussion
	discussionInitiationID string
	backlogItems           []*backlogitem.BacklogItem
}

func NewProduct(tenantID tenant.ID, productID ProductID, productOwnerID team.ProductOwnerID, name string, description string, discussionAvailability discussion.Availability) (*Product, error) {
	discussion, err := fromAvailability(discussionAvailability)
	if err != nil {
		return nil, err
	}

	product := &Product{
		tenantID:       tenantID,
		productID:      productID,
		productOwnerID: productOwnerID,
		name:           name,
		description:    description,
		discussion:     *discussion,
		backlogItems:   make([]*backlogitem.BacklogItem, 0),
	}

	// TODO: DomainEventPublisher

	return product, nil
}
