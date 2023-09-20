package repository

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gabefgonc/PetShopMGMT/internal/app/domain/costumers"
)

const (
	InsertOneCostumerSQL            = "INSERT INTO costumers (name, phone_number) VALUES ($1, $2) RETURNING id"
	FindOneCustomerByIdSQL          = "SELECT (id, name, phone_number, created_at) FROM costumers WHERE id = $1"
	FindOneCustomerByPhoneNumberSQL = "SELECT (id, name, phone_number, created_at) FROM costumers WHERE phone_number = $1"
	SearchCostumerByNameSQL         = "SELECT (id, name, phone_number, created_at) FROM costumers WHERE name LIKE '%' || $1 || '%'"
)

type CostumersRepository struct {
	db *pgxpool.Pool
}

func (c *CostumersRepository) Create(costumer *costumers.Costumer) (id int64, err error) {
	err = c.db.QueryRow(context.Background(), InsertOneCostumerSQL, costumer.Name, strings.Replace(costumer.PhoneNumber, " ", "", -1)).
		Scan(&id)

	return
}

func (c *CostumersRepository) FindByID(id string) (costumer *costumers.Costumer, err error) {
	err = c.db.QueryRow(context.Background(), FindOneCustomerByIdSQL, id).
		Scan(&costumer.ID, &costumer.Name, &costumer.PhoneNumber, &costumer.CreatedAt)
	return
}

func (c *CostumersRepository) FindByPhoneNumber(
	phoneNumber string,
) (costumer *costumers.Costumer, err error) {
	err = c.db.QueryRow(context.Background(), FindOneCustomerByPhoneNumberSQL, strings.Replace(phoneNumber, " ", "", -1)).
		Scan(&costumer.ID, &costumer.Name, &costumer.PhoneNumber, &costumer.CreatedAt)
	return
}

func (c *CostumersRepository) SearchByName(
	name string,
) ([]costumers.Costumer, error) {
	rows, err := c.db.Query(
		context.Background(),
		SearchCostumerByNameSQL,
		strings.ToLower(name),
	)
	if err != nil {
		log.Errorf("Error executing search: %v", err)
		return nil, err
	}

	defer rows.Close()

	costumersResult := make([]costumers.Costumer, 0)

	for rows.Next() {
		var costumer costumers.Costumer

		err := rows.Scan(
			&costumer.ID,
			&costumer.Name,
			&costumer.PhoneNumber,
			&costumer.CreatedAt,
		)
		if err != nil {
			log.Errorf("Error scanning row: %v", err)
			return nil, err
		}

		costumersResult = append(costumersResult, costumer)
	}

	return costumersResult, nil
}

func NewCostumersRepository(db *pgxpool.Pool) *CostumersRepository {
	return &CostumersRepository{
		db: db,
	}
}
