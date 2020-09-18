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

func artefactsEntity(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pEntity := viper.GetString("paths./artefacts-entity/{entitiesRef}.get.parameters.name")
    entityRef := params[pEntity]
	hEntityRef := common.HexToHash(entityRef)
	content := app.ArtefactEntity(hEntityRef)
    w.Write(content)
}

func artefactsAuthors(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pArtefact := viper.GetString("paths./artefacts-authors/{artefactsRef}.get.parameters.name")
    artefactsRef := params[pArtefact]
	hArtefactsRef := common.HexToHash(artefactsRef)
	content := app.ArtefactAuthors(hArtefactsRef)
    w.Write(content)
}

func artefactsCopyrightHolders(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pArtefact := viper.GetString("paths./artefacts-copyright-holders/{artefactsRef}.get.parameters.name")
    artefactsRef := params[pArtefact]
	hArtefactsRef := common.HexToHash(artefactsRef)
	content := app.ArtefactCopyrightHolders(hArtefactsRef)
    w.Write(content)
}

func artefactsPublishers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pArtefact := viper.GetString("paths./artefacts-publishers/{artefactsRef}.get.parameters.name")
    artefactsRef := params[pArtefact]
	hArtefactsRef := common.HexToHash(artefactsRef)
	content := app.ArtefactPublishers(hArtefactsRef)
    w.Write(content)
}
