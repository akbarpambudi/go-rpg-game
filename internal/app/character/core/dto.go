package core

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/dto"
	"github.com/devfeel/mapper"
)

func init() {
	_ = mapper.Register(&StatDTO{})
	_ = mapper.Register(&CharacterDTO{})
}

type (
	StatDTO struct {
		dto.EntityMetadata `mapper:"Model" json:"metadata"`
		BaseAttack         int  `mapper:"BaseAttack" json:"baseAttack,omitempty"`
		BaseDefense        int  `mapper:"BaseDefense" json:"baseDefense,omitempty"`
		BaseMP             int  `mapper:"BaseMP" json:"baseMP,omitempty"`
		BaseHP             int  `mapper:"BaseHP" json:"baseHP,omitempty"`
		Strength           int  `mapper:"Strength" json:"strength,omitempty"`
		Speed              int  `mapper:"Speed" json:"speed,omitempty"`
		Intelligence       int  `mapper:"Intelligence" json:"intelligence,omitempty"`
		CharacterID        uint `mapper:"CharacterID" json:"characterID,omitempty"`
	}

	StateDTO struct {
		dto.EntityMetadata `mapper:"Model" json:"metadata"`
		CurrentMP          int  `mapper:"CurrentMP" json:"currentMP,omitempty"`
		CurrentHP          int  `mapper:"CurrentHP" json:"currentHP,omitempty"`
		CharacterID        uint `mapper:"CharacterID" json:"characterID,omitempty"`
	}

	CharacterDTO struct {
		dto.EntityMetadata `mapper:"Model" json:"metadata"`
		Race               entity.CharacterRace `mapper:"Race" json:"race,omitempty"`
		Stat               StatDTO              `mapper:"Stat" json:"stat,omitempty"`
		State              StateDTO             `mapper:"State" json:"state,omitempty"`
		Name               string               `mapper:"Name" json:"name,omitempty"`
	}
)

func (d CharacterDTO) ApplyToEntity(chara *entity.Character) error {
	return mapper.AutoMapper(&d, chara)
}

func MakeCharacterDTOFromEntity(chara entity.Character) (CharacterDTO, error) {
	charaDTO := CharacterDTO{}

	err := mapper.Mapper(&chara, &charaDTO)
	if err != nil {
		return CharacterDTO{}, err
	}
	return charaDTO, nil
}
