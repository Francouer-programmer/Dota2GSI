package dota2gsi

import (
	"fmt"
	"time"
)

//GSIView methoods implementation
type GSIView interface {
	CheckAegisPlayer0()
	CheckAegisPlayer1()
	CheckAegisPlayer2()
	CheckAegisPlayer3()
	CheckAegisPlayer4()
	CheckAegisPlayer5()
	CheckAegisPlayer6()
	CheckAegisPlayer7()
	CheckAegisPlayer8()
	CheckAegisPlayer9()

	GetSmokePlayer0() bool
	GetSmokePlayer1() bool
	GetSmokePlayer2() bool
	GetSmokePlayer3() bool
	GetSmokePlayer4() bool
	GetSmokePlayer5() bool
	GetSmokePlayer6() bool
	GetSmokePlayer7() bool
	GetSmokePlayer8() bool
	GetSmokePlayer9() bool

	GetRoshanTime() string
}



//GetRoshanTime return respound roshans time in seconds.
func (dota *Dota2GSIData) GetRoshanTime() string{
	if dota.Map.RoshanState != "alive" {
		time.Sleep(2 * time.Second)
		fmt.Println("Roshan state")
		fmt.Println(dota.Map.RoshanStateEndSeconds)
		return fmt.Sprintf("Roshan state => %d", dota.Map.RoshanStateEndSeconds)
	}
	return ""
}

//CheckAegisPlayer0 launch search in slots for aegis in player0.
func (dota *Dota2GSIData) CheckAegisPlayer0() {

	if dota.Items.Team2.Player0.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot0")
	}
	if dota.Items.Team2.Player0.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot1")
	}
	if dota.Items.Team2.Player0.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot2")
	}
	if dota.Items.Team2.Player0.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot3")
	}
	if dota.Items.Team2.Player0.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot4")
	}
	if dota.Items.Team2.Player0.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot5")
	}
	if dota.Items.Team2.Player0.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player0 Slot6")
	}
}

//CheckAegisPlayer1 launch search in slots for aegis in player1.
func (dota *Dota2GSIData) CheckAegisPlayer1() {

	if dota.Items.Team2.Player1.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot0")
	}
	if dota.Items.Team2.Player1.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot1")
	}
	if dota.Items.Team2.Player1.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot2")
	}
	if dota.Items.Team2.Player1.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot3")
	}
	if dota.Items.Team2.Player1.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot4")
	}
	if dota.Items.Team2.Player1.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot5")
	}
	if dota.Items.Team2.Player1.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player1 Slot6")
	}
}

//CheckAegisPlayer2 launch search in slots for aegis in player2.
func (dota *Dota2GSIData) CheckAegisPlayer2() {

	if dota.Items.Team2.Player2.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot0")
	}
	if dota.Items.Team2.Player2.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot1")
	}
	if dota.Items.Team2.Player2.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot2")
	}
	if dota.Items.Team2.Player2.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot3")
	}
	if dota.Items.Team2.Player2.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot4")
	}
	if dota.Items.Team2.Player2.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot5")
	}
	if dota.Items.Team2.Player2.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player2 Slot6")
	}
}

//CheckAegisPlayer3 launch search in slots for aegis in player3.
func (dota *Dota2GSIData) CheckAegisPlayer3() {

	if dota.Items.Team2.Player3.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot0")
	}
	if dota.Items.Team2.Player3.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot1")
	}
	if dota.Items.Team2.Player3.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot2")
	}
	if dota.Items.Team2.Player3.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot3")
	}
	if dota.Items.Team2.Player3.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot4")
	}
	if dota.Items.Team2.Player3.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot5")
	}
	if dota.Items.Team2.Player3.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player3 Slot6")
	}
}

//CheckAegisPlayer4 launch search in slots for aegis in player4.
func (dota *Dota2GSIData) CheckAegisPlayer4() {

	if dota.Items.Team2.Player4.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot0")
	}

	if dota.Items.Team2.Player4.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot1")
	}

	if dota.Items.Team2.Player4.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot2")
	}

	if dota.Items.Team2.Player4.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot3")
	}

	if dota.Items.Team2.Player4.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot4")
	}

	if dota.Items.Team2.Player4.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot5")
	}

	if dota.Items.Team2.Player4.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player4 Slot6")
	}
}

