package storage

import "sasha/Desktop/30.8.1-main/pkg/storage/postgres"

type Interface interface {
	GetTasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	UpdateTask(int, postgres.Task) error
	DeleteTask(int) error
}
