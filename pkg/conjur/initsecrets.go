package conjur

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type secret struct {
	Path string `json:"id"`
}

var (
	emptyEnv     []string
	discoverPath string
)

func handleConjurAuthn(t string) string {

	if t == "" {
		emptyEnv = append(emptyEnv, "tempFile")
	}

	log.Println("Reading local authorization token.")

	token, err := os.ReadFile(t)

	if err != nil {

		log.Fatal(err)

	}

	return string(token)

}

func discover(t string, client *http.Client, query string) string {

	var sb strings.Builder

	authnUrl := discoverPath + query
	method := "GET"
	req, err := http.NewRequest(method, authnUrl, nil)

	if err != nil {

		log.Println(err)
		return "false"

	}
	tokenHeader := "Token token=\"" + t + "\""
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

func initSecrets(base string, token string, account string, safe string, query string) {

	if base == "" {
		log.Fatalf("%s not populated. Check your configuration and redeploy.", base)
	}
	if token == "" {
		log.Fatalf("%s not populated. Check your configuration and redeploy.", token)
	}
	if account == "" {
		log.Fatalf("%s not populated. Check your configuration and redeploy.", account)
	}
	if safe == "" {
		log.Fatalf("%s not populated. Check your configuration and redeploy.", safe)
	}
	if query == "" {
		log.Fatal("Base URI not populated. Check your configuration and redeploy.", query)
	}

	// Default Paths
	discoverPath := base + "/resources/" + account + "?kind=variable&search=" + safe + "/"
	retrievalPath := base + "/secrets?variable_ids="

	log.Println(discoverPath)
	log.Println(retrievalPath)

	log.Println("called.")
}
