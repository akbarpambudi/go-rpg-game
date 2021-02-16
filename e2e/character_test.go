package e2e_test

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/delivery"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorygorm"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service/serviceimpl"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/transport"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/gokithelper"
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"testing"
)

type CharacterTestSuite struct {
	suite.Suite
	db         *gorm.DB
	mockServer *apitest.APITest
}

func (s *CharacterTestSuite) SetupTest() {
	router := chi.NewRouter()
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	s.db = db

	err = entity.MigrateAllWithGORM(db)
	if err != nil {
		log.Fatal(err)
	}

	characterRepository := repositorygorm.NewCharacter(&repositorygorm.CharacterOptions{
		DB: db,
	})

	characterManager := serviceimpl.NewCharacterManager(&serviceimpl.CharacterManagerOptions{
		Repository: characterRepository,
	})

	characterHTTPHandler := httptransport.NewServer(
		delivery.MakeCharacterCreationEndpoint(characterManager),
		transport.DecodeCharacterCreationRequestFromHTTP,
		gokithelper.DecodeResponseToHTTP,
	)

	router.Route("/character", func(r chi.Router) {
		r.Post("/", characterHTTPHandler.ServeHTTP)
	})

	s.mockServer = apitest.New().Handler(router)
}

func (s CharacterTestSuite) TestCallCharacterCreationHandlerToCreateNewCharacter() {
	s.mockServer.Post("/character").
		Body(`
			{
				"race":1,
				"name":"Elvin",
				"stat":{
					"baseMP":10,
					"baseHP":20
				},
				"state":{
					"currentHP":10,
					"currentMP":10
				}
			}
		`).
		Expect(s.T()).
		Status(http.StatusOK).
		Assert(jsonpath.Chain().
			NotEqual("$.metadata.id", float64(0)).
			Present("$.name").
			NotEqual("$.race", float64(0)).
			NotEqual("$.stat.metadata.id", float64(0)).
			NotEqual("$.state.metadata.id", float64(0)).
			End(),
		).
		End()
}

func TestRunCharacterTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}
