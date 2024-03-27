package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Games Db Api...\nRequested path: %s", r.URL.Path[1:])
}

func main() {

	// İlk Örnek
	//http.HandleFunc("/", handler)
	//err := http.ListenAndServe(":9000", nil)

	// İkinci örnek
	// var iRobot Robot
	// err := http.ListenAndServe(":9000", iRobot)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Üçüncü örnek
	// http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
	// 	msg := RobotApi{"Robot Api Root"}
	// 	output, err := json.Marshal(msg)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Fprint(w, string(output))
	// })

	// http.HandleFunc(apiRoot+"/robots", func(w http.ResponseWriter, r *http.Request) {
	// 	robots := []Robot{
	// 		{Nickname: "Gemini", Region: "Earth Zone", Version: 1},
	// 		{Nickname: "WARP", Region: "Out of Space Zone", Version: 4},
	// 		{Nickname: "HAL - 2001", Region: "Black Hole Zone", Version: 99},
	// 	}
	// 	output, err := json.Marshal(robots)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Fprint(w, string(output))
	// })

	// http.HandleFunc(apiRoot+"/robots/me", func(w http.ResponseWriter, r *http.Request) {
	// 	robot := Robot{"T-1000", "Unknown Zone", 13}
	// 	output, err := json.Marshal(robot)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Fprint(w, string(output))
	// })

	// 4ncü örnek (Yeni Go versiyonunda Gorilla'nın routing mekanizması dahili paketlere eklenmiş sanırım.)
	// api/users/13 şeklinde test edilebilir
	gorillaRouter := mux.NewRouter()
	gorillaRouter.HandleFunc("/api/users/{id:[0-9]+}", TinyMuxHandler)
	http.Handle("/", gorillaRouter)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Web Server is listening...")
}

type Robot struct {
	Nickname string
	Region   string
	Version  int
}

func (r Robot) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Nickname = "Robo IO 1001"
	r.Region = "Red Zone"
	r.Version = 1

	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("Path : ", req.URL.Path)
	fmt.Fprintf(w, "<table style='border:1'><tr><td><b>Nickname</b></td><td><b>Region</b></td><td><b>Version</b></td></tr><tr><td>%s</td><td>%s</td><td>%d</td></tr></table>", r.Nickname, r.Region, r.Version)
}

type RobotApi struct {
	Message string `json:"message"`
}

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Point int    `json:"int"`
}

const apiRoot string = "/api"

func TinyMuxHandler(w http.ResponseWriter, req *http.Request) {
	urlParams := mux.Vars(req)
	id := urlParams["id"]
	content := "User Id : " + id
	message := RobotApi{content}
	output, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(output))
}
