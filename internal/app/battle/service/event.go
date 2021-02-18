package service

type BattleInitiatedEvent struct {
	SessionID uint	`json:"sessionId"`
	LeftSideCharacterIDs []uint `json:"leftSideCharacterIds"`
	RightSideCharacterIDs []uint `json:"rightSideCharacterIds"`
}
