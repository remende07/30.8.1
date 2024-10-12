DROP TABLE IF EXISTS tasks_labels, users, tasks, labels;

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    opened BIGINT NOT NULL DEFAULT extract(epoch from now()),
    closed BIGINT DEFAULT 0,
    author_id INTEGER REFERENCES users(id),
    assigned_id INTEGER REFERENCES users(id),
    title TEXT,
    content TEXT
);

CREATE TABLE labels(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE tasks_labels(
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);

INSERT INTO users (id, name) VALUES (1, 'User1'), (2, 'User2');

INSERT INTO labels (id, name) VALUES (1, 'Label1'), (2, 'Label2');