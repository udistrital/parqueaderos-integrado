// Playbook - http://play.golang.org/p/3wFl4lacjX

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	//"crypto/rand"
	"errors"
	"fmt"
	"github.com/gotamer/conv"
	//"io"
)

func Pad(src []byte) []byte {
	fmt.Println(len(src))
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encryptCBC(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	//fmt.Println(block)
	if err != nil {
		return "", err
	}

	msg := []byte(text)
	//fmt.Println(len(msg))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	//iv := ciphertext[:aes.BlockSize]
	//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//		return "", err
	//	}
	iv := []byte("abcdefghijklmnop")

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], msg)
	fmt.Println(ciphertext)
	finalMsg := conv.Bl2s(ciphertext)
	return finalMsg, nil
}

func decryptCBC(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedMsg := []byte{165, 7, 155, 74, 5, 229, 106, 210, 117, 186, 190, 152, 215, 220, 227, 30}
	fmt.Println((decodedMsg))
	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	//iv := decodedMsg[:aes.BlockSize]
	iv := []byte("abcdefghijklmnop")
	msg := decodedMsg
	fmt.Println(iv)
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(msg, msg)

	//unpadMsg, err := Unpad(msg)
	unpadMsg := (msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

func main() {
	key := []byte("omarleonardozamb")
	encryptMsg, _ := encryptCBC(key, "HelloWorldHolaMu")

	fmt.Println(encryptMsg)
	msg, _ := decryptCBC(key, encryptMsg)
	fmt.Println(msg) // Hello World
}
