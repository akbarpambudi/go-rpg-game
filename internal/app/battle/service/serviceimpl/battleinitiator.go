package serviceimpl

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/battle/repository"
	"github.com/akbarpambudi/go-rpg-game/internal/app/battle/service"
	charaRepository "github.com/akbarpambudi/go-rpg-game/internal/app/character/repository"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/battle"
)

type BattleInitiatorOptions struct {
	BattleRepository    repository.Battle
	CharacterRepository charaRepository.Character
}

type BattleInitiator struct {
	opts *BattleInitiatorOptions
}

func NewBattleInitiator(opts *BattleInitiatorOptions) *BattleInitiator {
	return &BattleInitiator{opts: opts}
}
//Note: this code is bad code, this is for codescene bad code scanning testing only

func (b BattleInitiator) Initiate(ctx context.Context, command service.InitiateBattleCommand) (successEvent service.BattleInitiatedEvent, err error) {
	var battleCharacters []entity.BattleCharacter
	//Add Left side Character
	leftSideCharacters,err := b.findCharactersWhereIdIn(ctx,command.LeftSideCharactersIDs)
	if err != nil {
		return successEvent,err
	}
	leftSideBattleCharacters := b.mapCharactersToBattleCharacters(ctx,leftSideCharacters,battle.SideLeft)
	battleCharacters = append(battleCharacters,leftSideBattleCharacters...)
	//Add Right side Character
	rightSideCharacters,err := b.findCharactersWhereIdIn(ctx,command.RightSideCharactersIDs)
	if err != nil {
		return successEvent,err
	}
	rightSideBattleCharacters := b.mapCharactersToBattleCharacters(ctx,rightSideCharacters,battle.SideRight)
	battleCharacters = append(battleCharacters,rightSideBattleCharacters...)
	//Create battle session
	session := entity.BattleSession{
		Characters: battleCharacters,
	}

	err = b.opts.BattleRepository.CreateOrUpdate(ctx, &session)
	if err != nil {
		return service.BattleInitiatedEvent{}, err
	}

	return service.BattleInitiatedEvent{
		SessionID:             session.ID,
		LeftSideCharacterIDs:  command.LeftSideCharactersIDs,
		RightSideCharacterIDs: command.RightSideCharactersIDs,
	}, nil
}

func (b BattleInitiator) mapCharactersToBattleCharacters(ctx context.Context, characters []*entity.Character,side entity.BattleSide) (result []entity.BattleCharacter) {
	for _, c := range characters {
		b := entity.BattleCharacter{
			CharacterID: c.ID,
			MaxMP:       c.Stat.BaseHP,
			MP:          c.Stat.BaseMP,
			HP:          c.State.CurrentHP,
			MaxHP:       c.State.CurrentMP,
			Side:        side,
		}
		result = append(result, b)
	}
	return result
}
//TODO: refactor this shit, move to repository or somewhere else has this responsibility
func (b BattleInitiator) findCharactersWhereIdIn(ctx context.Context,characterIds []uint) ([]*entity.Character,error) {
	var characters []*entity.Character

	for _, charaId := range characterIds {
		chara, err := b.opts.CharacterRepository.LoadByID(ctx, charaId)
		if err != nil {
			return nil,err
		}
		characters = append(characters,chara)
	}
	return characters,nil
}
