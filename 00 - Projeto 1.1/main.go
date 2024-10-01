//Exercício 1: Validação de CEP
//Implemente uma validação para verificar se o CEP contém exatamente 8 dígitos.
//Caso contrário, retorne um erro de requisição inválida (400 Bad Request).

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode"
)

type ViaCEP struct {
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

func isNumeric(cep string) bool {
	for _, ch := range cep {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found!\n"))
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if len(cepParam) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP must have 8 digits"))
		return
	} else if isNumeric(cepParam) != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP must be numeric"))
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Forma extensa
	// result, err := json.Marshal(cep)
	// if err != nil{
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	// Forma mais prática
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
