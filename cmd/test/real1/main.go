package main

import "fmt"

type ReptileType int

const FIRE_DRAGON ReptileType = iota + 1

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
	hatched       bool
	reptileType   ReptileType
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.hatched {
		return nil
	}

	egg.hatched = true
	switch egg.reptileType {
	case FIRE_DRAGON:
		return &FireDragon{}
	default:
		return nil
	}
}

type FireDragon struct {
}

func (f FireDragon) Lay() ReptileEgg {
	return ReptileEgg{
		reptileType: FIRE_DRAGON,
		CreateReptile: func() Reptile {
			return &FireDragon{}
		},
	}
}

func main() {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch()
	fmt.Println(childDragon)

}
