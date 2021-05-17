package backlogitem

type BusinessPriority struct {
	Ratings BusinessPriorityRatings
}

func NewBusinessPriority(ratings BusinessPriorityRatings) BusinessPriority {
	return BusinessPriority{Ratings: ratings}
}

func (p BusinessPriority) Priority(total BusinessPriorityTotals) float32 {
	costAndRisk := p.CostPercentage(total) + p.RiskPercentage(total)

	return p.ValuePercentage(total) / costAndRisk
}

func (p BusinessPriority) CostPercentage(total BusinessPriorityTotals) float32 {
	return float32(100 * p.Ratings.Cost / total.TotalCost)
}

func (p BusinessPriority) RiskPercentage(total BusinessPriorityTotals) float32 {
	return float32(p.Ratings.Risk / total.TotalRisk)
}

func (p BusinessPriority) ValuePercentage(total BusinessPriorityTotals) float32 {
	return p.TotalValue() / float32(total.TotalValue)
}

func (p BusinessPriority) TotalValue() float32 {
	return float32(p.Ratings.Benefit + p.Ratings.Penalty)
}
