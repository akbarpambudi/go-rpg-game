package repository

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
)
//go:generate mockgen -source=battle.go -destination=repositorymock/battle_mock.go -package=mockrepository
type Battle interface {
	CreateOrUpdate(ctx context.Context, battle *entity.BattleSession) error
	LoadByID(ctx context.Context, id uint) (*entity.BattleSession, error)
	LoadMany(ctx context.Context, predicate predicate.BattleSession) ([]*entity.BattleSession, error)
	RemoveByID(ctx context.Context, id uint) error
}
