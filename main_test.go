package main

import "testing"

const product = `{
    "name": "Игровой",
    "components": [
        {
            "isMain": true,
            "name": "Интернет",
            "prices": [
                {
                    "cost": 100,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "10"
                        }
                    ]
                },
                {
                    "cost": 150,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "15"
                        }
                    ]
                },
                {
                    "cost": 500,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "xpon"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "100"
                        }
                    ]
                },
                {
                    "cost": 900,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "xpon"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "200"
                        }
                    ]
                },
                {
                    "cost": 200,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "30"
                        }
                    ]
                },
                {
                    "cost": 400,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "50"
                        }
                    ]
                },
                {
                    "cost": 600,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "200"
                        }
                    ]
                },
                {
                    "cost": 10,
                    "priceType": "DISCOUNT",
                    "ruleApplicabilities": [
                        {
                            "codeName": "internetSpeed",
                            "operator": "GTE",
                            "value": "50"
                        }
                    ]
                }
            ]
        },
        {
            "name": "ADSL Модем",
            "prices": [
                {
                    "cost": 300,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        }
                    ]
                }
            ]
        }
    ]
}`

func TestCalculateNil(t *testing.T) {
	r, err := Calculate(nil, nil)
	if err != nil {
		t.Error("Error calculating", err)
	}
	if r != nil {
		t.Error(r)
	}
}

func TestCalculateNotNil(t *testing.T) {
	r, err := Calculate(&Product{}, &Condition{})
	if err != nil {
		t.Error("Error calculating", err)
	}
	if r == nil {
		t.Error(r)
	}
}
