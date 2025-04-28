package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/crypto"
)

func main() {
	input := "12345678900"

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := crypto.NewSymmetric(crypto.WithConfig(cfg))
	if err != nil {
		log.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric(input)
	if err != nil {
		log.Fatal("Erro ao criptografar:", err)
	}
	fmt.Println("Token:", token)

	plain, err := es.DecryptSymmetric(token)
	if err != nil {
		log.Fatal("Erro ao descriptografar:", err)
	}

	fmt.Println("Descriptografado:", string(plain))
}
