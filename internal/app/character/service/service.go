package service

import "context"

type CharacterManager interface {
	Create(ctx context.Context, command CharacterCreationCommand) (successEvent CharacterCreatedEvent, err error)
	Update(ctx context.Context, command CharacterUpdateCommand) (successEvent CharacterUpdatedEvent, err error)
}
