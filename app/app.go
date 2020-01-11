package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	dota2gsi "github.com/go-programmer/Dota2GSI"
)

type ViewPlayersData struct {
	DataMap     map[string]string
	RoshanTimer string
}

//GSIData - shortcut type for incoming data from dota2.
type GSIData dota2gsi.Dota2GSIData

//SmokerStatus - contains "visibility"/"hidden" style property
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

var SmokeData SmokeStatus

var roshanStateData string

func main() {
	fmt.Println("Started")
	http.HandleFunc("/", decodeBody)

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {

		var idMap = map[string]string{
			"player0": SmokeData.StPlayer0,
			"player1": SmokeData.StPlayer1,
			"player2": SmokeData.StPlayer2,
			"player3": SmokeData.StPlayer3,
			"player4": SmokeData.StPlayer4,
			"player5": SmokeData.StPlayer5,
			"player6": SmokeData.StPlayer6,
			"player7": SmokeData.StPlayer7,
			"player8": SmokeData.StPlayer8,
			"player9": SmokeData.StPlayer9,
		}

		data := ViewPlayersData{
			DataMap:     idMap,
			RoshanTimer: roshanStateData,
		}

		tmpl, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Println(err)
		}
		//
		everyThreeSecPrintMap(idMap)
		//
		tmpl.Execute(w, data)
	})
	//http.HandleFunc("/roshan", getRoshanTime)
	ifErrorDo(http.ListenAndServe("localhost:3000", nil))
}

func everyThreeSecPrintMap(idMap map[string]string) {
	time.Sleep(2 * time.Second)
	fmt.Println(idMap)
}

func decodeBody(rw http.ResponseWriter, req *http.Request) {
	var gsiData GSIData

	req.Close = true

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&gsiData)
	if err != nil {
		log.Println(err)
	}
	getRoshanTime(&gsiData)
	getPlayersStatus(&gsiData)
}

func getPlayersStatus(gsi *GSIData) {
	SmokeData = SmokeStatus{
		StPlayer0: Player0(gsi),
		StPlayer1: Player1(gsi),
		StPlayer2: Player2(gsi),
		StPlayer3: Player3(gsi),
		StPlayer4: Player4(gsi),
		StPlayer5: Player5(gsi),
		StPlayer6: Player6(gsi),
		StPlayer7: Player7(gsi),
		StPlayer8: Player8(gsi),
		StPlayer9: Player9(gsi),
	}
}

func writeToFile(req *http.Request) {
	output, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(string(output))
	if err != nil {
		log.Fatal(err)
	}
	f.Sync()
}

func Player0(gsi *GSIData) string {
	gsi.CheckAegisPlayer0()
	if gsi.getSmokePlayer0() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player1(gsi *GSIData) string {
	gsi.CheckAegisPlayer1()
	if gsi.getSmokePlayer1() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player2(gsi *GSIData) string {
	gsi.CheckAegisPlayer2()
	if gsi.getSmokePlayer2() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player3(gsi *GSIData) string {
	gsi.CheckAegisPlayer3()
	if gsi.getSmokePlayer3() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player4(gsi *GSIData) string {
	gsi.CheckAegisPlayer4()
	if gsi.getSmokePlayer4() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player5(gsi *GSIData) string {
	gsi.CheckAegisPlayer5()
	if gsi.getSmokePlayer5() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player6(gsi *GSIData) string {
	gsi.CheckAegisPlayer6()
	if gsi.getSmokePlayer6() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player7(gsi *GSIData) string {
	gsi.CheckAegisPlayer7()
	if gsi.getSmokePlayer7() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player8(gsi *GSIData) string {
	gsi.CheckAegisPlayer8()
	if gsi.getSmokePlayer8() {
		return "visible"
	} else {
		return "hidden"
	}
}

func Player9(gsi *GSIData) string {
	gsi.CheckAegisPlayer9()
	if gsi.getSmokePlayer9() {
		return "visible"
	} else {
		return "hidden"
	}
}

func getRoshanTime(gsiData *GSIData) {
	if gsiData.Map.RoshanState != "alive" {
		time.Sleep(2 * time.Second)
		fmt.Println("Roshan state")
		roshanStateData = fmt.Sprintf("Roshan state => %d", gsiData.Map.RoshanStateEndSeconds)
		fmt.Println(gsiData.Map.RoshanStateEndSeconds)
	}
}

func (dota GSIData) CheckAegisPlayer0() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer1() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer2() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer3() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer4() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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
func (dota GSIData) CheckAegisPlayer5() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
	if dota.Items.Team3.Player5.Slot0.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot0 %s\n")
	}
	if dota.Items.Team3.Player5.Slot1.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot1 %s\n")
	}
	if dota.Items.Team3.Player5.Slot2.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot2 %s\n")
	}
	if dota.Items.Team3.Player5.Slot3.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot3 %s\n")
	}
	if dota.Items.Team3.Player5.Slot4.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot4 %s\n")
	}
	if dota.Items.Team3.Player5.Slot5.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot5 %s\n")
	}
	if dota.Items.Team3.Player5.Slot6.Name == "item_aegis" {
		fmt.Println("Aegis Player5 Slot6 %s\n")
	}
}
func (dota GSIData) CheckAegisPlayer6() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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
func (dota GSIData) CheckAegisPlayer7() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer8() {
	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota GSIData) CheckAegisPlayer9() {

	//value := reflect.ValueOf(dota)
	//typeOfS := value.Type()
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	//}
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

func (dota *GSIData) getSmokePlayer0() bool {
	if dota.Hero.Team2.Player0.Smoked {
		return true
	}
	return false
}

func (dota *GSIData) getSmokePlayer1() bool {
	if dota.Hero.Team2.Player1.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer2() bool {
	if dota.Hero.Team2.Player2.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer3() bool {
	if dota.Hero.Team2.Player3.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer4() bool {
	if dota.Hero.Team2.Player4.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer5() bool {
	if dota.Hero.Team3.Player5.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer6() bool {
	if dota.Hero.Team3.Player6.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer7() bool {
	if dota.Hero.Team3.Player7.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer8() bool {
	if dota.Hero.Team3.Player8.Smoked {
		return true
	}
	return false
}
func (dota *GSIData) getSmokePlayer9() bool {
	if dota.Hero.Team3.Player9.Smoked {
		return true
	}
	return false
}

func ifErrorDo(err error) {
	if err != nil {
		log.Println(err)
	}
}
