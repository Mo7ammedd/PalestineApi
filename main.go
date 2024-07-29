package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "sync"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

var (
    faviconPath   string
    casualtiesUrl string
    killedUrl     string
    westbankUrl   string

    casualtiesJson []byte
    killedJson     []byte
    westbankJson   []byte
    fetchOnce      sync.Once
)

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
    faviconPath = os.Getenv("FAVICON_PATH")
    casualtiesUrl = os.Getenv("CASUALTIES_URL")
    killedUrl = os.Getenv("KILLED_URL")
    westbankUrl = os.Getenv("WESTBANK_URL")
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/killed", killedHandler).Methods("GET")
    r.HandleFunc("/casualties", casualtiesHandler).Methods("GET")
    r.HandleFunc("/westbank", westbankHandler).Methods("GET")
    r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")

    http.Handle("/", handlers.CORS()(r))
    fmt.Println("Server running on port 3000")
    http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "from the river to the sea"})
}

func killedHandler(w http.ResponseWriter, r *http.Request) {
    fetchOnce.Do(fetchData)
    w.Header().Set("Content-Type", "application/json")
    w.Write(killedJson)
}

func casualtiesHandler(w http.ResponseWriter, r *http.Request) {
    fetchOnce.Do(fetchData)
    w.Header().Set("Content-Type", "application/json")
    w.Write(casualtiesJson)
}

func westbankHandler(w http.ResponseWriter, r *http.Request) {
    fetchOnce.Do(fetchData)
    w.Header().Set("Content-Type", "application/json")
    w.Write(westbankJson)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filepath.Join(".", faviconPath))
}

func fetchData() {
    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        defer wg.Done()
        resp, err := http.Get(casualtiesUrl)
        if err != nil {
            fmt.Println("Error fetching casualties data:", err)
            return
        }
        defer resp.Body.Close()
        casualtiesJson, _ = ioutil.ReadAll(resp.Body)
    }()

    go func() {
        defer wg.Done()
        resp, err := http.Get(killedUrl)
        if err != nil {
            fmt.Println("Error fetching killed data:", err)
            return
        }
        defer resp.Body.Close()
        killedJson, _ = ioutil.ReadAll(resp.Body)
    }()

    go func() {
        defer wg.Done()
        resp, err := http.Get(westbankUrl)
        if err != nil {
            fmt.Println("Error fetching westbank data:", err)
            return
        }
        defer resp.Body.Close()
        westbankJson, _ = ioutil.ReadAll(resp.Body)
    }()

    wg.Wait()
}
