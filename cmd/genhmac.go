package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/token"
)

var (
	sig string
	ts  int64
)

// genhmacCmd represents the genhmac command
var genhmacCmd = &cobra.Command{
	Use:   "genhmac",
	Short: "Generate HMAC for a given input",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Validação simples
		if sig == "" {
			fmt.Println("❌ Erro: --sig é obrigatório.")
			os.Exit(1)
		}

		// Se não informar --ts, usa o time.Now()
		var timestamp int64
		if ts == 0 {
			timestamp = time.Now().Unix()
		} else {
			timestamp = ts
		}

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

		// Gera o token
		hmacToken := ti.GenerateTimedHMAC(sig, timestamp)

		// Exibe
		fmt.Println("📦 HMAC Token gerado:")
		fmt.Println("X-HMAC-Token:     ", hmacToken)
		fmt.Println("X-HMAC-Signature: ", sig)
		fmt.Println("X-HMAC-Timestamp: ", timestamp)
	},
}

func init() {
	rootCmd.AddCommand(genhmacCmd)

	genhmacCmd.PersistentFlags().StringVarP(&sig, "signature", "s", "", "Signature for HMAC")
	genhmacCmd.PersistentFlags().Int64VarP(&ts, "timestamp", "t", 0, "Timestamp for HMAC (optional)")

}
