package routes

import (
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/app"
)

func entitiesList (w http.ResponseWriter) {

	content := app.JobsList()
	w.Write(content)
}

func entitiesAll(w http.ResponseWriter) {

	content := app.JobsAll()
	w.Write(content)
}

func entitiesTotal(w http.ResponseWriter) {

    content := app.JobsTotal()
    w.Write(content)
}

func entity(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pJob := viper.GetString("paths./job/{jobRef}.get.parameters.name")
    jobRef := params[pJob]
	hJobRef := common.HexToHash(jobRef)
	content := app.Job(hJobRef)
    w.Write(content)
}
