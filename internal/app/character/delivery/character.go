package delivery

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/core"
	"github.com/akbarpambudi/go-rpg-game/internal/app/character/service"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/gokithelper"
	"github.com/go-kit/kit/endpoint"
)

func MakeCharacterCreationEndpoint(charaManager service.CharacterManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(core.CharacterDTO)
		successEvent, err := charaManager.Create(ctx, service.CharacterCreationCommand{
			CharacterDTO: req,
		})

		if err != nil {
			return gokithelper.NewFailedResponse(err), nil
		}

		return gokithelper.NewSuccessResponse(successEvent), nil
	}
}
