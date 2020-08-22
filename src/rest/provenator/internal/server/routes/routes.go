package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
)

// RegisterRoutes - register all out routes for the server
func RegisterRoutes(r *mux.Router) {

	uHome := viper.GetString("paths./.get.summary")
	uHelp := viper.GetString("paths./help.get.summary")

	uSponsList := viper.GetString("paths./spons.get.summary")
	uSponsAll := viper.GetString("paths./spons-all.get.summary")
	uSponsTotal := viper.GetString("paths./spons-total.get.summary")
	uSpons := viper.GetString("paths./spons/{sponsRef}.get.summary")

	uJobs := viper.GetString("paths./jobs.get.summary")
	uJobsAll := viper.GetString("paths./jobs-all.get.summary")
	uJobsTotal := viper.GetString("paths./jobs-total.get.summary")
	uJobsAddress := viper.GetString("paths./jobs-address/{address}.get.summary")
	uJobsAddressStatus := viper.GetString("paths./jobs-address-status/{address}/{status}.get.summary")
	uJobsStatus := viper.GetString("paths./jobs-status/{status}.get.summary")
	uJob := viper.GetString("paths./job/{jobRef}.get.summary")

	uSubbies := viper.GetString("paths./subbies.get.summary")
	uSubbiesAll := viper.GetString("paths./subbies-all.get.summary")
	uSubbiesTotal := viper.GetString("paths./subbies-total.get.summary")
	uSubbie := viper.GetString("paths./subbie/{subbieAddress}.get.summary")

	r.PathPrefix("/").HandlerFunc(postHandler).Methods(http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)

	r.HandleFunc(uHome, func(w http.ResponseWriter, r *http.Request) {
    	home(w)
	}).Methods(http.MethodGet)

    r.HandleFunc(uHelp, func(w http.ResponseWriter, r *http.Request) {
    	help(w)
	}).Methods(http.MethodGet)

	// Spons
	r.HandleFunc(uSponsList, func(w http.ResponseWriter, r *http.Request) {
    	sponsList(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSponsAll, func(w http.ResponseWriter, r *http.Request) {
    	sponsAll(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSponsTotal, func(w http.ResponseWriter, r *http.Request) {
    	sponsTotal(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSpons, func(w http.ResponseWriter, r *http.Request) {
    	spons(w, r)
	}).Methods(http.MethodGet)

	// Jobs
	r.HandleFunc(uJobs, func(w http.ResponseWriter, r *http.Request) {
	    jobs(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJobsAll, func(w http.ResponseWriter, r *http.Request) {
	    jobsAll(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJobsTotal, func(w http.ResponseWriter, r *http.Request) {
	    jobsTotal(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJobsAddress, func(w http.ResponseWriter, r *http.Request) {
	    jobsAddress(w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJobsAddressStatus, func(w http.ResponseWriter, r *http.Request) {
	    jobsAddressStatus(w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJobsStatus, func(w http.ResponseWriter, r *http.Request) {
	    jobsStatus(w, r)
	}).Methods(http.MethodGet)

	r.HandleFunc(uJob, func(w http.ResponseWriter, r *http.Request) {
	    job(w, r)
	}).Methods(http.MethodGet)

	// Subbies
	r.HandleFunc(uSubbies, func(w http.ResponseWriter, r *http.Request) {
	    subbiesList(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSubbiesAll, func(w http.ResponseWriter, r *http.Request) {
	    subbiesAll(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSubbiesTotal, func(w http.ResponseWriter, r *http.Request) {
	    subbiesTotal(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uSubbie, func(w http.ResponseWriter, r *http.Request) {
	    subbie(w, r)
	}).Methods(http.MethodGet)
}
