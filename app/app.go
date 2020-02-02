package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-programmer/Dota2GSI/app/smoke"
)

//ViewPlayersData in general player data information.
type ViewPlayersData struct {
	DataMap     map[string]string
	RoshanTimer string
}

//AdminModelData expiremental data usage.
type AdminModelData struct {
	AdminName   string
	ProgramName string
}

//SmokeData global variable for SmokeStatus
var SmokeData smoke.SmokeStatus

var roshanStateData string

func fileNameWin() string {

	r, err := regexp.Compile("([\\w]+)\\.exe")
	if err != nil {
		log.Println(err)
	}
	ProgramName := r.FindStringSubmatch(os.Args[0])[1]

	return ProgramName
}

func main() {
	fmt.Println("Started")
	http.HandleFunc("/", decodeBody)

	http.HandleFunc("/go", parseTemplateGo)

	http.HandleFunc(fmt.Sprintf("/go-%s", fileNameWin()), adminCabinetTemplateParse)

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	//http.HandleFunc("/roshan", getRoshanTime)
	ifErrorDo(http.ListenAndServe("localhost:3000", nil))
}

func adminCabinetTemplateParse(w http.ResponseWriter, r *http.Request) {
	adminData := AdminModelData{
		AdminName:   "Admin", //TODO: replace to func getName
		ProgramName: fileNameWin(),
	}
	tmpl, err := template.ParseFiles("./assets/html/admin.html")
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, adminData)
}

func parseTemplateGo(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template.ParseFiles("./assets/html/index.html")
	if err != nil {
		log.Println(err)
	}
	everyThreeSecPrintMap(idMap)
	tmpl.Execute(w, data)
}

func everyThreeSecPrintMap(idMap map[string]string) {
	time.Sleep(2 * time.Second)
	fmt.Println(idMap)
}

func decodeBody(rw http.ResponseWriter, req *http.Request) {
	var pl smoke.PlayerChecker

	req.Close = true

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&pl.GSI)
	if err != nil {
		log.Println(err)
	}
	pl.GSI.GetRoshanTime()
	pl.GetPlayersStatus()
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

func ifErrorDo(err error) {
	if err != nil {
		log.Println(err)
	}
}
