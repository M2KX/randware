# File Encryption and Decryption Tool

This repository contains two Go scripts, `enc.go` and `dec.go`, for encrypting and decrypting files using the **AES-GCM** (Advanced Encryption Standard - Galois/Counter Mode) algorithm.

---

## Prerequisites

- Go programming language installed on your system.
- A file to encrypt or decrypt.

---

## Usage

### Encryption (`enc.go`)

To encrypt a file, use the `enc.go` script. The script will encrypt the specified file and replace it !.

```
go run enc.go <input-file>
```

### Decryption (`dec.go`)

To decrypt a file, use the dec.go script. The script will decrypt the specified file and replace it with the original.

```
go run dec.go <input-file>
```

---

## Encryption Key

The encryption key is read from the `ENCRYPTION_KEY` environment variable. Ensure that the key is 16, 24, or 32 bytes long.

Example:

```
ENCRYPTION_KEY="examplekey123456"
```

---

## Important Notes

- Key Security: Do not hardcode the encryption key in the source code. Use environment variables or a secure key management system.

- File Overwriting: The scripts **do** overwrite the original files.make sure to backup important files.
