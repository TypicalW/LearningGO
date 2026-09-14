package main

import (
	"time"
)

type Cigarette struct {
	Name     string
	Brand    string
	SmokedAt time.Time
}

type Cigarettes []Cigarette

func (ciggs *Cigarettes) add(name string, brand string) {
	cigarette := Cigarette{
		Name:     name,
		Brand:    brand,
		SmokedAt: time.Now(),
	}

	*ciggs = append(*ciggs, cigarette)
}
