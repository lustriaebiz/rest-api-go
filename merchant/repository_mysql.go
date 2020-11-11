package merchant

import (
	"context"
	"fmt"
	"log"
	"time"

	"api-new/config"
	"api-new/models"
)

const (
	table          = "merchant"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Merchant, error) {

	var merchants []models.Merchant

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var merchant models.Merchant
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&merchant.ID,
			&merchant.Name,
			&merchant.Telp,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		merchant.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		merchant.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func Insert(ctx context.Context, merchants models.Merchant) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, telp, created_at, updated_at) values('%v','%v','%v','%v')", table,
		merchants.Name,
		merchants.Telp,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
