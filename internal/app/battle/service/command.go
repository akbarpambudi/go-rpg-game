package service

type InitiateBattleCommand struct {
	LeftSideCharactersIDs []uint `json:"leftSideCharacterIds"`
	RightSideCharactersIDs []uint `json:"rightSideCharacterIds"`
}
