package api

import "fmt"

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func (p *Pokemon) String() string {
	return fmt.Sprintf("%v is %v cm tall and weights %v Kgs", p.Name, p.Height, p.Weight)
}
