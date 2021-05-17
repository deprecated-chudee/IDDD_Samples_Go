package product

import "iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"

type Repository interface {
	NextID() (ID, error)
	ListProductsOfTenant(tenantID tenant.ID) ([]*Product, error)
	ProductOfDiscussionInitiationID(tenantID tenant.ID, discussionInitiationID string) (*Product, error)
	ProductOfID(tenantID tenant.ID, id ID) (*Product, error)
	Remove(product *Product) error
	RemoveAll(products []*Product) error
	Save(product *Product) error
	SaveAll(products []*Product) error
}