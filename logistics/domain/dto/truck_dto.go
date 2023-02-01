package dto

type Truck struct {
	Registration string  `db:"registration" json:"registration"`
	Tare         float32 `db:"tare" json:"tare"`
	Capacity     float32 `db:"capacity" json:"capacity"`
	Autonomy     float32 `db:"autonomy" json:"autonomy"`
	Active       bool    `db:"active" json:"active"`
}
