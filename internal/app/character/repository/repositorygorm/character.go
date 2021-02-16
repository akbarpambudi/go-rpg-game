package repositorygorm

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
	"gorm.io/gorm"
)

type CharacterOptions struct {
	DB *gorm.DB
}

type Character struct {
	opts *CharacterOptions
}

func NewCharacter(opts *CharacterOptions) *Character {
	return &Character{opts: opts}
}

func (c Character) CreateOrUpdate(ctx context.Context, chara *entity.Character) error {
	return c.opts.
		DB.
		WithContext(ctx).
		Save(chara).Error
}

func (c Character) LoadByID(ctx context.Context, id uint) (*entity.Character, error) {
	chara := entity.Character{}
	err := c.opts.
		DB.
		WithContext(ctx).
		Joins("Stat").
		Joins("State").
		Find(&chara, id).
		Error
	if err != nil {
		return nil, err
	}

	return &chara, nil
}

func (c Character) LoadMany(ctx context.Context, predicate predicate.Character) ([]*entity.Character, error) {
	var charas []*entity.Character
	structQuery, fields := c.predicateToGormStructCondition(predicate)
	err := c.opts.
		DB.
		WithContext(ctx).
		Where(structQuery, fields).
		Joins("Stat").
		Joins("State").
		Find(&charas).
		Error
	if err != nil {
		return nil, err
	}

	return charas, nil
}

func (c Character) RemoveByID(ctx context.Context, id uint) error {
	return c.opts.DB.WithContext(ctx).Delete(&entity.Character{}, id).Error
}

func (c Character) predicateToGormStructCondition(pred predicate.Character) (queryStruct interface{}, fields []interface{}) {
	var affectedFields []interface{}
	structQuery := pred()
	affectedFieldsInArrayString := structQuery.FilledFields()
	for _, f := range affectedFieldsInArrayString {
		affectedFields = append(affectedFields, f)
	}
	return structQuery, affectedFields
}
