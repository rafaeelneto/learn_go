package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AddressResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, args := range os.Args[1:] {
		url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", args)
		req, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		println(string(res))

		var data AddressResponse

		err = json.Unmarshal(res, &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Logradouro)

		file, err := os.Create("cep.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = json.NewEncoder(file).Encode(data)
		if err != nil {
			panic(err)
		}
	}
}
