package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var debug bool
var proxy, dockerUser, dockerPassword, nameSpace string
var githubRepo, githubToken string
var queryLimit, processLimit, monitorCount int
var httpTimeOut, syncTimeOut time.Duration

var rootCmd = &cobra.Command{
	Use:   "gcrsync",
	Short: "A docker image sync tool for Google container registry (gcr.io)",
	Long: `
A docker image sync tool for Google container registry (gcr.io).`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
	rootCmd.PersistentFlags().StringVar(&proxy, "proxy", "", "http client proxy")
	rootCmd.PersistentFlags().StringVar(&dockerUser, "user", "", "docker registry user")
	rootCmd.PersistentFlags().StringVar(&dockerPassword, "password", "", "docker registry user password")
	rootCmd.PersistentFlags().StringVar(&nameSpace, "namespace", "google-containers", "google container registry namespace")
	rootCmd.PersistentFlags().IntVar(&queryLimit, "querylimit", 50, "http query limit")
	rootCmd.PersistentFlags().DurationVar(&httpTimeOut, "httptimeout", 10*time.Second, "http request timeout")
	rootCmd.PersistentFlags().DurationVar(&syncTimeOut, "synctimeout", 0, "sync timeout")
	rootCmd.PersistentFlags().IntVar(&processLimit, "processlimit", 10, "image process limit")
	rootCmd.PersistentFlags().StringVar(&githubRepo, "githubrepo", "wyliog/gcr", "github commit repo")
	rootCmd.PersistentFlags().StringVar(&githubToken, "githubtoken", "", "github commit token")
}
