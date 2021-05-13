package product

import "iddd_samples_go/pkg/saasovation/agilepm/domain"

type ProductID struct {
	domain.ValueObject

	id string
}

func NewProductID(id string) *ProductID {
	return &ProductID{id: id}
}

func (i ProductID) ID() string {
	return i.id
}

func (i *ProductID) SetID(id string) error {
	if err := i.AssertArgumentNotEmpty(id, "The id must be provided."); err != nil {
		return err
	}
	if err := i.AssertArgumentMaxLength(id, 36, "The id must be 36 characters or less."); err != nil {
		return err
	}

	i.id = id
	return nil
}