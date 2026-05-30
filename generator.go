package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func main() {
	length := flag.Int("length", 12, "Lunghezza della password")
	flag.Parse()
	password, err := generatePassword(*length)
	if err != nil {
		fmt.Println("Errore nella generazione della password:", err)
		return
	}
	fmt.Println(password)
	fmt.Println("Premi Invio per uscire...")
	fmt.Scanln()
}

func generatePassword(length int) (string, error) {
	b := make([]byte, length)
	charsetLength := big.NewInt(int64(len(characters)))

	for i := range b {
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		b[i] = characters[n.Int64()]
	}

	return string(b), nil
}
