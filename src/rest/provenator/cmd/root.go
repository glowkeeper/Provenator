package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	pkgLog "github.com/glowkeeper/Provenator/src/rest/provenator/pkg/log"
	"github.com/glowkeeper/Provenator/src/rest/provenator/internal/server"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "provenator",
	Short: "provenator RESTful server",
	Long: `provenator RESTful server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {

	cobra.OnInitialize(server.Init)
	rootCmd.PersistentFlags().StringVar(&server.CfgFile, "config", ".provenator.yaml", "file is ./.provenator.yaml")
	rootCmd.PersistentFlags().StringVar(&server.OpenAPIFile, "api", "openapi.yaml", "file is ./openapi.yaml")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	pkgLog.Init(viper.GetString("env"))
}
