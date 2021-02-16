package entity

import "gorm.io/gorm"

func MigrateAllWithGORM(db *gorm.DB) error {
	err := db.AutoMigrate(
		&CharacterState{},
		&CharacterStat{},
		&Character{},
	)
	if err != nil {
		return err
	}
	return nil
}
