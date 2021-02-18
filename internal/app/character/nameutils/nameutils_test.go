package nameutils_test

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/nameutils"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorymock"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/character"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CharacterNamingRandomizedTestSuite struct {
	suite.Suite
	sut *nameutils.CharacterNamingRandomized
	repositoryMock *repositorymock.MockCharacter
}

func (s *CharacterNamingRandomizedTestSuite) SetupSuite()  {
	gomockCtrl := gomock.NewController(s.T())
	repository := repositorymock.NewMockCharacter(gomockCtrl)
	s.sut = nameutils.NewCharacterNamingRandomized(repository)
	s.repositoryMock = repository
}

func (s CharacterNamingRandomizedTestSuite) TearDownSuite()  {
	s.T().Log("Delete all data")
}

func (s CharacterNamingRandomizedTestSuite) TestCallGetRandomName() {
	want := "random name"
	got := s.sut.GetRandomName()
	s.Assert().Equal(want,got)
}

func (s CharacterNamingRandomizedTestSuite) TestCallGetVeryRandomName(){
	t := s.T()
	want := "very random name"
	got := s.sut.GetVeryRandomName()
	assert.Equal(t,want,got)
}

func (s CharacterNamingRandomizedTestSuite) TestCallGenerateCharacterWithRandomName()  {
	expectedEntity := &entity.Character{
		Name: "very random name",
		Race: character.RaceElf,
	}
	s.repositoryMock.EXPECT().CreateOrUpdate(gomock.Any(),gomock.Eq(expectedEntity)).Return(nil)
	err := s.sut.GenerateCharacterWithRandomName()
	s.Assert().NoError(err)
}

func TestRunCharacterNamingRandomizedTestSuite(t *testing.T) {
	suite.Run(t,new(CharacterNamingRandomizedTestSuite))
}
