package model

import "github.com/google/uuid"

type UUID string

func NewUUID(uuid string) UUID {
	return UUID(uuid)
}

func (u UUID) Validate() error {
	_, err := uuid.Parse(string(u))

	return err
}
