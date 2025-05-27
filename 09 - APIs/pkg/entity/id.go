package entity

import "github.com/google/uuid"

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func (id ID) String() string {
	return id.value.String()
}

// Como o ID vira de um JSON
// importante validar com o ParseID
// para garantir que se trata de um ID
// valido
/*
func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
*/
