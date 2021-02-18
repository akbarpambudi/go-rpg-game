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
//TODO: refactor this shit
func (b BattleInitiator) Initiate(ctx context.Context, command service.InitiateBattleCommand) (successEvent service.BattleInitiatedEvent, err error) {
	var battleCharacters []entity.BattleCharacter
	//Add Left side Character
	leftCharaIds := command.LeftSideCharactersIDs
	var leftCharacters []*entity.Character

	for _,charaId := range leftCharaIds {
		chara,err := b.opts.CharacterRepository.LoadByID(ctx,charaId)
		if err != nil {
			return successEvent,err
		}
		leftCharacters = append(leftCharacters,chara)
	}

	for _,c := range leftCharacters {
		leftBattleCharacter := entity.BattleCharacter{
			CharacterID:     c.ID,
			MaxMP:           c.Stat.BaseHP,
			MP:              c.Stat.BaseMP,
			HP:              c.State.CurrentHP,
			MaxHP:           c.State.CurrentMP,
			Side:            battle.SideLeft,
		}
		battleCharacters = append(battleCharacters,leftBattleCharacter)
	}

	//Add Left side Character
	rightCharaIds := command.RightSideCharactersIDs
	var rightCharacters []*entity.Character
	for _,charaId := range rightCharaIds {
		chara,err := b.opts.CharacterRepository.LoadByID(ctx,charaId)
		if err != nil {
			return successEvent,err
		}
		rightCharacters = append(rightCharacters,chara)
	}

	for _,c := range rightCharacters {
		rightBattleCharacter := entity.BattleCharacter{
			CharacterID:     c.ID,
			MaxMP:           c.Stat.BaseHP,
			MP:              c.Stat.BaseMP,
			HP:              c.State.CurrentHP,
			MaxHP:           c.State.CurrentMP,
			Side:            battle.SideRight,
		}
		battleCharacters = append(battleCharacters,rightBattleCharacter)
	}
	//Create battle session
	session := entity.BattleSession{
		Characters: battleCharacters,
	}

	err = b.opts.BattleRepository.CreateOrUpdate(ctx,&session)
	if err != nil {
		return service.BattleInitiatedEvent{},err
	}

	return service.BattleInitiatedEvent{
		SessionID:             session.ID,
		LeftSideCharacterIDs:  command.LeftSideCharactersIDs,
		RightSideCharacterIDs: command.RightSideCharactersIDs,
	},nil
}
