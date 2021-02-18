package nameutils

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/character"
)

type CharacterNamingRandomized struct {
	repository repository.Character
}

func NewCharacterNamingRandomized(repository repository.Character) *CharacterNamingRandomized {
	return &CharacterNamingRandomized{repository: repository}
}


func (r CharacterNamingRandomized) GetRandomName() string {
	return "random name"
}


func (r CharacterNamingRandomized) GetVeryRandomName() string {
	return "very random name"
}

func (r CharacterNamingRandomized) GenerateCharacterWithRandomName() error {
	return r.repository.CreateOrUpdate(context.Background(),&entity.Character{
		Race:  character.RaceElf,
		Name:  r.GetVeryRandomName(),
	})
}
