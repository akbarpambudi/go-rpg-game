package serviceimpl

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/core"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
)

type CharacterManagerOptions struct {
	Repository repository.Character
}

type CharacterManager struct {
	opts *CharacterManagerOptions
}

func NewCharacterManager(opts *CharacterManagerOptions) *CharacterManager {
	return &CharacterManager{opts: opts}
}

func (c CharacterManager) Create(ctx context.Context, command service.CharacterCreationCommand) (successEvent service.CharacterCreatedEvent, err error) {
	chara := entity.Character{}
	applyingEntityErr := command.ApplyToEntity(&chara)
	if applyingEntityErr != nil {
		return service.CharacterCreatedEvent{}, applyingEntityErr
	}
	creationErr := c.opts.Repository.CreateOrUpdate(ctx, &chara)
	if creationErr != nil {
		return service.CharacterCreatedEvent{}, creationErr
	}

	updatedCharaDTO, characterMakingErr := core.MakeCharacterDTOFromEntity(chara)
	if characterMakingErr != nil {
		return service.CharacterCreatedEvent{}, characterMakingErr
	}
	successEvent.CharacterDTO = updatedCharaDTO
	return successEvent, nil
}

func (c CharacterManager) Update(ctx context.Context, command service.CharacterUpdateCommand) (successEvent service.CharacterUpdatedEvent, err error) {
	chara := entity.Character{}
	applyingEntityErr := command.ApplyToEntity(&chara)
	if applyingEntityErr != nil {
		return service.CharacterUpdatedEvent{}, applyingEntityErr
	}

	updatingErr := c.opts.Repository.CreateOrUpdate(ctx, &chara)
	if updatingErr != nil {
		return service.CharacterUpdatedEvent{}, updatingErr
	}

	updatedCharaDTO, characterMakingErr := core.MakeCharacterDTOFromEntity(chara)
	if characterMakingErr != nil {
		return service.CharacterUpdatedEvent{}, characterMakingErr
	}
	successEvent.CharacterDTO = updatedCharaDTO
	return successEvent, nil
}
