package entity

import "gorm.io/gorm"

func MigrateAllWithGORM(db *gorm.DB) error {
	err := db.AutoMigrate(
		&CharacterState{},
		&CharacterStat{},
		&Character{},
		&BattleCharacter{},
		&BattleSession{},
	)
	if err != nil {
		return err
	}
	return nil
}
