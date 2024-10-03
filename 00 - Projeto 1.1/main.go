//Exercício 1: Validação de CEP
//Implemente uma validação para verificar se o CEP contém exatamente 8 dígitos.
//Caso contrário, retorne um erro de requisição inválida (400 Bad Request).

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

var cache = make(map[string]ViaCEP)
var errors = make([]string, 5)

func isNumeric(cep string) bool {
	for _, ch := range cep {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Servidor executando normalmente.")
}

func CreateLogFile() {
	t := time.Now()
	file_name := t.Format("2006-01-02 15:04:05") + ".txt"

	log_file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}

	for i := len(errors) - 1; i >= 0; i-- {
		_, err := log_file.Write([]byte(errors[i] + "\n"))
		if err != nil {
			panic(err)
		}
	}

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/status" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found!\n"))
		t := time.Now()
		error_entry := t.Format("2006-01-02 15:04:05") + "-" + "Not Found"
		errors = append(errors, error_entry)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP cannot be an empty value"))
		t := time.Now()
		error_entry := t.Format("2006-01-02 15:04:05") + "-" + "CEP cannot be an empty value"
		errors = append(errors, error_entry)
		return
	} else if len(cepParam) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP must have 8 digits"))
		t := time.Now()
		error_entry := t.Format("2006-01-02 15:04:05") + "-" + "CEP must have 8 digits"
		errors = append(errors, error_entry)
		return
	} else if isNumeric(cepParam) != true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP must be numeric"))
		t := time.Now()
		error_entry := t.Format("2006-01-02 15:04:05") + "-" + "CEP must be numeric"
		errors = append(errors, error_entry)
		return
	}

	found := cache[cepParam]
	if (ViaCEP{}) != found {
		json.NewEncoder(w).Encode(found)
	} else {

		cep, error := BuscaCep(cepParam)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		//for key, value := range cache {
		//	fmt.Println(key, " - ", value)
		//}

		// Forma extensa
		// result, err := json.Marshal(cep)
		// if err != nil{
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		// w.Write(result)

		// Forma mais prática
		json.NewEncoder(w).Encode(cep)
		t := time.Now()
		error_entry := t.Format("2006-01-02 15:04:05") + "-" + "Success!"
		errors = append(errors, error_entry)
	}

	CreateLogFile()

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

	cache[cep] = c

	return &c, nil
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.HandleFunc("/status", StatusHandler)
	http.ListenAndServe(":8080", nil)

}
