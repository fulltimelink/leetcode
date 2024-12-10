package abstract_factory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{Shoe{
		logo: "AdidasShoe",
		size: 14,
	}}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{Shirt{
		logo: "AdidasShirt",
		size: 14,
	}}
}
