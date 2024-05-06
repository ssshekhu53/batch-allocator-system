package service

import "batch-allocation-system/models"

type Student interface {
	Register(name, gender string) int
	Enroll(id int, stream string) error
	AllocateBatch(studentID, batchID int) error
	GetStudent(id int) (*models.Student, error)
}

type Admin interface {
	Register(name, gender string) int
	CreateBatch(id, capacity int, stream, timing string) (int, error)
	AllocateBatch(adminID, studentID int, allocationCriteria string) (int, error)
}
