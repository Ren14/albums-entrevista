package album

type StorageRepository interface {
	Add(album Album) error
	GetByID(albumID string) (*Album, error)
	GetAll() []Album
}

type CurrencyService interface {
	GetUSDPrice() float64
}

type Service struct {
	StorageRepository StorageRepository
	CurrencyService   CurrencyService
}

func NewService(storageRepo StorageRepository, currencyService CurrencyService) *Service {
	return &Service{
		StorageRepository: storageRepo,
		CurrencyService:   currencyService,
	}
}
