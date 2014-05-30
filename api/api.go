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

type Application struct {
	DbMap   *gorp.DbMap
	Router  *mux.Router
	Handler http.Handler
}

func NewApp(dbname, dbuser, dbpass,
	logPrefix, secretKey, apiPrefix, staticPrefix, staticDir string) (*Application, error) {

	app := &Application{}

	if err := app.InitDb(dbname, dbuser, dbpass, logPrefix); err != nil {
		return nil, err
	}
	app.InitRouter(apiPrefix, staticPrefix, staticDir, secretKey)
	return app, nil
}

func (app *Application) InitDb(dbname, dbuser, dbpass, logPrefix string) error {

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s", dbuser, dbname, dbpass))
	if err != nil {
		return err
	}

	app.DbMap, err = models.Configure(db, logPrefix)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) InitRouter(apiPrefix, staticPrefix, staticDir, secretKey string) {

	app.Router = mux.NewRouter()

	// API

	routes.Configure(app.Router.PathPrefix(apiPrefix).Subrouter())

	// STATIC FILES

	app.Router.PathPrefix(staticPrefix).Handler(http.FileServer(http.Dir(staticDir)))
	app.Handler = csrf.NewCSRF(secretKey, app.Router)
}

func (app *Application) Shutdown() {
	app.DbMap.Db.Close()
}
