package shared

import (
	"log"

	"github.com/google/uuid"
)

func GetID() uuid.UUID {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatal("Erro", err)
	}

	return uuid
}

func ParseIDString(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func GetUuidEmpty() uuid.UUID {
	return uuid.Nil
}
