package api

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/danjac/angular-react-compare/api/csrf"
	"github.com/danjac/angular-react-compare/api/models"
	"github.com/danjac/angular-react-compare/api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

type Config struct {
	DbName, DbUser, DbPassword, LogPrefix, ApiPrefix, StaticPrefix, StaticDir, SecretKey string
}

type Application struct {
	Config  *Config
	DbMap   *gorp.DbMap
	Router  *mux.Router
	Handler http.Handler
}

func NewApp(config *Config) (*Application, error) {

	app := &Application{Config: config}

	if err := app.InitDb(); err != nil {
		return nil, err
	}
	app.InitRouter()
	return app, nil
}

func (app *Application) InitDb() error {

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s",
		app.Config.DbUser,
		app.Config.DbName,
		app.Config.DbPassword))
	if err != nil {
		return err
	}

	app.DbMap, err = models.Configure(db, app.Config.LogPrefix)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) InitRouter() {

	app.Router = mux.NewRouter()

	// API

	routes.Configure(app.Router.PathPrefix(app.Config.ApiPrefix).Subrouter())

	// STATIC FILES

	app.Router.PathPrefix(app.Config.StaticPrefix).Handler(http.FileServer(http.Dir(app.Config.StaticDir)))
	app.Handler = csrf.NewCSRF(app.Config.SecretKey, app.Router)
}

func (app *Application) Shutdown() {
	app.DbMap.Db.Close()
}
