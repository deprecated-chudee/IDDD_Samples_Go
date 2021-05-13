package domain

import "iddd_samples_go/pkg/saasovation/common"

type Entity struct {
	common.AssertionConcern
	concurrencyVersion int
}

func NewEntity() *Entity {
	return &Entity{concurrencyVersion: 0}
}

func (e *Entity) ConcurrencyVersion() int {
	return e.concurrencyVersion
}

func (e *Entity) SetConcurrencyVersion(concurrencyVersion int) {
	e.concurrencyVersion = concurrencyVersion
}