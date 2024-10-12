package memdb

import "sasha/Desktop/30.8.1-main/pkg/storage/postgres"

type DB []postgres.Task

func (db DB) GetTasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) UpdateTask(int, postgres.Task) error {
	return nil
}

func (db DB) DeleteTask(int) error {
	return nil
}
