package routes

import (
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/common"

	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/internal/app"
)

func jobs (w http.ResponseWriter) {

	content := app.JobsList()
	w.Write(content)
}

func jobsAddress(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pAddress := viper.GetString("paths./jobs-address/{address}.get.parameters.name")
    address := params[pAddress]
	hAddress := common.HexToAddress(address)
	content := app.JobsForAddress(hAddress)
    w.Write(content)
}

func jobsAddressStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pAddress := viper.GetString("paths./jobs-address-status/{address}/{status}.get.parameters.name")
	pStatus := viper.GetString("paths./jobs-address-status/{address}/{status}.get.otherParameters.name")
    address := params[pAddress]
	parsedStatus, _ := strconv.ParseUint(params[pStatus], 10, 8)
	status := uint8(parsedStatus)
	hAddress := common.HexToAddress(address)
	content := app.JobsForAddressStatus(hAddress, status)
    w.Write(content)
}

func jobsStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pStatus := viper.GetString("paths./jobs-status/{status}.get.parameters.name")
	parsedStatus, _ := strconv.ParseUint(params[pStatus], 10, 8)
	status := uint8(parsedStatus)
	content := app.JobsForStatus(status)
    w.Write(content)
}

func jobsAll(w http.ResponseWriter) {

	content := app.JobsAll()
	w.Write(content)
}

func jobsTotal(w http.ResponseWriter) {

    content := app.JobsTotal()
    w.Write(content)
}

func job(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	pJob := viper.GetString("paths./job/{jobRef}.get.parameters.name")
    jobRef := params[pJob]
	hJobRef := common.HexToHash(jobRef)
	content := app.Job(hJobRef)
    w.Write(content)
}
