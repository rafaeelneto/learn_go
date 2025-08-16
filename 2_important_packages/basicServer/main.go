package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	http.HandleFunc("/", GetCEPHandler)
	http.ListenAndServe(":8080", nil)

}

func GetCEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addr, err := searchCep(cep)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(addr)

	// Longer alternative
	// jsonResponse, err := json.Marshal(addr)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// w.Write([]byte(jsonResponse))
}

func searchCep(cep string) (*AddressResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep)
	req, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var addressResponse AddressResponse
	err = json.Unmarshal(res, &addressResponse)
	if err != nil {
		return nil, err
	}

	return &addressResponse, nil
}
