package repositorygorm_test

import (
	"context"
	"errors"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorygorm"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/character"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type CharacterTestSuite struct {
	suite.Suite
	db  *gorm.DB
	sut *repositorygorm.Character
}

func (s *CharacterTestSuite) SetupTest() {
	s.T().Log("Setup Test")
	s.db, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	err := entity.MigrateAllWithGORM(s.db)
	if err != nil {
		s.T().Fatal(err)
	}
	s.sut = repositorygorm.NewCharacter(&repositorygorm.CharacterOptions{
		DB: s.db,
	})
}

func (s *CharacterTestSuite) TestCallCreateOrUpdateToCreateNewEntity() {
	//arrange
	ctx := context.Background()
	testEntity := entity.Character{
		Race: character.RaceElf,
		Stat: entity.CharacterStat{
			BaseAttack:   10,
			BaseDefense:  10,
			BaseHP:       100,
			BaseMP:       100,
			Intelligence: 10,
			Speed:        10,
			Strength:     20,
		},
		State: entity.CharacterState{
			CurrentHP: 100,
			CurrentMP: 100,
		},
		Name: "Mozart Dragon",
	}
	//act
	err := s.sut.CreateOrUpdate(ctx, &testEntity)
	//assert
	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})
	s.Run("ShouldGenerateCharacterID", func() {
		s.Assert().NotZero(testEntity.ID)
	})
	s.Run("ShouldGenerateStatID", func() {
		s.Assert().NotZero(testEntity.Stat.ID)
	})

	s.Run("ShouldGenerateStateID", func() {
		s.Assert().NotZero(testEntity.State.ID)
	})

	s.Run("StatShouldReferenceToCharacterID", func() {
		s.Assert().Equal(testEntity.Stat.CharacterID, testEntity.ID)
	})

	s.Run("StateShouldReferenceToCharacterID", func() {
		s.Assert().Equal(testEntity.State.CharacterID, testEntity.ID)
	})
}

func (s *CharacterTestSuite) TestCallCreateOrUpdateToUpdateExistingEntity() {
	//arrange
	ctx := context.Background()
	testEntity := entity.Character{
		Race: character.RaceElf,
		Stat: entity.CharacterStat{
			BaseAttack:   10,
			BaseDefense:  10,
			BaseHP:       100,
			BaseMP:       100,
			Intelligence: 10,
			Speed:        10,
			Strength:     20,
		},
		State: entity.CharacterState{
			CurrentHP: 100,
			CurrentMP: 100,
		},
		Name: "Mozart Dragon",
	}

	//seed data
	s.db.Create(&testEntity)
	//act
	err := s.sut.CreateOrUpdate(ctx, &testEntity)
	//assert
	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})
	s.Run("ShouldGenerateCharacterID", func() {
		s.Assert().NotZero(testEntity.ID)
	})
	s.Run("ShouldGenerateStatID", func() {
		s.Assert().NotZero(testEntity.Stat.ID)
	})

	s.Run("ShouldGenerateStateID", func() {
		s.Assert().NotZero(testEntity.State.ID)
	})

	s.Run("StatShouldReferenceToCharacterID", func() {
		s.Assert().Equal(testEntity.Stat.CharacterID, testEntity.ID)
	})

	s.Run("StateShouldReferenceToCharacterID", func() {
		s.Assert().Equal(testEntity.State.CharacterID, testEntity.ID)
	})
}

