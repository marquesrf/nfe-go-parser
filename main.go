package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"nfe-parser/models"
	"os"
)

func main() {

	data, err := os.Open("NFE.xml")
	if err != nil {
		fmt.Printf("Deu ruim ao abrir o arquivo: %s", err.Error())
	}
	defer data.Close()

	bData, err := io.ReadAll(data)
	if err != nil {
		fmt.Printf("Deu ruim ao ler o arquivo: %s", err.Error())
	}

	var nfeProc models.NfeProc

	err = xml.Unmarshal(bData, &nfeProc)
	if err != nil {

	}

	nfeJson, err := json.MarshalIndent(nfeProc, "", " ")
	if err != nil {
		fmt.Printf("Deu ruim ao tentar imprimir em JSON: %s\n", err.Error())
		fmt.Println("Imprimindo no formato padrão...")
		fmt.Printf("%+v\n", nfeProc)
	}

	fmt.Printf("Aqui estão as principais informações da sua NFe: %s\n", string(nfeJson))

}
