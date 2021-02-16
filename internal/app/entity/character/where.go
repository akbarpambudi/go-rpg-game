package character

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
)

func RaceEQ(race entity.CharacterRace) predicate.Character {
	return func() predicate.CharacterQuery {
		return predicate.CharacterQuery{
			Race: race,
		}
	}
}
