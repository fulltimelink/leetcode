package abstract_factory

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{Shoe{
		logo: "NikeShoe",
		size: 22,
	}}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{Shirt{
		logo: "NikeShirt",
		size: 14,
	}}
}
