package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run enc.go <file-to-encrypt>")
        return
    }

    filePath := os.Args[1]
    key := []byte("examplekey123456") // Key must be 16, 24, or 32 bytes long

    // Encrypt the file and overwrite it
    encrypted := encryptFile(filePath, key)

    // Write the encrypted content back to the original file
    if err := ioutil.WriteFile(filePath, encrypted, 0644); err != nil {
        log.Fatal(err)
    }

    fmt.Println("File encrypted and saved as", filePath)
}

func encryptFile(filePath string, key []byte) []byte {
    plaintext, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err)
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        log.Fatal(err)
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        log.Fatal(err)
    }

    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext
}
