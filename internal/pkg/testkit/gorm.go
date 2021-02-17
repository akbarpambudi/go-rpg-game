package testkit

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gorm struct {
	suite.Suite
	db *gorm.DB
}

func (g *Gorm) SetupTest()  {
	g.T().Log("Setup Test")
	g.db, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	err := entity.MigrateAllWithGORM(g.db)
	if err != nil {
		g.T().Fatal(err)
	}
}

func (g Gorm) DB() *gorm.DB {
	return g.db
}
