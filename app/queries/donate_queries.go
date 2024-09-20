package queries

import (
	"donation_app/app/models"
	"donation_app/platform/db"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DonatesQueries struct {
	*sqlx.DB
}

// ListAllMessages if for returning all messages from the database table
func ListAllDonates(loginStrimer string) ([]models.Donate, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return []models.Donate{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM donates WHERE Logintodonate = $1 \n", loginStrimer)
	if err != nil {
		log.Println(err)
		return []models.Donate{}
	}

	all := []models.Donate{}
	var c1, c6 int
	var c2 time.Time
	var c4, c3, c5 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := models.Donate{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	return all
}

func (q *DonatesQueries) GetDonateByID(id uuid.UUID) (models.Donate, error) {
	donate := models.Donate{}

	query := `SELECT * FROM donates WHERE id = $1`

	err := q.Get(&donate, query, id)
	if err != nil {
		return donate, err
	}

	return donate, nil
}

func InsertDonate(d models.Donate) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO donates(Created_At, Loginwhodonate, Logintodonate, Message, Summary) values($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	stmt.Exec(d.CreatedAt, d.LoginWhoDonate, d.LoginToDonate, d.Message, d.Summary)
	return nil
}
