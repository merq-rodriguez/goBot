package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Port    string `json:"PORT"`
	CertPem string `json:"CERT_PEM"`
	KeyPem  string `json:"KEY_PEM"`
}

//Server configurations
var config Config



func main() {
	loadConfig()
	http.HandleFunc("/", saludar)

	log.Printf("Server running in https://localhost%s", config.Port)
	err := http.ListenAndServeTLS(config.Port, config.CertPem, config.KeyPem, nil)
	if err != nil {
		log.Println(err)
	}
}

func saludar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func loadConfig(){
	log.Println("Starting reading of the configuration file")
	
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Err reading file: %v", err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("Err converting file: %v",err)
	}

	log.Println("Ending reading of the configuration file")

}