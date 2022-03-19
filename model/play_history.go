package model

import "time"

type PlayHistory struct {
	listCard  []Card
	winner    PlayerDetail
	player    []PlayerDetail
	CreatedAt time.Time
	UpdatedAt time.Time
}
