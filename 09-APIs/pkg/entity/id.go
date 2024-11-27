package entity

import "github.com/google/uuid"

// seguindo os padroes de DDD
// cada objeto criado deve gerar
// valor ao sistema
// assim o ID deve gerar novos IDs
// sempre que necessario

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
