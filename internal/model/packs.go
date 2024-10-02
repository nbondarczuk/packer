package model

type Pack struct {
	Value int `json: "value"`
}

type Packs struct {
	packs []Packs `json: "packs"`
}
