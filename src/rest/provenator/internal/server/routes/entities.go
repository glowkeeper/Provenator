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

func entity(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pJob := viper.GetString("paths./job/{jobRef}.get.parameters.name")
    jobRef := params[pJob]
	hJobRef := common.HexToHash(jobRef)
	content := app.Entity(hJobRef)
    w.Write(content)
}
