package routes

import (
	"net/http"

	"github.com/spf13/viper"
	"github.com/gorilla/mux"
)

// RegisterRoutes - register all out routes for the server
func RegisterRoutes(r *mux.Router) {

	uHome := viper.GetString("paths./.get.summary")

	uArtefactsList := viper.GetString("paths./artefacts.get.summary")
	uArtefactsAll := viper.GetString("paths./artefacts-all.get.summary")
	uArtefactsTotal := viper.GetString("paths./artefacts-total.get.summary")
	uArtefacts := viper.GetString("paths./artefacts/{artefactsRef}.get.summary")

	uEntitiesList := viper.GetString("paths./entities.get.summary")
	uEntitiesAll := viper.GetString("paths./entities-all.get.summary")
	uEntitiesTotal := viper.GetString("paths./entities-total.get.summary")
	uEntity := viper.GetString("paths./entities/{entitiesRef}.get.summary")

	r.PathPrefix("/").HandlerFunc(postHandler).Methods(http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)

	r.HandleFunc(uHome, func(w http.ResponseWriter, r *http.Request) {
    	home(w)
	}).Methods(http.MethodGet)

	// Artefacts
	r.HandleFunc(uArtefactsList, func(w http.ResponseWriter, r *http.Request) {
    	artefactsList(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uArtefactsAll, func(w http.ResponseWriter, r *http.Request) {
    	artefactsAll(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uArtefactsTotal, func(w http.ResponseWriter, r *http.Request) {
    	artefactsTotal(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uArtefacts, func(w http.ResponseWriter, r *http.Request) {
    	artefacts(w, r)
	}).Methods(http.MethodGet)

	// Entities
	r.HandleFunc(uEntitiesList, func(w http.ResponseWriter, r *http.Request) {
	    entitiesList(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uEntitiesAll, func(w http.ResponseWriter, r *http.Request) {
	    entitiesAll(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uEntitiesTotal, func(w http.ResponseWriter, r *http.Request) {
	    entitiesTotal(w)
	}).Methods(http.MethodGet)

	r.HandleFunc(uEntity, func(w http.ResponseWriter, r *http.Request) {
	    entity(w, r)
	}).Methods(http.MethodGet)
}