//CheckAegisPlayer5 launch search in slots for aegis in player5.
func (dota *Dota2GSIData) CheckAegisPlayer5() {

	if dota.Items.Team3.Player5.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot0")
	}
	if dota.Items.Team3.Player5.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot1")
	}
	if dota.Items.Team3.Player5.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot2")
	}
	if dota.Items.Team3.Player5.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot3")
	}
	if dota.Items.Team3.Player5.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot4")
	}
	if dota.Items.Team3.Player5.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot5")
	}
	if dota.Items.Team3.Player5.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot6")
	}
}

//CheckAegisPlayer6 launch search in slots for aegis in player6.
func (dota *Dota2GSIData) CheckAegisPlayer6() {

	if dota.Items.Team3.Player6.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot0")
	}
	if dota.Items.Team3.Player6.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot1")
	}
	if dota.Items.Team3.Player6.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot2")
	}
	if dota.Items.Team3.Player6.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot3")
	}
	if dota.Items.Team3.Player6.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot4")
	}
	if dota.Items.Team3.Player6.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot5")
	}
	if dota.Items.Team3.Player6.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player6 Slot6")
	}
}

//CheckAegisPlayer7 launch search in slots for aegis in player7.
func (dota *Dota2GSIData) CheckAegisPlayer7() {

	if dota.Items.Team3.Player7.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot0")
	}
	if dota.Items.Team3.Player7.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot1")
	}
	if dota.Items.Team3.Player7.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot2")
	}
	if dota.Items.Team3.Player7.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot3")
	}
	if dota.Items.Team3.Player7.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot4")
	}
	if dota.Items.Team3.Player7.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot5")
	}
	if dota.Items.Team3.Player7.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player7 Slot6")
	}
}

//CheckAegisPlayer8 launch search in slots for aegis in player8.
func (dota *Dota2GSIData) CheckAegisPlayer8() {
	if dota.Items.Team3.Player8.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot0")
	}
	if dota.Items.Team3.Player8.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot1")
	}
	if dota.Items.Team3.Player8.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot2")
	}
	if dota.Items.Team3.Player8.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot3")
	}
	if dota.Items.Team3.Player8.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot4")
	}
	if dota.Items.Team3.Player8.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot5")
	}
	if dota.Items.Team3.Player8.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player8 Slot6")
	}
}

//CheckAegisPlayer9 launch search in slots for aegis in player9.
func (dota *Dota2GSIData) CheckAegisPlayer9() {

	if dota.Items.Team3.Player9.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot0")
	}
	if dota.Items.Team3.Player9.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot1")
	}
	if dota.Items.Team3.Player9.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot2")
	}
	if dota.Items.Team3.Player9.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot3")
	}
	if dota.Items.Team3.Player9.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot4")
	}
	if dota.Items.Team3.Player9.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot5")
	}
	if dota.Items.Team3.Player9.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player9 Slot6")
	}
}

func (dota *Dota2GSIData) GetSmokePlayer0() bool {
	if dota.Hero.Team2.Player0.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer1() bool {
	if dota.Hero.Team2.Player1.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer2() bool {
	if dota.Hero.Team2.Player2.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer3() bool {
	if dota.Hero.Team2.Player3.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer4() bool {
	if dota.Hero.Team2.Player4.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer5() bool {
	if dota.Hero.Team3.Player5.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer6() bool {
	if dota.Hero.Team3.Player6.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer7() bool {
	if dota.Hero.Team3.Player7.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer8() bool {
	if dota.Hero.Team3.Player8.Smoked {
		return true
	}
	return false
}

func (dota *Dota2GSIData) GetSmokePlayer9() bool {
	if dota.Hero.Team3.Player9.Smoked {
		return true
	}
	return false
}
