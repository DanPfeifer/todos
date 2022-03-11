package resources

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	Token string
}


func (d database) Login(email string, password string) (u User, err error) {

	err = d.db.QueryRow("SELECT id, first_name, last_name, email FROM `user` WHERE email=? AND password=?", email, password).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)

	if err != nil {
		return
	}

	return
}

func (d database) FindByEmail(email string) (u User, err error) {

	err = d.db.QueryRow("SELECT id, first_name, last_name, email, token FROM `user` WHERE email=?", email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Token)

	if err != nil {
		return
	}

	return
}

func (d database) getOneUser(id int) (u User, err error) {

	err = d.db.QueryRow("SELECT id, first_name, last_name, email, token FROM `user` WHERE id=?", id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Token)

	if err != nil {
		return
	}

	return
}

func (d database) SetToken(user User, token string) (u User, err error) {

	stmt, err := d.db.Prepare("UPDATE `user` SET token=? WHERE id=?")
	if err != nil {
		return
	}

	_, err = stmt.Exec(token, user.ID)
	if err != nil {
		return
	}
	u = User{
		ID: user.ID,
		FirstName:user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Token: token,
	}

	return
}