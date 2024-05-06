package admin

import (
	"errors"

	"batch-allocation-system/models"
	"batch-allocation-system/service"
)

type admin struct {
	student service.Student
	admins  []models.Admin
	batches []models.Batch
}

func New(student service.Student) service.Admin {
	admins := make([]models.Admin, 0)
	batches := make([]models.Batch, 0)

	return &admin{student: student, admins: admins, batches: batches}
}

func (a *admin) Register(name, gender string) int {
	id := len(a.admins)

	a.admins = append(a.admins, models.Admin{ID: id, Name: name, Gender: gender})

	return id
}

func (a *admin) CreateBatch(id, capacity int, stream, timing string) (int, error) {
	index := a.findAdmin(id)
	if index == -1 {
		return -1, errors.New("admin not found")
	}

	batchID := len(a.batches) + 1

	a.batches = append(a.batches, models.Batch{ID: batchID, Capacity: capacity, InitialCapacity: capacity, Stream: stream,
		Timing: timing})

	return batchID, nil
}

func (a *admin) AllocateBatch(adminID, studentID int, allocationCriteria string) (int, error) {
	index := a.findAdmin(adminID)
	if index == -1 {
		return -1, errors.New("admin not found")
	}

	student, err := a.student.GetStudent(studentID)
	if err != nil {
		return -1, nil
	}

	var batchIndex int

	switch allocationCriteria {
	case "Gender Based":
		batchIndex = a.getBatchOnGender(student.Stream)
	case "Higher Capacity":
		batchIndex = a.getBatchOnCapacity(student.Stream)
	}

	if batchIndex == -1 {
		return -1, errors.New("no batch available")
	}

	err = a.student.AllocateBatch(studentID, a.batches[batchIndex].ID)
	if err != nil {
		return -1, err
	}

	a.batches[batchIndex].Capacity -= 1

	return a.batches[batchIndex].ID, nil
}

func (a *admin) findAdmin(id int) int {
	for i := range a.admins {
		if a.admins[i].ID == id {
			return i
		}
	}

	return -1
}

func (a *admin) findBatch(batchId int) int {
	for i := range a.batches {
		if a.batches[i].ID == batchId {
			return i
		}
	}

	return -1
}

func (a *admin) getBatchOnGender(stream string) int {
	for i, batch := range a.batches {
		if batch.Stream == stream && batch.Timing == "Morning" && batch.Capacity > 0 {
			return i
		}
	}

	for i, batch := range a.batches {
		if batch.Stream == stream && batch.Timing == "Afternoon" && batch.Capacity > 0 {
			return i
		}
	}

	for i, batch := range a.batches {
		if batch.Stream == stream && batch.Timing == "Evening" && batch.Capacity > 0 {
			return i
		}
	}

	return -1
}

func (a *admin) getBatchOnCapacity(stream string) int {
	maxCapacityBatchIndex := 0

	for i, batch := range a.batches {
		if (batch.Stream == stream && batch.InitialCapacity > a.batches[maxCapacityBatchIndex].InitialCapacity &&
			batch.Capacity != 0) || (a.batches[maxCapacityBatchIndex].Capacity == 0 && batch.Capacity != 0) {
			maxCapacityBatchIndex = i
		}
	}

	return maxCapacityBatchIndex
}
