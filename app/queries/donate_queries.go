package queries

import (
	"donation_app/app/models"
	"donation_app/platform/db"
	"log"
	"time"
)

func ListAllDonates(loginStrimer string) ([]models.Donate, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return []models.Donate{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM donates WHERE Logintodonate = $1 \n", loginStrimer)
	if err != nil {
		log.Println(err)
		return []models.Donate{}, err
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

	return all, nil
}

func GetAllDonates() ([]models.Donate, error) {
	db, err := db.ConnectPostgres()
	if err != nil {
		db.Close()
		return []models.Donate{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM donates \n")
	if err != nil {
		return []models.Donate{}, err
	}
	donates := []models.Donate{}

	var c1, c6 int
	var c2 time.Time
	var c3, c4, c5 string

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := models.Donate{c1, c2, c3, c4, c5, c6}
		donates = append(donates, temp)
	}

	return donates, nil
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

func DeleteDonate(ID int) error {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return err
	}
	defer db.Close()

	//check is the user ID exist
	t := FindDonateID(ID)
	if t.ID == 0 {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM donates WHERE ID = $1")
	if err != nil {
		log.Println("DeleteUser:", err)
		return err
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println("DeleteUser:", err)
		return err
	}

	return nil
}

func FindDonateID(ID int) models.Donate {
	db, err := db.ConnectPostgres()
	if err != nil {
		log.Println("Cannot connect to PostreSQL!")
		db.Close()
		return models.Donate{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM donates WHERE ID = $1\n", ID)
	if err != nil {
		log.Println("Query:", err)
		return models.Donate{}
	}
	defer rows.Close()

	d := models.Donate{}
	var c1, c6 int
	var c2 time.Time
	var c3, c4, c5 string

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return models.Donate{}
		}
		d = models.Donate{c1, c2, c3, c4, c5, c6}
	}
	return d
}
