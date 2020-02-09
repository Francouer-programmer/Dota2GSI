package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"time"

	dota2gsi "github.com/go-programmer/Dota2GSI"
	"github.com/go-programmer/Dota2GSI/app/smoke"
)

var data dota2gsi.Dota2GSIData

//HandlersView takes all hadnlers of the app.
type HandlersView interface {
	RunHandlers()
	DecodeBody(rw http.ResponseWriter, req *http.Request)
}

//Handler implement handlers data.
type Handler struct {
	Checker smoke.PlayersView
}

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

//Run handler
func Run() error {
	dotaData := &dota2gsi.GSIData{
		Data: data,
	}

	playerCheck := &smoke.PlayerChecker{
		GSI: *dotaData,
	}

	handle := &Handler{
		Checker: playerCheck,
	}
	if err := handle.RunHandlers(); err != nil {
		return err
	}
	return nil
}

//RunHandlers launch all handlers.
func (h *Handler) RunHandlers() error {

	fmt.Println("Started")

	http.HandleFunc("/", h.DecodeBody)

	http.HandleFunc("/go", h.parseTemplateGo)

	http.HandleFunc(fmt.Sprintf("/go-%s", "app"), adminCabinetTemplateParse)

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	//http.HandleFunc("/roshan", getRoshanTime)
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		return err
	}
	return nil
}


//ReturnGSI return dota data.
func (h *Handler) ReturnGSI() *dota2gsi.Dota2GSIData {
	return h.Checker.GetGSIData()
}

//DecodeBody handler wrapper that decodes body of the dota2 data send request.
func (h *Handler) DecodeBody(rw http.ResponseWriter, req *http.Request) {
	var data dota2gsi.Dota2GSIData
	req.Close = true
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err)
	}
	h.Checker.GetPlayersStatus()
	h.Checker.GetRoshanTime()
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

func (h *Handler) parseTemplateGo(w http.ResponseWriter, r *http.Request) {
	var idMap = map[string]string{
		"player0": h.Checker.Player0(),
		"player1": h.Checker.Player1(),
		"player2": h.Checker.Player2(),
		"player3": h.Checker.Player3(),
		"player4": h.Checker.Player4(),
		"player5": h.Checker.Player5(),
		"player6": h.Checker.Player6(),
		"player7": h.Checker.Player7(),
		"player8": h.Checker.Player8(),
		"player9": h.Checker.Player9(),
	}

	data := ViewPlayersData{
		DataMap:     idMap,
		RoshanTimer: h.Checker.GetRoshanTime(),
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

func fileNameWin() string {

	r, err := regexp.Compile("([\\w]+)\\.exe")
	if err != nil {
		log.Println(err)
	}
	ProgramName := r.FindStringSubmatch(os.Args[0])[1]

	return ProgramName
}
