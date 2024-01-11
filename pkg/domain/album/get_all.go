package album

const defuaultPrice = 1.0

func (s *Service) GetAll() []Album {
	return s.StorageRepository.GetAll()
}

func (s *Service) GetAllUSD() []Album {
	albums := s.StorageRepository.GetAll()

	usdPrice := s.CurrencyService.GetUSDPrice() // No voy directo al repositorio, sino que al servicio que entiende sobre cotizaciones
	if usdPrice <= 0 {
		usdPrice = defuaultPrice
	}

	for i, a := range albums {
		albums[i].Price = a.Price / usdPrice // TODO: redondear a 2 decimales
	}

	return albums
}
