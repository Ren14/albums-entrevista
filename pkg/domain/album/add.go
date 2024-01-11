package album

func (s *Service) Add(album Album) error {
	return s.StorageRepository.Add(album)
}
