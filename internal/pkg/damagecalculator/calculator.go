package damagecalculator

func CalculateDamage(defense int,enemyAttack int) int {
	if enemyAttack < defense {
		return 0
	}
	return enemyAttack - defense
}
