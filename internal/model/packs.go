package model

type Packs struct {
	Value   int
	Buckets []int `json: "buckets"`
}
