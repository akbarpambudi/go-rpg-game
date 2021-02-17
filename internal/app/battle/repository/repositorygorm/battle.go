package repositorygorm

import (
	"context"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity"
	"github.com/akbarpambudi/go-rpg-game/internal/app/entity/predicate"
	"gorm.io/gorm"
)

type BattleOptions struct {
	DB *gorm.DB
}

type Battle struct {
	opts *BattleOptions
}

func NewBattle(opts *BattleOptions) *Battle {
	return &Battle{opts: opts}
}

func (b Battle) CreateOrUpdate(ctx context.Context, battle *entity.BattleSession) error {
	return b.opts.
		DB.
		WithContext(ctx).
		Save(battle).Error
}

func (b Battle) LoadByID(ctx context.Context, id uint) (*entity.BattleSession, error) {

	battleSess := entity.BattleSession{}

	err := b.opts.
		DB.
		WithContext(ctx).
		Joins("Characters").
		Find(&battleSess, id).
		Error
	if err != nil {
		return nil, err
	}

	return &battleSess,nil
}

func (b Battle) LoadMany(ctx context.Context, predicate predicate.BattleSession) ([]*entity.BattleSession, error) {
	var sessions []*entity.BattleSession
	structQuery, fields := b.predicateToGormStructCondition(predicate)
	err := b.opts.
		DB.
		WithContext(ctx).
		Where(structQuery, fields).
		Joins("Stat").
		Joins("State").
		Find(&sessions).
		Error

	if err != nil {
		return nil,err
	}

	return sessions,nil
}

func (b Battle) RemoveByID(ctx context.Context, id uint) error {
	return b.opts.DB.WithContext(ctx).Delete(&entity.BattleSession{}, id).Error
}

func (b Battle) predicateToGormStructCondition(pred predicate.BattleSession) (queryStruct interface{}, fields []interface{}) {
	var affectedFields []interface{}
	structQuery := pred()
	affectedFieldsInArrayString := structQuery.FilledFields()
	for _, f := range affectedFieldsInArrayString {
		affectedFields = append(affectedFields, f)
	}
	return structQuery, affectedFields
}
