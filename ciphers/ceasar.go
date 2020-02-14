package ciphers

import (
	"fmt"
	"strings"
)

/**
A function that encrypts a plaintext with Ceasar Cipher
*/
func CeasarCipherEncrypt(plaintext string, alphabet string, shifts int) {

}

/**
A function that decrypts a ciphertext with Ceasar Cipher
*/
func CeasarCipherDecrypt(ciphertext string, alphabet string, shifts int) {

}

/**
Brute Force a ciphertext to find the plaintext
*/
func CeasarBF(ciphertext string, alphabet string) {

	var decrypted string

	for key := 0; key < len(alphabet); key++ {
		decrypted = ""
		for j := 0; j < len(ciphertext); j++ {
			if strings.ContainsAny(alphabet, string(ciphertext[j])) {
				pos := strings.IndexAny(alphabet, string(ciphertext[j]))
				pos = pos - key

				if pos < 0 {
					pos = pos + len(alphabet)
				}

				decrypted = decrypted + string(alphabet[pos])
			} else {
				decrypted = decrypted + string(ciphertext[j])
			}
		}
		fmt.Printf("Key#%d: %s\n", key, decrypted)

	}

}
