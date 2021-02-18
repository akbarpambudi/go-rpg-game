package entity

import "gorm.io/gorm"

type BattleSide int

type BattleCharacter struct {
	gorm.Model
	CharacterID     uint
	BattleSessionID uint
	MaxMP           int
	MP              int
	HP              int
	MaxHP           int
	Side            BattleSide
}

type BattleSession struct {
	gorm.Model
	Characters []BattleCharacter
}
