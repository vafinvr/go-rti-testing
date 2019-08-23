package main

const (
	PriceTypeCost     = "COST"
	PriceTypeDiscount = "Discount"
)

type RuleApplicability struct {
	CodeName string `json:"codeName"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Price struct {
	Cost                float64             `json:"cost"`
	PriceType           string              `json:"priceType"`
	RuleApplicabilities []RuleApplicability `json:"ruleApplicabilities"`
}

type Component struct {
	Name   string `json:"name"`
	IsMain bool   `json:"isMain"`
}

type Product struct {
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}

type Condition struct {
	RuleName string `json:"ruleName"`
	Value    string `json:"value"`
}

type Offer struct {
	Product
	TotalCost Price `json:"totalCost"`
}

func (p *Price) GetCostInt64() int64 {
	return int64(p.Cost)
}
