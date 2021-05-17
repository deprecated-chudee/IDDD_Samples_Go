package backlogitem

type Type string

const (
	Feature     Type = "FEATURE"
	Enhancement Type = "ENHANCEMENT"
	Defect      Type = "DEFECT"
	Foundation  Type = "FOUNDATION"
	Integration Type = "INTEGRATION"
)

func (t Type) IsFeature() bool {
	return t == Feature
}

func (t Type) IsEnhancement() bool {
	return t == Enhancement
}

func (t Type) IsDefect() bool {
	return t == Defect
}

func (t Type) IsFoundation() bool {
	return t == Foundation
}

func (t Type) IsIntegration() bool {
	return t == Integration
}