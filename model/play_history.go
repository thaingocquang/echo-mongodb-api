package model

import "time"

type PlayHistory struct {
	winner    PlayerDetail
	player    []PlayerDetail
	CreatedAt time.Time
	UpdatedAt time.Time
}
