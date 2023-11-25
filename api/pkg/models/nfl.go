package models

type TeamCode string

const (
	ARI TeamCode = "ARI"
	ATL TeamCode = "ATL"
	BAL TeamCode = "BAL"
	BUF TeamCode = "BUF"
	CAR TeamCode = "CAR"
	CHI TeamCode = "CHI"
	CIN TeamCode = "CIN"
	CLE TeamCode = "CLE"
	DAL TeamCode = "DAL"
	DEN TeamCode = "DEN"
	DET TeamCode = "DET"
	GB  TeamCode = "GB"
	HOU TeamCode = "HOU"
	IND TeamCode = "IND"
	JAX TeamCode = "JAX"
	KC  TeamCode = "KC"
	LAC TeamCode = "LAC"
	LAR TeamCode = "LAR"
	LV  TeamCode = "LV"
	MIA TeamCode = "MIA"
	MIN TeamCode = "MIN"
	NE  TeamCode = "NE"
	NO  TeamCode = "NO"
	NYG TeamCode = "NYG"
	NYJ TeamCode = "NYJ"
	PHI TeamCode = "PHI"
	PIT TeamCode = "PIT"
	SEA TeamCode = "SEA"
	SF  TeamCode = "SF"
	TB  TeamCode = "TB"
	TEN TeamCode = "TEN"
	WAS TeamCode = "WAS"
)

func GetAllTeamCodes() []TeamCode {
	return []TeamCode{
		ARI,
		ATL,
		BAL,
		BUF,
		CAR,
		CHI,
		CIN,
		CLE,
		DAL,
		DEN,
		DET,
		GB,
		HOU,
		IND,
		JAX,
		KC,
		LAC,
		LAR,
		LV,
		MIA,
		MIN,
		NE,
		NO,
		NYG,
		NYJ,
		PHI,
		PIT,
		SEA,
		SF,
		TB,
		TEN,
		WAS,
	}
}

type NFLTeam struct {
	Code TeamCode `json:"code" gorm:"primaryKey"`
}

type Position string

const (
	QB  Position = "QB"
	RB  Position = "RB"
	WR  Position = "WR"
	TE  Position = "TE"
	K   Position = "K"
	DEF Position = "DEF"
)
