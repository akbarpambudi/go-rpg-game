package service

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/core"
)

type (
	CharacterCreatedEvent struct {
		core.CharacterDTO
	}

	CharacterUpdatedEvent struct {
		core.CharacterDTO
	}
)
