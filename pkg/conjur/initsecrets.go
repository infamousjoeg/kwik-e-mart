package conjur

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	enc "encoding/base64"
)

type secret struct {
	Path string `json:"id"`
}

var (
	USER         string
	PASS         string
	ADDR         string
	DBNAME       string
	PORT         string
)

func httpClient() *http.Client {

	// Demo app

	log.Println("This is a demo and uses self signed certificate validation only. Please update if using this for other purposes.")

	c := &http.Client{

		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	log.Println("Successfully built http client.")

	return c

}

func handleConjurAuthn(t string) string {

	log.Println("Reading local authorization token.")

	rawToken, err := os.ReadFile(t)
	if err != nil {

		log.Fatal(err)

	}

	log.Println("Encoding Token")

	token := enc.StdEncoding.EncodeToString([]byte(rawToken))

	return string(token)

}

func discover(authn string, client *http.Client, query string, uri string) string {

	var sb strings.Builder

	authnUrl := uri + query
	method := "GET"
	req, err := http.NewRequest(method, authnUrl, nil)

	if err != nil {

		log.Println(err)
		return "false"

	}
	tokenHeader := "Token token=\"" + authn + "\""
	req.Header.Add("Authorization", tokenHeader)

	res, err := client.Do(req)
	if err != nil {

		log.Println(err)
		return "false"

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {

		log.Println(err)
		return "false"

	}

	secretPaths := []secret{}
	jsonErr := json.Unmarshal(body, &secretPaths)
	if jsonErr != nil {

		log.Fatal(jsonErr)

	}

	log.Println("Successfully found secrets, preparing paths")

	for _, values := range secretPaths {

		sb.WriteString(values.Path + ",")

	}

	log.Println("Converting buffer to string.")
	dirtyVariables := sb.String()
	log.Println("Cleaning variables.")
	cleanVariables := strings.TrimSuffix(dirtyVariables, ",")

	return cleanVariables

}

func initSecretData(authn string, client *http.Client, paths string, uri string) [5]string {

	authnUrl := uri + paths
	method := "GET"
	req, err := http.NewRequest(method, authnUrl, nil)

	if err != nil {

		log.Fatal(err)

	}
	tokenHeader := "Token token=\"" + authn + "\""
	req.Header.Add("Authorization", tokenHeader)

	res, err := client.Do(req)
	if err != nil {

		log.Fatal(err)

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {

		log.Fatal(err)

	}

	responseDecoded := make(map[string]string)

	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&responseDecoded)
	if err != nil {
		log.Fatal("Error coding JSON response.")
	}

	for k, v := range responseDecoded {

		lastIndex := strings.LastIndex(k, "/")

		attrChk := k[lastIndex+1:]

		if attrChk == "username" {
			USER = v
		}
		if attrChk == "password" {
			PASS = v
		}
		if attrChk == "Port" {
			PORT = v
		}
		if attrChk == "Database" {
			DBNAME = v
		}
		if attrChk == "address" {
			ADDR = v
		}
	}

	var payload [5]string

	/*
	*	Working with discovery I build a mapped array as follows:
	*	payload[1] := USER
	*	payload[2] := PASS
	*
	* 	This is an example of implementing discovery and can be retrofit for any need.
	*/

	payload[0] = USER
	payload[1] = PASS
	payload[2] = PORT
	payload[3] = DBNAME
	payload[4] = ADDR

	return payload
}


//Returns data based on the query provided. This will discover secrets from the provided safe and return data as implemented.
func GetData(base string, token string, account string, safe string, query string) [5]string {

	log.Println("Validating Variables.")

	if base == "" {
		log.Fatal("Base URI not populated. Check your configuration and redeploy.")
	}
	if token == "" {
		log.Fatal("Token Path not populated. Check your configuration and redeploy.")
	}
	if account == "" {
		log.Fatal("Conjur Account not populated. Check your configuration and redeploy.")
	}
	if safe == "" {
		log.Fatal("Safe not populated. Check your configuration and redeploy.")
	}
	if query == "" {
		log.Fatal("Query not populated. Check your configuration and redeploy.")
	}

	// Set up http Client
	hClient := httpClient()

	// Default Paths
	discoverPath := base + "/resources/" + account + "?kind=variable&search=" + safe + "/"
	retrievalPath := base + "/secrets?variable_ids="

	localToken := handleConjurAuthn(token)

	paths := discover(localToken, hClient, query, discoverPath)

	secrets := initSecretData(localToken, hClient, paths, retrievalPath)

	return secrets
}
