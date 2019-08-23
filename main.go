package main

func Calculate(product *Product, condition *Condition) (*Offer, error) {
	if product == nil || condition == nil {
		return nil, nil
	}

	return &Offer{}, nil
}
