package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	ww "github.com/tnantoka/webwindow"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"time"
)

func renderTemplate(name string, w http.ResponseWriter) {
	t := template.New(name)
	files := []string{"data/views/layout.html", "data/views/" + name + ".html"}

	for _, f := range files {
		t.Parse(string(MustAsset(f)))
	}

	data := map[string]string{
		"name":   name,
		name:     "1",
		"now":    strconv.Itoa(int(time.Now().Unix())),
		"roomID": NewSettings().RoomID,
	}

	t.ExecuteTemplate(w, "layout", data)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	renderTemplate("index", w)
}

func handleSettings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate("settings", w)
	case "POST":
		r.ParseForm()
		roomID := r.Form["room_id"][0]
		apiToken := r.Form["api_token"][0]

		settings := NewSettings()
		if roomID != "" {
			settings.RoomID = roomID
		}
		if apiToken != "" {
			settings.APIToken = apiToken
		}
		settings.Save()

		http.Redirect(w, r, "/index.html", http.StatusFound)
	}
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	chatwork := Chatwork{
		Settings: NewSettings(),
	}
	chatwork.PostMessage(r.Form["body"][0])
	http.Redirect(w, r, "/index.html", http.StatusFound)
}

func assetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "data/public"}
}

func initHTTP() {
	http.HandleFunc("/index.html", handleRoot)
	http.HandleFunc("/settings", handleSettings)
	http.HandleFunc("/messages", handleMessages)
	http.Handle("/", http.FileServer(assetFS()))
}

func storePath(component string) string {
	user, _ := user.Current()
	return user.HomeDir + "/Library/Application Support/cwww/" + component
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func init() {
	dir := storePath("")
	if !fileExists(dir) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	initHTTP()
	config := ww.NewConfig(33333, "CWWW")
	config.Width = 300
	config.Height = 340
	ww.Open(config)
}
