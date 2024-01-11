package album

func (s *Service) GetByID(albumID string) (*Album, error) {
	return s.StorageRepository.GetByID(albumID)
}
