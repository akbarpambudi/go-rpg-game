package service

type InitiateBattleCommand struct {
	CharacterIDs []uint `json:"characterIds"`
}
