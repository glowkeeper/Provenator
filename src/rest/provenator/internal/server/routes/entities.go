package routes

import (
	"net/http"
	//"strconv"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/app"
)

func entitiesList (w http.ResponseWriter) {

	content := app.EntitiesList()
	w.Write(content)
}

func entitiesAll(w http.ResponseWriter) {

	content := app.EntitiesAll()
	w.Write(content)
}

func entitiesTotal(w http.ResponseWriter) {

    content := app.EntitiesTotal()
    w.Write(content)
}

func entitiesTypes(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pEntity := viper.GetString("paths./entities-types/{entitiesRef}.get.parameters.name")
    entityRef := params[pEntity]
	hEntityRef := common.HexToHash(entityRef)
	content := app.EntityTypes(hEntityRef)
    w.Write(content)
}

func entity(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pEntity := viper.GetString("paths./entities/{entitiesRef}.get.parameters.name")
    entityRef := params[pEntity]
	hEntityRef := common.HexToHash(entityRef)
	content := app.Entity(hEntityRef)
    w.Write(content)
}
