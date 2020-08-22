package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/app"
)

func sponsList(w http.ResponseWriter) {

	content := app.SponsList()
	w.Write(content)
}

func sponsAll(w http.ResponseWriter) {

	content := app.SponsAll()
	w.Write(content)
}

func sponsTotal(w http.ResponseWriter) {

    content := app.SponsTotal()
    w.Write(content)
}

func spons(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pSpons := viper.GetString("paths./spons/{sponsRef}.get.parameters.name")
    sponsRef := params[pSpons]
	hSponsRef := common.HexToHash(sponsRef)
	content := app.Spons(hSponsRef)
    w.Write(content)
}
