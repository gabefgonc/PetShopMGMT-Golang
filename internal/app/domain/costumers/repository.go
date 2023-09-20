package costumers

type Repository interface {
	Create(costumer *Costumer) error
	FindByID(id string) (*Costumer, error)
	FindByPhoneNumber(phoneNumber string) (*Costumer, error)
	SearchByName(term string) ([]Costumer, error)
}
