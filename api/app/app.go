package app

import (
	"log"
	"net/http"

	"github.com/VReference/api/app/handler"
	"github.com/VReference/api/app/model"
	"github.com/VReference/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	db, err := gorm.Open("mysql", "ujhn57abqqlauq4w:4D3xwIKjyfILPn0tEB5V@(bubxvxkwb1z9i36sqsy0-mysql.services.clever-cloud.com:3306)/bubxvxkwb1z9i36sqsy0?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/headsets", a.GetAllHeadsets)
	a.Post("/headsets", a.CreateHeadset)
	a.Get("/headsets/{title}", a.GetHeadset)
	a.Put("/headsets/{title}", a.UpdateHeadset)
	a.Delete("/headsets/{title}", a.DeleteHeadset)
	a.Put("/headsets/{title}/disable", a.DisableHeadset)
	a.Put("/headsets/{title}/enable", a.EnableHeadset)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Headset Data
func (a *App) GetAllHeadsets(w http.ResponseWriter, r *http.Request) {
	handler.GetAllHeadsets(a.DB, w, r)
}

func (a *App) CreateHeadset(w http.ResponseWriter, r *http.Request) {
	handler.CreateHeadset(a.DB, w, r)
}

func (a *App) GetHeadset(w http.ResponseWriter, r *http.Request) {
	handler.GetHeadset(a.DB, w, r)
}

func (a *App) UpdateHeadset(w http.ResponseWriter, r *http.Request) {
	handler.UpdateHeadset(a.DB, w, r)
}

func (a *App) DeleteHeadset(w http.ResponseWriter, r *http.Request) {
	handler.DeleteHeadset(a.DB, w, r)
}

func (a *App) DisableHeadset(w http.ResponseWriter, r *http.Request) {
	handler.DisableHeadset(a.DB, w, r)
}

func (a *App) EnableHeadset(w http.ResponseWriter, r *http.Request) {
	handler.EnableHeadset(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}