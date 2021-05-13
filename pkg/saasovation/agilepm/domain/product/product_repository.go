package product

import "iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"

type Repository interface {
	AllProductsOfTenant(tenantID *tenant.ID) ([]*Product, error)
	NextIdentity() *ProductID
	ProductOfDiscussionInitiationID(tenantID *tenant.ID, discussionInitiationID string) (*Product, error)
	ProductOfID(tenantID *tenant.ID, productID *ProductID) (*Product, error)
	Remove(product *Product) error
	RemoveAll(products []*Product) error
	Save(product *Product) error
	SaveAll(products []*Product) error
}