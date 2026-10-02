package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func generatePassword(length int, charset string) (string, error) {
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}

func askYesNo(reader *bufio.Reader, prompt string) bool {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password length: ")
	input, _ := reader.ReadString('\n')
	length, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || length <= 0 {
		fmt.Println("❌ Please enter a valid positive number.")
		return
	}

	charset := ""
	if askYesNo(reader, "Include lowercase letters? (y/n): ") {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if askYesNo(reader, "Include uppercase letters? (y/n): ") {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if askYesNo(reader, "Include digits? (y/n): ") {
		charset += "0123456789"
	}
	if askYesNo(reader, "Include symbols? (y/n): ") {
		charset += "!@#$%^&*"
	}

	if charset == "" {
		fmt.Println("❌ You must select at least one character type.")
		return
	}

	password, err := generatePassword(length, charset)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("🔐 Your password:", password)
}