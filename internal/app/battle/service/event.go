package service

type BattleInitiatedEvent struct {
	SessionID uint	`json:"sessionId"`
	CharacterIDs []uint `json:"characterIds"`
}
