package repositorygorm_test

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/battle/repository/repositorygorm"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/battle"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/testkit"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BattleTestSuite struct {
	testkit.Gorm
	sut *repositorygorm.Battle
}

func (s *BattleTestSuite) SetupTest() {
	s.Gorm.SetupTest()
	s.sut = repositorygorm.NewBattle(&repositorygorm.BattleOptions{
		DB: s.DB(),
	})
}

func (s BattleTestSuite) TestCreateOrUpdate() {
	ctx := context.Background()
	testEntity := entity.BattleSession{
		Characters: []entity.BattleCharacter{
			{
				CharacterID: 1,
				MaxHP:       10,
				MaxMP:       10,
				HP:          10,
				MP:          10,
				Side:        battle.SideLeft,
			},
			{
				CharacterID: 2,
				MaxHP:       10,
				MaxMP:       10,
				HP:          10,
				MP:          10,
				Side:        battle.SideRight,
			},
		},
	}

	err := s.sut.CreateOrUpdate(ctx, &testEntity)

	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})

	s.Run("ShouldCreateAllBattleCharacter", func() {
		var size int64
		s.DB().Model(&entity.BattleCharacter{}).Where(&entity.BattleCharacter{BattleSessionID: testEntity.ID}).Count(&size)
		s.Assert().Equal(size, int64(2))
	})
}

func TestRunBattleTestSuite(t *testing.T) {
	suite.Run(t, new(BattleTestSuite))
}
