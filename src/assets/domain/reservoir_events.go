package domain

import (
	"time"

	"github.com/satori/go.uuid"
)

type ReservoirCreated struct {
	UID         uuid.UUID
	Name        string
	WaterSource WaterSource
	FarmUID     uuid.UUID
	CreatedDate time.Time
}

type ReservoirWaterSourceChanged struct {
	WaterSource WaterSource
}

type ReservoirNoteAdded struct {
	ReservoirUID uuid.UUID
	UID          uuid.UUID
	Content      string
	CreatedDate  time.Time
}

type ReservoirNoteRemoved struct {
	ReservoirUID uuid.UUID
	UID          uuid.UUID
}