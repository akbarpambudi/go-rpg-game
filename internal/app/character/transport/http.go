package transport

import (
	"context"
	"encoding/json"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/core"
	"net/http"
)

func DecodeCharacterCreationRequestFromHTTP(_ context.Context, r *http.Request) (interface{}, error) {
	var request core.CharacterDTO
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
