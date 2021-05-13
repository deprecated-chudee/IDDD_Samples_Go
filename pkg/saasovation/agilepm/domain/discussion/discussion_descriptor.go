package discussion

import "iddd_samples_go/pkg/saasovation/agilepm/domain"

const UndefinedId = "UNDEFINED"

type Descriptor struct {
	domain.ValueObject

	id string
}

func NewDescriptor(id string) *Descriptor {
	return &Descriptor{id: id}
}

func (d *Descriptor) IsUndefined() bool {
	return d.id == UndefinedId
}

func (d *Descriptor) SetID(id string) error {
	if err := d.AssertArgumentNotEmpty(id, "The discussion identity must be provided."); err != nil {
		return err
	}
	if err := d.AssertArgumentMaxLength(id, 36, "The discussion identity must be 36 characters or less."); err != nil {
		return err
	}

	d.id = id
	return nil
}