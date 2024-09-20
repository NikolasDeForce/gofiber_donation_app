package queries

import (
	"donation_app/app/models"
	"donation_app/platform/db"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DonatesQueries struct {
	*sqlx.DB
}

func (q *DonatesQueries) GetAllDonates() ([]models.Donate, error) {
	donates := []models.Donate{}

	query := `SELECT * FROM donates`

	err := q.Get(&donates, query)
	if err != nil {
		return donates, err
	}

	return donates, nil
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
