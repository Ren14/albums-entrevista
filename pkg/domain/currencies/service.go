package currencies

type CurrencyRepo interface {
	GetPrice() float64
}

type CurrencyService struct {
	CurrencyRepo CurrencyRepo
}

func NewService(currencyRepo CurrencyRepo) *CurrencyService {
	return &CurrencyService{
		CurrencyRepo: currencyRepo,
	}
}
