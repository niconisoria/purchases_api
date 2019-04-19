package models

import (
	"github.com/google/uuid"
)

type Purchase struct {
	ID     string  `json:"id"`
	Image  string  `json:"image"`
	Title  string  `json:"title"`
	Status string  `json:"status"`
	Amount float32 `json:"amount"`
	User   User    `json:"user"`
}

func (p *Purchase) IsValid() bool {
	return p.Amount > 0 && p.User.IsValid()
}

func (p *Purchase) GenerateID() {
	uuid, _ := uuid.NewRandom()
	p.ID = uuid.String()
}
