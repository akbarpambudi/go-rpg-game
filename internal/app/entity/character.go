package entity

import (
	"gorm.io/gorm"
)

type (
	CharacterRace int

	CharacterStat struct {
		gorm.Model
		BaseAttack   int
		BaseDefense  int
		BaseMP       int
		BaseHP       int
		Strength     int
		Speed        int
		Intelligence int
		CharacterID  uint
	}

	CharacterState struct {
		gorm.Model
		CurrentMP   int
		CurrentHP   int
		CharacterID uint
	}

	Character struct {
		gorm.Model
		Race  CharacterRace
		Stat  CharacterStat
		State CharacterState
		Name  string
	}
)
