package queries

import (
	"donation_app/app/models"
	"donation_app/platform/db"
	"log"
	"time"
)

func GetAllUsers() ([]models.User, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return []models.User{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		return []models.User{}, err
	}
	users := []models.User{}

	var c1 int
	var c2 time.Time
	var c3, c4, c5 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5)
		temp := models.User{c1, c2, c3, c4, c5}
		users = append(users, temp)
	}

	return users, nil
}

// func GetUserByID(id uuid.UUID) (models.User, error) {
// 	user := models.User{}

// 	query := `SELECT * FROM users WHERE id = $1`

// 	err := q.Get(&user, query, id)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func GetUserByEmail(email string) (models.User, error) {
// 	user := models.User{}

// 	query := `SELECT * FROM users WHERE email = $1`

// 	err := q.Get(&user, query, email)
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

func InsertUser(u models.User) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	if IsUserValid(u) {
		log.Println("User", u.Email, "already exist!")
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users(Created_At, Login, Email, Password) values($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	stmt.Exec(u.CreatedAt, u.Login, u.Email, u.Password)
	return nil
}

func IsUserValid(u models.User) bool {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Email = $1 \n", u.Email)
	if err != nil {
		return false
	}

	temp := models.User{}
	var c1 int
	var c2 time.Time
	var c3, c4, c5 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5)
		if err != nil {
			return false
		}
		temp = models.User{c1, c2, c3, c4, c5}
	}
	if u.Email == temp.Email {
		return true
	}

	return false
}
