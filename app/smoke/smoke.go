package smoke

import (
	dota2gsi "github.com/go-programmer/Dota2GSI"
)

//SmokeStatus contains "visibility"/"hidden" style property
//visibility.
type SmokeStatus struct {
	StPlayer0 string
	StPlayer1 string
	StPlayer2 string
	StPlayer3 string
	StPlayer4 string
	StPlayer5 string
	StPlayer6 string
	StPlayer7 string
	StPlayer8 string
	StPlayer9 string
}

//PlayersView all players implementation.
type PlayersView interface {
	Player0() string
	Player1() string
	Player2() string
	Player3() string
	Player4() string
	Player5() string
	Player6() string
	Player7() string
	Player8() string
	Player9() string

	GetPlayersStatus() SmokeStatus

	GetRoshanTime() string

	GetGSIData() *dota2gsi.Dota2GSIData
}

//PlayerChecker data
type PlayerChecker struct {
	GSI dota2gsi.GSIData
}

//GetGSIData share GSIData with another modules.
func (p *PlayerChecker) GetGSIData() *dota2gsi.Dota2GSIData{
	return p.GSI.Data.GetDota2GSIType()
}

//GetRoshanTime return roshan time in secoonds from dota 2 data.
func (p *PlayerChecker) GetRoshanTime() string{
	return p.GSI.Data.GetRoshanTime()
}

//GetPlayersStatus connects together Player's tread. 
func (p *PlayerChecker) GetPlayersStatus() SmokeStatus {
	var SmokeData SmokeStatus
	SmokeData = SmokeStatus{
		StPlayer0: p.Player0(),
		StPlayer1: p.Player1(),
		StPlayer2: p.Player2(),
		StPlayer3: p.Player3(),
		StPlayer4: p.Player4(),
		StPlayer5: p.Player5(),
		StPlayer6: p.Player6(),
		StPlayer7: p.Player7(),
		StPlayer8: p.Player8(),
		StPlayer9: p.Player9(),
	}
	return SmokeData
}

//Player0 launch payer 4 tread.
func (p *PlayerChecker) Player0() string {
	p.GSI.Data.CheckAegisPlayer0()
	if p.GSI.Data.GetSmokePlayer0() {
		return "visible"
	}
	return "hidden"
}

//Player1 launch payer 4 tread.
func (p *PlayerChecker) Player1() string {
	p.GSI.Data.CheckAegisPlayer1()
	if p.GSI.Data.GetSmokePlayer1() {
		return "visible"
	}
	return "hidden"
}

//Player2 launch payer 4 tread.
func (p *PlayerChecker) Player2() string {
	p.GSI.Data.CheckAegisPlayer2()
	if p.GSI.Data.GetSmokePlayer2() {
		return "visible"
	}
	return "hidden"
}

//Player3 launch payer 4 tread.
func (p *PlayerChecker) Player3() string {
	p.GSI.Data.CheckAegisPlayer3()
	if p.GSI.Data.GetSmokePlayer3() {
		return "visible"
	}
	return "hidden"
}

//Player4 launch payer 4 tread.
func (p *PlayerChecker) Player4() string {
	p.GSI.Data.CheckAegisPlayer4()
	if p.GSI.Data.GetSmokePlayer4() {
		return "visible"
	}
	return "hidden"
}

//Player5 launch payer 4 tread.
func (p *PlayerChecker) Player5() string {
	p.GSI.Data.CheckAegisPlayer5()
	if p.GSI.Data.GetSmokePlayer5() {
		return "visible"
	}
	return "hidden"
}

//Player6 launch payer 4 tread.
func (p *PlayerChecker) Player6() string {
	p.GSI.Data.CheckAegisPlayer6()
	if p.GSI.Data.GetSmokePlayer6() {
		return "visible"
	}
	return "hidden"
}

//Player7 launch payer 4 tread.
func (p *PlayerChecker) Player7() string {
	p.GSI.Data.CheckAegisPlayer7()
	if p.GSI.Data.GetSmokePlayer7() {
		return "visible"
	}
	return "hidden"
}

//Player8 launch payer 4 tread.
func (p *PlayerChecker) Player8() string {
	p.GSI.Data.CheckAegisPlayer8()
	if p.GSI.Data.GetSmokePlayer8() {
		return "visible"
	}
	return "hidden"
}

//Player9 launch payer 4 tread.
func (p *PlayerChecker) Player9() string {
	p.GSI.Data.CheckAegisPlayer9()
	if p.GSI.Data.GetSmokePlayer9() {
		return "visible"
	}
	return "hidden"
}
