package resources

type ToDo struct {
	ID int
	Title string
	Message string
	Priority int
	Status int
	InsertDate string
	ModifiedDate string
	UserID int
	DueDate string
}

func (d database) GetAllTodos(user User) (results []ToDo, err error) {
	stmt, err := d.db.Prepare("SELECT id, title, message, priority, status, insert_date, modified_date, user_id, due_date  FROM `todo` WHERE user_id = ? ORDER BY due_date ASC")

	if err != nil {
		return
	}

	rows, err := stmt.Query(user.ID)

	if err != nil {
		return
	}
	defer rows.Close()

	t := ToDo{}
	for rows.Next() {

		err = rows.Scan(&t.ID, &t.Title, &t.Message, &t.Priority, &t.Status, &t.InsertDate, &t.ModifiedDate, &t.UserID, &t.DueDate)

		if err != nil {
			return
		}
		results = append(results, t)

	}
	return
}

func (d database) GetOneTodo(id int) (t ToDo, err error) {

	err = d.db.QueryRow("SELECT id, title, message, priority, status, insert_date, modified_date, user_id, due_date  FROM `todo` WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Message, &t.Priority, &t.Status, &t.InsertDate, &t.ModifiedDate, &t.UserID, &t.DueDate)

	if err != nil {
		return
	}

	return
}

func (d database) CreateTodo(t ToDo) (todo ToDo, err error) {

	stmt, err := d.db.Prepare("INSERT INTO `todo` (title, message, priority, status, insert_date, modified_date, user_id, due_date) VALUES (?, ?,?, ?, NOW(), NOW(), ?, ? )")

	if err != nil {
		return
	}
	res, err := stmt.Exec(t.Title, t.Message, t.Priority, t.Status, t.UserID, t.DueDate)

	if err != nil {
		return
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return
	}
	todo, err = d.GetOneTodo(int(lastId))

	if err != nil {
		return
	}
	return
}

func (d database) UpdateTodo(t ToDo, id int) ( todo ToDo, err error) {

	stmt, err := d.db.Prepare("UPDATE `todo` SET title=?, message=?, priority=?, status=?, modified_date=NOW(), due_date=? WHERE id=?")

	if err != nil {
		return
	}
	_, err = stmt.Exec(t.Title, t.Message, t.Priority, t.Status, t.DueDate, id)

	if err != nil {
		return
	}

	todo, err = d.GetOneTodo(id)

	if err != nil {
		return
	}

	return
}

func (d database) DeleteTodo(id int) ( err error) {

	stmt, err := d.db.Prepare("DELETE FROM `todo` WHERE id = ?")

	if err != nil {
		return
	}
	_, err = stmt.Exec(id)

	if err != nil {
		return
	}

	return
}
