package server

import (
    //"fmt"
    //"reflect"
    "time"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    "go.uber.org/zap"
	"github.com/spf13/viper"

    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/text"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/server/routes"
    "github.com/glowkeeper/Provenator/src/rest/provenator/internal/contracts"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/server/middleware"

    pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
)

// CfgFile - server configuration
var CfgFile string
// OpenAPIFile - openapi paths  et al.
var OpenAPIFile string

// Init - sets files
func Init() {

	if CfgFile != "" {

		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
		if err := viper.ReadInConfig(); err != nil {
			pkgLog.SLogger.Fatal(err)
		}

	} else {

		pkgLog.SLogger.Fatal(text.ErrorConfig)
	}

	if OpenAPIFile != "" {

		// Use config file from the flag.
		viper.SetConfigFile(OpenAPIFile)
		if err := viper.MergeInConfig(); err != nil {
			pkgLog.SLogger.Fatal(err)
		}

	} else {

		pkgLog.SLogger.Fatal(text.ErrorConfig)
	}

    contracts.InitContracts()

}


// Start - fire up the server
func Start() {

    uRL := viper.GetString("servers.url")
    port := viper.GetString("servers.variables.port.default")
    sDomain := uRL + port

    wTimeOut := time.Duration(viper.GetInt("timeout.write"))
    rTimeOut := time.Duration(viper.GetInt("timeout.read"))
    iTimeOut := time.Duration(viper.GetInt("timeout.idle"))

    cors := handlers.CORS(
        handlers.AllowedHeaders([]string{"content-type"}),
        handlers.AllowedMethods([]string{"GET"}),
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowCredentials(),
    )

    r := mux.NewRouter().StrictSlash(true)
	r.Use(cors, middleware.Logging(), middleware.ContentType())
    routes.RegisterRoutes(r)

    srv := &http.Server{
		Handler:          r,
        Addr:             sDomain,
        WriteTimeout:     time.Second * wTimeOut,
		ReadTimeout:      time.Second * rTimeOut,
		IdleTimeout:      time.Second * iTimeOut,
	}

    pkgLog.SLogger.Info(sDomain)

	if err := srv.ListenAndServe(); err != nil {
		pkgLog.SLogger.Fatal(text.ErrorServer, zap.Error(err))
	}
}
