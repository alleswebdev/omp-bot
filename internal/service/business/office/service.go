package office

import (
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type DummyOfficeService struct {
	allEntities []business.Office
}

func NewDummyOfficeService() *DummyOfficeService {
	return &DummyOfficeService{allEntities: []business.Office{
		{
			Id:          1,
			Name:        "One",
			Description: "Office one",
		},
		{
			Id:          2,
			Name:        "Two",
			Description: "Office two",
		},
		{
			Id:          3,
			Name:        "three",
			Description: "Office tree",
		},
		{
			Id:          4,
			Name:        "four",
			Description: "Office four",
		},
		{
			Id:          5,
			Name:        "five",
			Description: "Office 5",
		},
		{
			Id:          6,
			Name:        "six",
			Description: "Office 6",
		},
		{
			Id:          7,
			Name:        "seven",
			Description: "Office 7",
		},
		{
			Id:          8,
			Name:        "eight",
			Description: "Office 8",
		},
		{
			Id:          9,
			Name:        "nine",
			Description: "Office 9",
		},
	}}
}

func (s *DummyOfficeService) List(cursor uint64, limit uint64) ([]business.Office, error) {
	if len(s.allEntities) == 0 {
		return nil, ErrorEmptyList
	}

	// когда сущеностей осталось меньше, чем лимит на выдачу, но их надо показать
	if uint64(len(s.allEntities)) > cursor && uint64(len(s.allEntities)) < cursor+limit {
		return s.allEntities[cursor:], nil
	}

	if uint64(len(s.allEntities)) <= cursor {
		return nil, ErrorOutRange
	}

	return s.allEntities[cursor : cursor+limit], nil
}

func (s *DummyOfficeService) Describe(officeId uint64) (*business.Office, error) {
	index, ok := s.getIndexById(officeId)

	if !ok {
		return nil, ErrorNotFound
	}

	return &s.allEntities[index], nil
}

func (s *DummyOfficeService) Remove(officeId uint64) (bool, error) {
	index, ok := s.getIndexById(officeId)

	if !ok {
		return false, ErrorNotFound
	}

	s.allEntities = append(s.allEntities[:index], s.allEntities[index+1:]...)

	return true, nil
}

func (s *DummyOfficeService) Create(o business.Office) (uint64, error) {
	o.Id = s.getNextEntityId()
	s.allEntities = append(s.allEntities, o)

	return o.Id, nil
}

func (s *DummyOfficeService) Update(officeId uint64, office business.Office) error {
	if len(s.allEntities) == 0 {
		return ErrorEmptyList
	}

	index, ok := s.getIndexById(officeId)

	if !ok {
		return ErrorNotFound
	}

	s.allEntities[index].Name = office.Name
	s.allEntities[index].Description = office.Description

	return nil
}

func (s *DummyOfficeService) getIndexById(officeId uint64) (int, bool) {
	if len(s.allEntities) == 0 {
		return 0, false
	}

	for k, entity := range s.allEntities {
		if entity.Id == officeId {
			return k, true
		}
	}

	return 0, false
}

func (s *DummyOfficeService) getNextEntityId() uint64 {
	maxId := uint64(1)

	if len(s.allEntities) == 0 {
		return maxId
	}

	for _, entity := range s.allEntities {
		if entity.Id > maxId {
			maxId = entity.Id
		}
	}

	return maxId + 1
}
