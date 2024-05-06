package student

import (
	"batch-allocation-system/models"
	"batch-allocation-system/service"
	"errors"
)

type student struct {
	students []models.Student
}

func New() service.Student {
	students := make([]models.Student, 0)

	return &student{students: students}
}

func (s *student) Register(name, gender string) int {
	id := len(s.students) + 1
	s.students = append(s.students, models.Student{ID: id, Name: name, Gender: gender})

	return id
}

func (s *student) Enroll(id int, stream string) error {
	index := s.findStudent(id)
	if index == -1 {
		return errors.New("student not found")
	}

	s.students[index].Stream = stream

	return nil
}

func (s *student) AllocateBatch(studentID, batchID int) error {
	index := s.findStudent(studentID)
	if index == -1 {
		return errors.New("student not found")
	}

	s.students[index].Batch = batchID

	return nil
}

func (s *student) GetStudent(id int) (*models.Student, error) {
	index := s.findStudent(id)
	if index == -1 {
		return nil, errors.New("student not found")
	}

	return &s.students[index], nil
}

func (s *student) findStudent(id int) int {
	for i := range s.students {
		if s.students[i].ID == id {
			return i
		}
	}

	return -1
}
