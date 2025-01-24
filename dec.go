package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run dec.go <file-to-decrypt>")
        return
    }

    filePath := os.Args[1]
    key := []byte("examplekey123456") // Key must be 16, 24, or 32 bytes long

    // Read the encrypted content
    encrypted, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }

    // Decrypt the content
    decrypted := decryptFile(encrypted, key)

    // Write the decrypted content back to the original file
    if err := ioutil.WriteFile(filePath, decrypted, 0644); err != nil {
        log.Fatal(err)
    }

    fmt.Println("File decrypted and saved as", filePath)
}

func decryptFile(ciphertext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err)
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        log.Fatal(err)
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        log.Fatal("ciphertext too short")
    }
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        log.Fatal(err)
    }

    return plaintext
}
