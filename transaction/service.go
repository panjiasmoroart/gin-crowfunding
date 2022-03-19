package transaction

type Service interface {
	GetTransactionsByCampaignID(camapignID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(camapignID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByCampaignID(camapignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
