package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// ID representa um identificador único universal (UUID).
type ID struct {
	value uuid.UUID
}

// NewID cria um novo ID com um UUID gerado automaticamente.
func NewID() ID {
	return ID{value: uuid.New()}
}

// ParseID cria um ID a partir de uma string.
// Retorna erro caso o formato não seja um UUID válido.
func ParseID(s string) (ID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID{value: parsed}, nil
}

// String retorna o ID no formato string.
func (id ID) String() string {
	return id.value.String()
}

// Equal compara dois IDs e retorna verdadeiro se forem iguais.
func (id ID) Equal(other ID) bool {
	return id.value == other.value
}

// IsEmpty verifica se o ID está vazio (UUID zero).
func (id ID) IsEmpty() bool {
	return id.value == uuid.Nil
}

// MarshalJSON converte o ID para JSON como uma string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON lê um ID de uma string no JSON.
func (id *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ID deve ser uma string: %w", err)
	}
	parsed, err := uuid.Parse(s)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}
	id.value = parsed
	return nil
}

// Value implementa driver.Valuer
func (id ID) Value() (driver.Value, error) {
	if id.IsEmpty() {
		return nil, nil
	}
	return id.value.String(), nil
}

// Scan implementa sql.Scanner
func (id *ID) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("ID precisa ser uma string")
	}
	parsed, err := uuid.Parse(str)
	if err != nil {
		return err
	}
	id.value = parsed
	return nil
}
