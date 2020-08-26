package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/app"
)

func artefactsList(w http.ResponseWriter) {

	content := app.ArtefactsList()
	w.Write(content)
}

func artefactsAll(w http.ResponseWriter) {

	content := app.ArtefactsAll()
	w.Write(content)
}

func artefactsTotal(w http.ResponseWriter) {

    content := app.ArtefactsTotal()
    w.Write(content)
}

func artefacts(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pArtefact := viper.GetString("paths./artefacts/{artefactsRef}.get.parameters.name")
    artefactsRef := params[pArtefact]
	hArtefactsRef := common.HexToHash(artefactsRef)
	content := app.Artefact(hArtefactsRef)
    w.Write(content)
}