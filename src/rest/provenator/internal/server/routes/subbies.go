package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/app"
)

func subbiesList (w http.ResponseWriter) {

	content := app.SubbiesList()
	w.Write(content)
}

func subbiesAll(w http.ResponseWriter) {

	content := app.SubbiesAll()
	w.Write(content)
}

func subbiesTotal(w http.ResponseWriter) {

    content := app.SubbiesTotal()
    w.Write(content)
}

func subbie(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pSubbie := viper.GetString("paths./subbie/{subbieAddress}.get.parameters.name")
    subbieAddress := params[pSubbie]
	hSubbieAddress := common.HexToHash(subbieAddress)
	content := app.Subbie(hSubbieAddress)
    w.Write(content)
}
