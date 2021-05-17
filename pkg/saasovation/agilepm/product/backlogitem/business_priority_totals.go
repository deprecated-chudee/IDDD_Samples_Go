package backlogitem

type BusinessPriorityTotals struct {
	TotalBenefit int
	TotalCost int
	TotalPenalty int
	TotalRisk int
	TotalValue int
}

func NewBusinessPriorityTotals(totalBenefit int, totalCost int, totalPenalty int, totalRisk int, totalValue int) BusinessPriorityTotals {
	return BusinessPriorityTotals{TotalBenefit: totalBenefit, TotalCost: totalCost, TotalPenalty: totalPenalty, TotalRisk: totalRisk, TotalValue: totalValue}
}