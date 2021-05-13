package tenant

import "iddd_samples_go/pkg/saasovation/agilepm/domain"

type ID struct {
	domain.ValueObject

	id string
}

func NewID(id string) ID {
	return ID{id: id}
}

func (i ID) ID() string {
	return i.id
}

func (i *ID) SetID(id string) error {
	if err := i.AssertArgumentNotEmpty(id, "The tenant identity is required."); err != nil {
		return err
	}
	if err := i.AssertArgumentMaxLength(id, 36, "The tenant identity must be 36 characters or less."); err != nil {
		return err
	}

	i.id = id
	return nil
}
