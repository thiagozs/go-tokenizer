package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/crypto"
	"github.com/thiagozs/go-tokenizer/handler"
	"github.com/thiagozs/go-tokenizer/token"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		config, err := config.NewConfig()
		if err != nil {
			fmt.Printf("Erro ao criar configuração: %v", err)
			os.Exit(1)
		}

		ti, err := token.NewToken(token.WithConfig(config))
		if err != nil {
			fmt.Printf("Erro ao criar token: %v", err)
			os.Exit(1)
		}

		crypto, err := crypto.NewSymmetric(crypto.WithConfig(config))
		if err != nil {
			fmt.Printf("Erro ao criar instância de criptografia simétrica: %v", err)
			os.Exit(1)
		}

		hadlr, err := handler.NewHandlers(
			handler.WithConfig(config),
			handler.WithCrypto(crypto),
			handler.WithToken(ti),
		)
		if err != nil {
			fmt.Printf("Erro ao criar manipulador: %v", err)
			os.Exit(1)
		}

		http.HandleFunc("/tokenize", hadlr.RequireHMAC(hadlr.Tokenize))
		http.HandleFunc("/detokenize", hadlr.RequireHMAC(hadlr.Detokenize))

		http.HandleFunc("/selftest-protected", hadlr.ProtectedEndpoint)

		http.HandleFunc("/genhmac", hadlr.GenerateHMACOnline)

		hostPort := fmt.Sprintf("%s:%s", config.GetHost(), config.GetPort())

		log.Println("🔐 Tokenizer API")
		if config.GetHost() != "" {
			log.Println("Running in:", hostPort)
			log.Fatal(http.ListenAndServe(hostPort, nil))
		} else {
			hostPort = ":8880"
			log.Println("Running in: ", hostPort)
			log.Fatal(http.ListenAndServe(hostPort, nil))
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
