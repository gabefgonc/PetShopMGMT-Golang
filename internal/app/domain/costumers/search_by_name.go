package costumers

type SearchPersonByName struct {
	repository Repository
}

func (s *SearchPersonByName) Execute(name string) ([]Costumer, error) {
	if len(name) < 2 {
		return nil, ErrSearchTermTooShort
	}

	costumers, err := s.repository.SearchByName(name)
	if err != nil {
		return nil, err
	}

	return costumers, nil
}
