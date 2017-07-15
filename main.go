package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

// Adapter adapter
type Adapter struct {
	h func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

// ServeHTTP serve http
func (a Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// App application
type App struct {
	BlockChain *Chain
	Logger     *log.Logger
}

// GetBlocks get blocks
func (app *App) GetBlocks(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, app.BlockChain, nil
}

// CreateBlockRequest create block request
type CreateBlockRequest struct {
	Data string `json:"data"`
}

// CreateBlock get blocks
func (app *App) CreateBlock(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var req CreateBlockRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	app.Logger.Printf("%v", req)
	b, err := app.BlockChain.GenerateNextBlock(req.Data)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	app.Logger.Printf("%v", b)
	return http.StatusOK, app.BlockChain, nil
}

func main() {
	port := flag.String("port", "8080", "server port")
	flag.Parse()

	app := App{
		BlockChain: NewBlockChain(),
		Logger:     log.New(os.Stdout, "[app]: ", log.Lmicroseconds),
	}
	mux := http.NewServeMux()
	mux.Handle("/blocks", Adapter{h: app.GetBlocks})
	mux.Handle("/blocks/add", Adapter{h: app.CreateBlock})

	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		log.Fatal(err)
	}
}
