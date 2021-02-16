package main

import (
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/delivery"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/repository/repositorygorm"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service/serviceimpl"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/transport"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/gokithelper"
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	db, err := gorm.Open(sqlite.Open("character.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
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

	log.Println("Serving character service on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
