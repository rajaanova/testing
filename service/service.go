package service

import "codebase/intve/storage"

type Service struct {
	Store storage.MapStorage
}

func (s *Service) GetUser(id string) (*storage.User, error) {
	return s.Store.GetUserById(id)

}
