package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return data, err
}

func encrypt(plainstring, keystring string) string {
	plaintext := []byte(plainstring)
	key := []byte(keystring)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(reflect.TypeOf(key))
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return string(ciphertext)
}

func decrypt(cipherstring string, keystring string) string {
	ciphertext := []byte(cipherstring)
	key := []byte(keystring)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("text too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}

func main() {
	key := "YCbM/d49s5CiDQzr/bXLZg=="

	for {
		fmt.Print("What would you like to do ? \n")
		line := readline()
		fmt.Println("so you want to " + line)

		switch line {
		case "help":
			fmt.Println("You can:\nencrypt\ndecrypt\nexit")
		case "exit":
			os.Exit(0)
		case "encrypt":
			fmt.Print("What would you like to encrypt? ")
			line2 := readline()
			ciphertext := encrypt(line2, key)
			fmt.Print("what is the file name: ")
			line3 := readline()
			writeToFile(ciphertext, line3)
			fmt.Println("wrote to file")
		case "decrypt":
			fmt.Print("What is the name of the file to decrypt? ")
			line2 := readline()
			if ciphertext, err := readFromFile(line2); err != nil {
				fmt.Println("File is not found")
			} else {
				plaintext := decrypt(string(ciphertext), key)
				fmt.Println(plaintext)
			}
		}
	}
}
