package handlers

import (
	"math/rand"
)

type Law struct {
	Name       string `json:"name,omitempty"`
	Definition string `json:"definition,omitempty"`
}

var HackerLaws = []Law{
	{
		Name:       "Amdahl's Law",
		Definition: "Amdahl's Law is a formula which shows the potential speedup of a computational task which can be achieved by increasing the resources of a system.",
	},
	{
		Name:       "Conway's Law",
		Definition: "This law suggests that the technical boundaries of a system will reflect the structure of the organisation.",
	},
	{
		Name:       "Gall's Law",
		Definition: "A complex system that works is invariably found to have evolved from a simple system that worked.",
	},
}

func GetRandomLaw() Law {
	return HackerLaws[rand.Intn(len(HackerLaws))]
}
