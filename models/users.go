package models

import "time"

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

// Create a user
func (user *User) Create() error {
	// Query
	insert, err := db.Exec("INSERT INTO users (username, first_name, last_name, actif, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", user.IDUser, user.FirstName, user.LastName, user.Actif, time.Now(), time.Now())
	if err != nil {
		return err
	}
	// Get last ID
	lastID, err := insert.LastInsertId()
	if err != nil {
		return err
	}
	// assign last id into User
	user.IDUser = int(lastID)

	// Find the new user
	if err = user.FindOne(); err != nil {
		return err
	}

	return nil
}

// Update a user
func (user *User) Update() error {
	// Query
	_, err := db.Exec("UPDATE FROM users SET first_name = ?, last_name = ?, actif = ?, updated_at = ? WHERE id_user = ?", user.FirstName, user.LastName, user.Actif, time.Now(), user.IDUser)
	if err != nil {
		return err
	}

	return nil
}
