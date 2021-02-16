package serviceimpl_test

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/core"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorymock"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service/serviceimpl"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/character"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CharacterTestSuite struct {
	suite.Suite
	mockCharacterRepository *repositorymock.MockCharacter
	sut                     *serviceimpl.CharacterManager
}

func (s *CharacterTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	s.mockCharacterRepository = repositorymock.NewMockCharacter(mockCtrl)
	s.sut = serviceimpl.NewCharacterManager(&serviceimpl.CharacterManagerOptions{
		Repository: s.mockCharacterRepository,
	})
}

func (s *CharacterTestSuite) TestCallCreateToCreateCharacterShouldBeSuccess() {
	//arrange
	ctx := context.Background()
	command := service.CharacterCreationCommand{
		CharacterDTO: core.CharacterDTO{
			State: core.StateDTO{
				CurrentMP: 100,
				CurrentHP: 100,
			},
			Stat: core.StatDTO{
				BaseMP:      100,
				BaseHP:      100,
				BaseAttack:  200,
				BaseDefense: 200,
			},
			Name: "Norman",
			Race: character.RaceElf,
		},
	}
	expectedEntity := entity.Character{
		Race: command.Race,
		Stat: entity.CharacterStat{
			BaseAttack:   command.Stat.BaseAttack,
			BaseDefense:  command.Stat.BaseDefense,
			BaseMP:       command.Stat.BaseMP,
			BaseHP:       command.Stat.BaseHP,
			Strength:     command.Stat.Strength,
			Speed:        command.Stat.Speed,
			Intelligence: command.Stat.Intelligence,
		},
		State: entity.CharacterState{
			CurrentMP:   command.State.CurrentMP,
			CurrentHP:   command.State.CurrentHP,
		},
		Name:  command.Name,
	}
	s.mockCharacterRepository.EXPECT().CreateOrUpdate(gomock.Eq(ctx), gomock.Eq(&expectedEntity)).Return(nil)
	//act
	successEvent, err := s.sut.Create(ctx, command)
	//assert
	s.Assert().NoError(err)
	s.Assert().NotZero(successEvent)
}

func TestRunCharacterTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}
