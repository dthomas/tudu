package datastore

// Todo Model
type Todo struct {
	ID        int    `db:"id" json:"id"`
	Completed bool   `db:"completed" json:"completed"`
	Notes     string `db:"notes,omitempty" json:"notes,omitempty"`
}

// Validate a todo
func (todo *Todo) Validate() (map[string][]string, bool) {
	errors := make(map[string][]string)

	if todo.Notes == "" {
		errors["notes"] = append(errors["notes"], "can't be blank")
	}

	if len(todo.Notes) < 4 {
		errors["notes"] = append(errors["notes"], "lenght must be atleast 4")
	}

	if len(todo.Notes) > 192 {
		errors["notes"] = append(errors["notes"], "length must be atmost 192")
	}

	return errors, len(errors) == 0
}

// Get : Fetch a Todo by ID
func (todo *Todo) Get(id int) error {
	err := DB.QueryRowx("SELECT * FROM todos WHERE id = $1", id).StructScan(todo)
	if err != nil {
		return err
	}
	return nil
}

// GetAllTodo : Fetch all todo items
func GetAllTodo() ([]Todo, error) {
	todos := []Todo{}
	rows, err := DB.Queryx("SELECT * FROM todos")
	if err != nil {
		return todos, err
	}
	defer rows.Close()
	for rows.Next() {
		todo := Todo{}
		err := rows.StructScan(&todo)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

// Save a new todo
func (todo *Todo) Save() error {
	stmt := `INSERT INTO todos(completed, notes) VALUES($1, $2) RETURNING id`
	err := DB.QueryRow(stmt, todo.Completed, todo.Notes).Scan(&todo.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update a todo
func (todo *Todo) Update() error {
	stmt := `UPDATE todos SET completed = $1, notes = $2 WHERE id = $3`
	_, err := DB.Query(stmt, todo.Completed, todo.Notes, todo.ID)
	return err
}

// Delete a todo
func (*Todo) Delete(id int) error {
	stmt := `DELETE FROM todos WHERE id = $1`
	_, err := DB.Query(stmt, id)
	return err
}
