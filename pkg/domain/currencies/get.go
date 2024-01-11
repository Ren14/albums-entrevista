package currencies

func (s *CurrencyService) GetUSDPrice() float64 {
	return s.CurrencyRepo.GetPrice()
}
