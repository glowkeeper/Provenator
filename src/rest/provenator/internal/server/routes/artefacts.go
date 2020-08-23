package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/app"
)

func artefactsList(w http.ReartefactseWriter) {

	content := app.SponsList()
	w.Write(content)
}

func artefactsAll(w http.ReartefactseWriter) {

	content := app.SponsAll()
	w.Write(content)
}

func artefactsTotal(w http.ReartefactseWriter) {

    content := app.SponsTotal()
    w.Write(content)
}

func artefacts(w http.ReartefactseWriter, r *http.Request) {

	params := mux.Vars(r)
	pSpons := viper.GetString("paths./artefacts/{artefactsRef}.get.parameters.name")
    artefactsRef := params[pSpons]
	hSponsRef := common.HexToHash(artefactsRef)
	content := app.Spons(hSponsRef)
    w.Write(content)
}
