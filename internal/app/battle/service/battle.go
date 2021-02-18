package service

import "context"

type (
	BattleInitiator interface {
		Initiate(ctx context.Context,command InitiateBattleCommand) (successEvent BattleInitiatedEvent,err error)
	}
)
