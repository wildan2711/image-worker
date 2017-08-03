package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gophercode/base/security"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd *cobra.Command

type Option struct {
	Verbose          bool   `mapstructure:"verbose"`
	JobServerAddress string `mapstructure:"job-server"`
	AuthKey          string `mapstructure:"auth-key"`
	SSLPath          string `mapstructure:"ssl-path"`
	SSLEnabled       bool   `mapstructure:"ssl-enabled"`
	HookEnabled      bool   `mapstructure:"hook-enabled"`
	ManagerAddress   string `mapstructure:"manager-address"`
}

func init() {

	progName := filepath.Base(os.Args[0])

	viper.SetConfigType("yaml")
	viper.SetConfigName(progName)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", progName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", progName))
	viper.AddConfigPath(".")

	opt := &Option{}

	cobra.OnInitialize(func() {

		if err := viper.ReadInConfig(); err != nil {
			log.Println(err)
		}

		if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
			log.Println(err)
		}

		if err := viper.Unmarshal(opt); err != nil {
			log.Println(err)
		}

		if !opt.Verbose {
			log.SetOutput(ioutil.Discard)
		}

	})

	RootCmd = &cobra.Command{
		Use:   progName,
		Short: fmt.Sprintf("%s is a image manipulation service", progName),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	RootCmd.PersistentFlags().Bool("verbose", true, "Verbose logging information")
	RootCmd.PersistentFlags().Bool("ssl", false, "Enable/disable SSL support")
	RootCmd.PersistentFlags().Bool("hook-enabled", true, "Enable/disable http hook")
	RootCmd.PersistentFlags().String("job-server", "127.0.0.1:4730", "Gearman job server")
	RootCmd.PersistentFlags().String("manager-address", ":8000", "Manager listening address")
	RootCmd.PersistentFlags().String("auth-key", "", "Authentication key")
	RootCmd.PersistentFlags().String("ssl-path", "", "SSL certificate file path")
	RootCmd.PersistentFlags().String("reports-dir", "./_reports", "Report files directory")
	RootCmd.PersistentFlags().String("caches-dir", "./_caches", "Generated report files directory")

	RootCmd.AddCommand(&cobra.Command{
		Use: "jobs",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})

	RootCmd.AddCommand(&cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})

	RootCmd.AddCommand(&cobra.Command{
		Use: "gen-auth-key",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})

	RootCmd.AddCommand(&cobra.Command{
		Use: "gen-secret-key",
		Run: func(cmd *cobra.Command, args []string) {
			key, err := security.GenerateKey()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(os.Stderr, "%s\n", base64.StdEncoding.EncodeToString(key[:]))
		},
	})

	RootCmd.AddCommand(&cobra.Command{
		Use:   "diagnostics",
		Short: "Perform self diagnostic testing",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})

}
