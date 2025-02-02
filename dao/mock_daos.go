package dao

import (
	"errors"
	"fmt"

	m "github.com/RedHatInsights/sources-api-go/model"
	"github.com/RedHatInsights/sources-api-go/util"
)

type MockSourceDao struct {
	Sources []m.Source
}

type MockApplicationDao struct {
	Applications []m.Application
}

type MockEndpointDao struct {
	Endpoints []m.Endpoint
}
type MockApplicationTypeDao struct {
	ApplicationTypes []m.ApplicationType
	Compatible       bool
}

type MockSourceTypeDao struct {
	SourceTypes []m.SourceType
}

type MockMetaDataDao struct {
	MetaDatas []m.MetaData
}

func (src *MockSourceDao) SubCollectionList(primaryCollection interface{}, limit, offset int, filters []util.Filter) ([]m.Source, int64, error) {
	var sources []m.Source

	for index, i := range src.Sources {
		switch object := primaryCollection.(type) {
		case m.SourceType:
			if i.SourceTypeID == object.Id {
				sources = append(sources, src.Sources[index])
			}
		case m.ApplicationType:
			if i.ID == object.Id {
				sources = append(sources, src.Sources[index])
			}
		default:
			return nil, 0, fmt.Errorf("unexpected primary collection type")
		}
	}

	count := int64(len(sources))
	if count == 0 {
		return nil, count, util.NewErrNotFound("source")
	}

	return sources, count, nil
}

func (src *MockSourceDao) List(limit, offset int, filters []util.Filter) ([]m.Source, int64, error) {
	count := int64(len(src.Sources))
	return src.Sources, count, nil
}

func (src *MockSourceDao) ListInternal(limit, offset int, filters []util.Filter) ([]m.Source, int64, error) {
	count := int64(len(src.Sources))
	return src.Sources, count, nil
}

func (src *MockSourceDao) GetById(id *int64) (*m.Source, error) {
	for _, i := range src.Sources {
		if i.ID == *id {
			return &i, nil
		}
	}

	return nil, util.NewErrNotFound("source")
}

func (src *MockSourceDao) Create(s *m.Source) error {
	return nil
}

func (src *MockSourceDao) Update(s *m.Source) error {
	panic("implement me")
}

func (src *MockSourceDao) Delete(id *int64) (*m.Source, error) {
	panic("implement me")
}

func (src *MockSourceDao) Tenant() *int64 {
	tenant := int64(1)
	return &tenant
}

// NameExistsInCurrentTenant returns always false because it's the safe default in case the request gets validated
// in the tests.
func (src *MockSourceDao) NameExistsInCurrentTenant(name string) bool {
	return false
}

func (src *MockSourceDao) GetByIdWithPreload(id *int64, preloads ...string) (*m.Source, error) {
	for _, i := range src.Sources {
		if i.ID == *id {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("source not found")
}

func (a *MockApplicationTypeDao) List(limit int, offset int, filters []util.Filter) ([]m.ApplicationType, int64, error) {
	count := int64(len(a.ApplicationTypes))
	return a.ApplicationTypes, count, nil
}

func (a *MockApplicationTypeDao) GetById(id *int64) (*m.ApplicationType, error) {
	for _, i := range a.ApplicationTypes {
		if i.Id == *id {
			return &i, nil
		}
	}

	return nil, util.NewErrNotFound("application type")
}

func (a *MockMetaDataDao) List(limit int, offset int, filters []util.Filter) ([]m.MetaData, int64, error) {
	count := int64(len(a.MetaDatas))
	return a.MetaDatas, count, nil
}

func (a *MockMetaDataDao) SubCollectionList(primaryCollection interface{}, limit, offset int, filters []util.Filter) ([]m.MetaData, int64, error) {
	var appMetaDataList []m.MetaData

	for index, i := range a.MetaDatas {
		switch object := primaryCollection.(type) {
		case m.ApplicationType:
			if i.ApplicationTypeID == object.Id {
				appMetaDataList = append(appMetaDataList, a.MetaDatas[index])
			}
		default:
			return nil, 0, fmt.Errorf("unexpected primary collection type")
		}
	}
	count := int64(len(appMetaDataList))
	if count == 0 {
		return nil, count, util.NewErrNotFound("metadata")
	}

	return appMetaDataList, count, nil
}

func (a *MockMetaDataDao) GetById(id *int64) (*m.MetaData, error) {
	for _, i := range a.MetaDatas {
		if i.ID == *id {
			return &i, nil
		}
	}

	return nil, util.NewErrNotFound("metadata")
}

func (a *MockMetaDataDao) Create(src *m.MetaData) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockMetaDataDao) Update(src *m.MetaData) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockMetaDataDao) Delete(id *int64) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockApplicationTypeDao) Create(src *m.ApplicationType) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockApplicationTypeDao) Update(src *m.ApplicationType) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockApplicationTypeDao) Delete(id *int64) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockApplicationTypeDao) SubCollectionList(primaryCollection interface{}, limit, offset int, filters []util.Filter) ([]m.ApplicationType, int64, error) {
	var appTypes []m.ApplicationType

	for index, i := range a.ApplicationTypes {
		switch object := primaryCollection.(type) {
		case m.Source:
			if i.Id == object.ID {
				appTypes = append(appTypes, a.ApplicationTypes[index])
			}
		default:
			return nil, 0, fmt.Errorf("unexpected primary collection type")
		}
	}
	count := int64(len(appTypes))
	if count == 0 {
		return nil, count, util.NewErrNotFound("application type")
	}

	return appTypes, count, nil
}

