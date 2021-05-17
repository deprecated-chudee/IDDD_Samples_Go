package backlogitem

import "fmt"

const (
	minRatingsLength = 1
	maxRatingsLength = 9
)

var (
	ErrInvalidBenefit = fmt.Errorf("relative benefit must be between %d and %d", minRatingsLength, maxRatingsLength)
	ErrInvalidCost = fmt.Errorf("relative cost must be between %d and %d", minRatingsLength, maxRatingsLength)
	ErrInvalidPenalty = fmt.Errorf("relative penalty must be between %d and %d", minRatingsLength, maxRatingsLength)
	ErrInvalidRisk = fmt.Errorf("relative risk must be between %d and %d", minRatingsLength, maxRatingsLength)
)

type BusinessPriorityRatings struct {
	Benefit int
	Cost int
	Penalty int
	Risk int
}

func NewBusinessPriorityRatings(benefit int, cost int, penalty int, risk int) (BusinessPriorityRatings, error) {
	if benefit < 1 || benefit > 9 {
		return BusinessPriorityRatings{}, ErrInvalidBenefit
	}
	if cost < 1 || cost > 9 {
		return BusinessPriorityRatings{}, ErrInvalidCost
	}
	if penalty < 1 || penalty > 9 {
		return BusinessPriorityRatings{}, ErrInvalidPenalty
	}
	if risk < 1 || risk > 9 {
		return BusinessPriorityRatings{}, ErrInvalidRisk
	}

	return BusinessPriorityRatings{Benefit: benefit, Cost: cost, Penalty: penalty, Risk: risk}, nil
}