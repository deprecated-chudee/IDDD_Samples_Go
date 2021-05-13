package product

import (
	"errors"
	"iddd_samples_go/pkg/saasovation/agilepm/domain"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/discussion"
)

var (
	ErrCannotBeCreatedReady = errors.New("cannot be created ready")
)

type Discussion struct {
	domain.ValueObject

	availability discussion.Availability
	descriptor discussion.Descriptor
}

func NewProductDiscussion(availability discussion.Availability, descriptor discussion.Descriptor) *Discussion {
	return &Discussion{availability: availability, descriptor: descriptor}
}

func fromAvailability(availability discussion.Availability) (*Discussion, error) {
	if availability.IsReady() {
		return nil, ErrCannotBeCreatedReady
	}

	descriptor := discussion.NewDescriptor(discussion.UndefinedId)

	return NewProductDiscussion(availability, *descriptor), nil
}

