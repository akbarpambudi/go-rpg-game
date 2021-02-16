package repository

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
)

type Character interface {
	CreateOrUpdate(ctx context.Context, chara *entity.Character) error
	LoadByID(ctx context.Context, id uint) (*entity.Character, error)
	LoadMany(ctx context.Context, predicate predicate.Character) ([]*entity.Character, error)
	RemoveByID(ctx context.Context, id uint) error
}
