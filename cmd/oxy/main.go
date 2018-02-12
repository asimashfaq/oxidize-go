package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	walletCmd "github.com/tclchiam/oxidize-go/wallet/cmd"
)

var (
	rootCmd = &cobra.Command{
		Use:   "oxy",
		Short: "Oxidize client CLI",
		Long:  "CLI for Oxidize protocol",
	}
)

func init() {
	rootCmd.AddCommand(walletCmd.WalletCmd)

	viper.SetDefault("data.dir", "~/.oxy/data")
	viper.SetDefault("node.addr", "localhost:8080")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}
}
