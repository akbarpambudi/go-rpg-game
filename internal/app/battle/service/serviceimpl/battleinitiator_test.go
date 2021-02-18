package serviceimpl_test

import (
	"context"
	mockrepository "github.com/akbarpambudi/go-rpg-game/internal/app/battle/repository/repositorymock"
	"github.com/akbarpambudi/go-rpg-game/internal/app/battle/service"
	"github.com/akbarpambudi/go-rpg-game/internal/app/battle/service/serviceimpl"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorymock"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/battle"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/character"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type BattleInitiatorTestSuite struct {
	suite.Suite
	sut *serviceimpl.BattleInitiator
	mockBattleRepository *mockrepository.MockBattle
	mockCharacterRepository *repositorymock.MockCharacter
}

func (s *BattleInitiatorTestSuite) SetupSuite()  {

	goMockCtrl := gomock.NewController(s.T())

	s.mockBattleRepository = mockrepository.NewMockBattle(goMockCtrl)
	s.mockCharacterRepository = repositorymock.NewMockCharacter(goMockCtrl)
	s.sut = serviceimpl.NewBattleInitiator(&serviceimpl.BattleInitiatorOptions{
		BattleRepository:    s.mockBattleRepository,
		CharacterRepository: s.mockCharacterRepository,
	})
}

func (s BattleInitiatorTestSuite) TestCallInitiateToInitiateNewBattleSessionShouldBeSuccess()  {
	//arrange
	ctx := context.Background()
	cmd := service.InitiateBattleCommand{
		LeftSideCharactersIDs:  []uint{1,2},
		RightSideCharactersIDs: []uint{4,5},
	}
	s.mockCharacterRepository.EXPECT().LoadByID(gomock.Any(),gomock.Any()).DoAndReturn(func(_ context.Context,id uint) (*entity.Character,error) {

		return &entity.Character{
			Model: gorm.Model{
				ID: id,
			},
			Race:  character.RaceElf,
			Stat:  entity.CharacterStat{
				BaseHP: 10,
				BaseMP: 10,
			},
			State: entity.CharacterState{
				CurrentMP: 10,
				CurrentHP: 10,
			},
			Name:  "Lewis",
		},nil
	}).AnyTimes()
	s.mockBattleRepository.EXPECT().CreateOrUpdate(gomock.Any(),gomock.Any()).Return(nil)
	//act
	ev,err := s.sut.Initiate(ctx,cmd)
	//assert
	s.Run("SuccessEventShouldBeNotZeroValue", func() {
		s.Assert().NotZero(ev)
	})

	s.Run("ErrShouldBeNil", func() {
		s.Assert().NoError(err)
	})
}


func (s BattleInitiatorTestSuite) TestCallInitiateToInitiateNewBattleSessionShouldPassEntityBattleRepository()  {
	//arrange
	ctx := context.Background()
	cmd := service.InitiateBattleCommand{
		LeftSideCharactersIDs:  []uint{1,2},
		RightSideCharactersIDs: []uint{4,5},
	}

	expectedEntity := &entity.BattleSession{
		Characters: []entity.BattleCharacter{
			{
				CharacterID: 1,
				HP: 10,
				MP:10,
				MaxMP: 10,
				MaxHP: 10,
				Side: battle.SideLeft,
			},
			{
				CharacterID: 2,
				HP: 10,
				MP:10,
				MaxMP: 10,
				MaxHP: 10,
				Side: battle.SideLeft,
			},
			{
				CharacterID: 4,
				HP: 10,
				MP:10,
				MaxMP: 10,
				MaxHP: 10,
				Side: battle.SideRight,
			},
			{
				CharacterID: 5,
				HP: 10,
				MP:10,
				MaxMP: 10,
				MaxHP: 10,
				Side: battle.SideRight,
			},
		},
	}

	s.mockCharacterRepository.EXPECT().LoadByID(gomock.Any(),gomock.Any()).DoAndReturn(func(_ context.Context,id uint) (*entity.Character,error) {

		return &entity.Character{
			Model: gorm.Model{
				ID: id,
			},
			Race:  character.RaceElf,
			Stat:  entity.CharacterStat{
				BaseHP: 10,
				BaseMP: 10,
			},
			State: entity.CharacterState{
				CurrentMP: 10,
				CurrentHP: 10,
			},
			Name:  "Lewis",
		},nil
	}).AnyTimes()
	s.mockBattleRepository.EXPECT().CreateOrUpdate(gomock.Eq(ctx),gomock.Eq(expectedEntity)).Return(nil)
	//act
	_,_ = s.sut.Initiate(ctx,cmd)

}

func TestRunBattleInitiatorTestSuite(t *testing.T) {
	suite.Run(t,new(BattleInitiatorTestSuite))
}
