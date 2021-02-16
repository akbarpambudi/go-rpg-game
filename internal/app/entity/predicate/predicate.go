package predicate

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/haritsfahreza/libra"
)

type CharacterQuery struct {
	ID   int
	Race entity.CharacterRace
	Name string
}

func (c CharacterQuery) FilledFields() []string {
	var filledFields []string
	empty := CharacterQuery{}
	diffs, err := libra.Compare(context.Background(), empty, c)
	if err != nil {
		panic(err)
	}

	for _, diff := range diffs {
		filledFields = append(filledFields, diff.Field)
	}

	return filledFields
}

type Character func() CharacterQuery
