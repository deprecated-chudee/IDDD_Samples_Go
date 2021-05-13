package team

import (
	"iddd_samples_go/pkg/saasovation/agilepm/domain"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"
)

type ProductOwnerID struct {
	domain.ValueObject

	id string
	tenantID tenant.ID
}

func NewProductOwnerID(tenantID tenant.ID, id string) ProductOwnerID {
	return ProductOwnerID{tenantID: tenantID, id: id}
}

func (i ProductOwnerID) ID() string {
	return i.id
}

func (i ProductOwnerID) TenantID() tenant.ID {
	return i.tenantID
}

func (i *ProductOwnerID) SetID(id string) error {
	if err := i.AssertArgumentNotEmpty(id, "The id must be provided."); err != nil {
		return err
	}
	if err := i.AssertArgumentMaxLength(id, 36, "The id must be 36 characters or less."); err != nil {
		return err
	}

	i.id = id
	return nil
}

func (i *ProductOwnerID) SetTenantID(tenantID tenant.ID) error {
	if err := i.AssertArgumentNotNil(tenantID, "TThe tenantId must be provided"); err != nil {
		return err
	}

	i.tenantID = tenantID
	return nil
}