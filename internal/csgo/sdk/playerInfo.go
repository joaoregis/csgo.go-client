package sdk

type PlayerInfo_T struct {
	PAD_000      [16]byte
	SzPlayerName [32]byte
}

type CSTeam uint

const (
	CSTeam_None CSTeam = iota
	CSTeam_Spectators
	CSTeam_TT
	CSTeam_CT
)