func (s CharacterTestSuite) TestCallLoadByIDToLoadEntityByID() {
	//arrange
	ctx := context.Background()
	testEntity := entity.Character{
		Race: character.RaceElf,
		Stat: entity.CharacterStat{
			BaseAttack:   10,
			BaseDefense:  10,
			BaseHP:       100,
			BaseMP:       100,
			Intelligence: 10,
			Speed:        10,
			Strength:     20,
		},
		State: entity.CharacterState{
			CurrentHP: 100,
			CurrentMP: 100,
		},
		Name: "Mozart Dragon",
	}
	//seed data
	s.db.Create(&testEntity)
	//act
	loadedEntity, err := s.sut.LoadByID(ctx, testEntity.ID)
	//assert
	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})
	s.Run("ShouldLoadTheCharacter", func() {
		s.Assert().Equal(testEntity.ID, loadedEntity.ID)
	})
	s.Run("ShouldPreloadStat", func() {
		s.Assert().Equal(testEntity.Stat.ID, loadedEntity.Stat.ID)
	})
	s.Run("ShouldPreloadState", func() {
		s.Assert().Equal(testEntity.State.ID, loadedEntity.State.ID)
	})
}

func (s CharacterTestSuite) TestCallLoadManyToLoadEntitiesMatchesPredicate() {
	//arrange
	ctx := context.Background()
	testEntities := []entity.Character{
		{
			Race: character.RaceElf,
			Stat: entity.CharacterStat{
				BaseAttack:   10,
				BaseDefense:  10,
				BaseHP:       100,
				BaseMP:       100,
				Intelligence: 10,
				Speed:        10,
				Strength:     20,
			},
			State: entity.CharacterState{
				CurrentHP: 100,
				CurrentMP: 100,
			},
			Name: "Mozart Dragon",
		},
		{
			Race: character.RaceDemon,
			Stat: entity.CharacterStat{
				BaseAttack:   10,
				BaseDefense:  10,
				BaseHP:       100,
				BaseMP:       100,
				Intelligence: 10,
				Speed:        10,
				Strength:     20,
			},
			State: entity.CharacterState{
				CurrentHP: 200,
				CurrentMP: 200,
			},
			Name: "Mira",
		},
	}
	//seed data
	s.db.CreateInBatches(&testEntities, 2)
	//act
	entities, err := s.sut.LoadMany(ctx, character.RaceEQ(character.RaceElf))
	//assert
	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})
	s.Run("ShouldReturnArrayWithLen1", func() {
		s.Assert().Len(entities, 1)
	})
	s.Run("ShouldLoadStatAssociation", func() {
		s.Assert().NotZero(entities[0].Stat)
	})
	s.Run("ShouldLoadStateAssociation", func() {
		s.Assert().NotZero(entities[0].State)
	})
}

func (s CharacterTestSuite) TestCallRemoveByIDToRemoveEntityByID() {
	//arrange
	ctx := context.Background()
	testEntities := []entity.Character{
		{
			Race: character.RaceElf,
			Stat: entity.CharacterStat{
				BaseAttack:   10,
				BaseDefense:  10,
				BaseHP:       100,
				BaseMP:       100,
				Intelligence: 10,
				Speed:        10,
				Strength:     20,
			},
			State: entity.CharacterState{
				CurrentHP: 100,
				CurrentMP: 100,
			},
			Name: "Mozart Dragon",
		},
		{
			Race: character.RaceDemon,
			Stat: entity.CharacterStat{
				BaseAttack:   10,
				BaseDefense:  10,
				BaseHP:       100,
				BaseMP:       100,
				Intelligence: 10,
				Speed:        10,
				Strength:     20,
			},
			State: entity.CharacterState{
				CurrentHP: 200,
				CurrentMP: 200,
			},
			Name: "Mira",
		},
	}
	//seed data
	s.db.CreateInBatches(&testEntities, 2)
	//act
	err := s.sut.RemoveByID(ctx, testEntities[0].ID)
	//assert
	s.Run("ShouldNotReturnAnyError", func() {
		s.Assert().NoError(err)
	})
	s.Run("RemoveTheEntity", func() {
		removedEntity := entity.Character{}
		err := s.db.First(&removedEntity, testEntities[0].ID).Error
		s.Assert().True(errors.Is(err, gorm.ErrRecordNotFound))
	})
}

func TestRunCharacterTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}
