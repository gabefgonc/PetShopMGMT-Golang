package costumers

import "time"

type Costumer struct {
	ID          int64
	Name        string
	PhoneNumber string
	CreatedAt   time.Time
}
