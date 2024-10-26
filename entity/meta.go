package entity

type Meta struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value interface{} `json:"value"`
}
