// Ref: https://www.systutorials.com/how-to-generate-rsa-private-and-public-key-pair-in-go-lang/
package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/public", sendPublicKey)

	fmt.Println("Server is at 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func sendPublicKey(w http.ResponseWriter, r *http.Request) {

	//publicKeyBytes := genKeys()
	publicKey := fetchPubKey()

	body := map[string]interface{}{"pubkey": publicKey, "msg": "success"}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func fetchPubKey() []byte {

	pubkey, err := os.ReadFile("./pubkey.pem")
	if err != nil {
		fmt.Println("Cannot read pubkey.pem")
	}

	block, _ := pem.Decode([]byte(pubkey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed: " + err.Error())
	}
	rsaPubKey, _ := pub.(*rsa.PublicKey)
	if err != nil {
		fmt.Println("Unable to convert to rsa public key")
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(rsaPubKey)
	if err != nil {
		fmt.Println("Cannot Marshal pubkey")
	}
	return publicKeyBytes
}

/*
func genKeys() []byte {
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
	return publicKeyBytes
}
*/