func (a *MockApplicationTypeDao) ApplicationTypeCompatibleWithSource(_, _ int64) error {
	if a.Compatible {
		return nil
	}

	return errors.New("Not compatible!")
}

func (a *MockSourceTypeDao) List(limit int, offset int, filters []util.Filter) ([]m.SourceType, int64, error) {
	count := int64(len(a.SourceTypes))
	return a.SourceTypes, count, nil
}

func (a *MockSourceTypeDao) GetById(id *int64) (*m.SourceType, error) {
	for _, i := range a.SourceTypes {
		if i.Id == *id {
			return &i, nil
		}
	}

	return nil, util.NewErrNotFound("source type")
}

func (a *MockSourceTypeDao) Create(src *m.SourceType) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockSourceTypeDao) Update(src *m.SourceType) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockSourceTypeDao) Delete(id *int64) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockApplicationDao) SubCollectionList(primaryCollection interface{}, limit, offset int, filters []util.Filter) ([]m.Application, int64, error) {
	var applications []m.Application

	for index, i := range a.Applications {
		switch object := primaryCollection.(type) {
		case m.Source:
			if i.SourceID == object.ID {
				applications = append(applications, a.Applications[index])
			}
		default:
			return nil, 0, fmt.Errorf("unexpected primary collection type")
		}
	}
	count := int64(len(applications))
	if count == 0 {
		return nil, count, util.NewErrNotFound("application")
	}

	return applications, count, nil
}

func (a *MockApplicationDao) List(limit int, offset int, filters []util.Filter) ([]m.Application, int64, error) {
	count := int64(len(a.Applications))
	return a.Applications, count, nil
}

func (a *MockApplicationDao) GetById(id *int64) (*m.Application, error) {
	for _, app := range a.Applications {
		if app.ID == *id {
			return &app, nil
		}
	}

	return nil, util.NewErrNotFound("application")
}

func (a *MockApplicationDao) Create(src *m.Application) error {
	return nil
}

func (a *MockApplicationDao) Update(src *m.Application) error {
	return nil
}

func (a *MockApplicationDao) Delete(id *int64) (*m.Application, error) {
	return nil, nil
}

func (a *MockApplicationDao) Tenant() *int64 {
	tenant := int64(1)
	return &tenant
}

func (a *MockEndpointDao) SubCollectionList(primaryCollection interface{}, limit, offset int, filters []util.Filter) ([]m.Endpoint, int64, error) {
	var endpoints []m.Endpoint

	for index, i := range a.Endpoints {
		switch object := primaryCollection.(type) {
		case m.Source:
			if i.SourceID == object.ID {
				endpoints = append(endpoints, a.Endpoints[index])
			}
		default:
			return nil, 0, fmt.Errorf("unexpected primary collection type")
		}
	}
	count := int64(len(endpoints))
	if count == 0 {
		return nil, count, util.NewErrNotFound("endpoint")
	}

	return endpoints, count, nil
}

func (a *MockEndpointDao) List(limit int, offset int, filters []util.Filter) ([]m.Endpoint, int64, error) {
	count := int64(len(a.Endpoints))
	return a.Endpoints, count, nil
}

func (a *MockEndpointDao) GetById(id *int64) (*m.Endpoint, error) {
	for _, app := range a.Endpoints {
		if app.ID == *id {
			return &app, nil
		}
	}

	return nil, util.NewErrNotFound("endpoint")
}

func (a *MockEndpointDao) Create(src *m.Endpoint) error {
	return nil
}

func (a *MockEndpointDao) Update(src *m.Endpoint) error {
	panic("not implemented") // TODO: Implement
}

func (a *MockEndpointDao) Delete(id *int64) (*m.Endpoint, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockEndpointDao) Tenant() *int64 {
	tenant := int64(1)
	return &tenant
}

func (m *MockEndpointDao) CanEndpointBeSetAsDefaultForSource(sourceId int64) bool {
	return true
}

func (m *MockEndpointDao) IsRoleUniqueForSource(role string, sourceId int64) bool {
	return true
}

func (m *MockEndpointDao) SourceHasEndpoints(sourceId int64) bool {
	return true
}
