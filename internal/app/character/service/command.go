package service

import "github.com/akbarpambudi/go-rpg-game/internal/app/character/core"

type (
	CharacterCreationCommand struct {
		core.CharacterDTO
	}

	CharacterUpdateCommand struct {
		core.CharacterDTO
	}
)
