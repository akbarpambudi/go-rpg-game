package damagecalculator_test

import (
	"fmt"
	"github.com/akbarpambudi/go-rpg-game/internal/pkg/damagecalculator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

func (s TestSuite) TestCallCalculateDamageShouldReturnDamageCalculation() {
	tt := []struct {
		attack         int
		defense        int
		expectedDamage int
	}{
		{
			attack:         10,
			defense:        0,
			expectedDamage: 10,
		},
		{
			attack:         100,
			defense:        120,
			expectedDamage: 0,
		},
		{
			attack:         0,
			defense:        120,
			expectedDamage: 0,
		},
	}

	for i, r := range tt {
		s.Run(fmt.Sprintf("Case#%v",i), func() {
			got := damagecalculator.CalculateDamage(r.defense,r.attack)
			s.Assert().Equal(r.expectedDamage,got)
		})
	}
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
