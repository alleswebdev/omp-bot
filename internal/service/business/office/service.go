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
	if len(s.allEntities) == 0 {
		return nil, ErrorEmptyList
	}

	for _, entity := range s.allEntities {
		if entity.Id == officeId {
			return &entity, nil
		}
	}

	return nil, ErrorNotFound
}

func (s *DummyOfficeService) Remove(officeId uint64) (bool, error) {
	if len(s.allEntities) == 0 {
		return false, ErrorEmptyList
	}

	for key, entity := range s.allEntities {
		if entity.Id == officeId {
			s.allEntities = append(s.allEntities[:key], s.allEntities[key+1:]...)
			return true, nil
		}
	}

	return false, ErrorNotFound
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

	for k, entity := range s.allEntities {
		if entity.Id == officeId {
			s.allEntities[k].Name = office.Name
			s.allEntities[k].Description = office.Description
			return nil
		}
	}

	return ErrorNotFound
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
