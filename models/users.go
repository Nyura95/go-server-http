package models

// User models
type User struct {
	IDUser    int
	Username  string
	FirstName string
	LastName  string
	Actif     int
}

// FindOne user with the userID
func (user *User) FindOne() error {
	// Query
	err := db.QueryRow("SELECT * FROM users WHERE id_user = ? LIMIT 1", user.IDUser).Scan(&user.IDUser, &user.Username, &user.FirstName, &user.LastName)
	if err != nil {
		return err
	}

	return nil
}

// FindAllActif user
func FindAllActif() ([]*User, error) {
	// Create tabs User
	users := make([]*User, 0)
	// Query
	rows, err := db.Query("SELECT * FROM users WHERE actif = 1")
	if err != nil {
		// If error return tab empty
		return users, err
	}
	// Close connect after the function or error
	defer rows.Close()

	for rows.Next() {
		// Create new user
		user := new(User)
		// Assign the values
		err := rows.Scan(&user.IDUser, &user.Username, &user.FirstName, &user.LastName, &user.Actif)
		if err != nil {
			return users, err
		}
		// Add user
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
