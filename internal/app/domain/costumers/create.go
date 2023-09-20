package costumers

type CreateCostumer struct {
	repository Repository
}

func (c *CreateCostumer) Execute(name string, phoneNumber string) (*Costumer, error) {
	alreadyExists, err := c.repository.FindByPhoneNumber(phoneNumber)

	if alreadyExists != nil && err != nil {
		return nil, ErrAlreadyExists
	}

	costumer := &Costumer{
		Name:        name,
		PhoneNumber: phoneNumber,
	}

	err = c.repository.Create(costumer)
	if err != nil {
		return nil, err
	}

	return costumer, err
}

func NewCreateCostumer(repository Repository) *CreateCostumer {
	return &CreateCostumer{
		repository: repository,
	}
}
