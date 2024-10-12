package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(connstr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}

	return &s, nil
}

type Task struct {
	ID         int
	Opened     int
	Closed     int
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

func (s *Storage) NewTask(task Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks(title, content)
		VALUES ($1, $2) RETURNING id
		`,
		task.Title, task.Content).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetTasks(labelID int, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
			FROM tasks, tasks_labels
			WHERE
				($1 = 0 OR (label_id = $1 AND task_id = id)) AND
				($2 = 0 OR author_id = $2)
			ORDER BY id;
		`,
		labelID, authorID)

	if err != nil {
		return nil, err
	}

	var tasks []Task

	for rows.Next() {
		var task Task
		err = rows.Scan(
			&task.ID,
			&task.Opened,
			&task.Closed,
			&task.AuthorID,
			&task.AssignedID,
			&task.Title,
			&task.Content,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, rows.Err()
}

func (s *Storage) UpdateTask(taskID int, task Task) error {
	_, err := s.db.Exec(context.Background(), `
		UPDATE tasks
		SET title = $1, content = $2, closed = $3
		WHERE id = $4`, task.Title, task.Content, task.Closed, taskID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteTask(taskID int) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks_labels WHERE task_id = $1
		`, taskID)

	if err != nil {
		return err
	}

	_, err = s.db.Exec(context.Background(), `
		DELETE FROM tasks WHERE id = $1
		`, taskID)

	if err != nil {
		return err
	}

	return nil
}
