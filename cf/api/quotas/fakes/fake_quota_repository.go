// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/api/quotas"

	"github.com/cloudfoundry/cli/cf/models"

	"sync"
)

type FakeQuotaRepository struct {
	FindAllStub        func() (quotas []models.QuotaFields, apiErr error)
	findAllMutex       sync.RWMutex
	findAllArgsForCall []struct{}
	findAllReturns     struct {
		result1 []models.QuotaFields
		result2 error
	}
	FindByNameStub        func(name string) (quota models.QuotaFields, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.QuotaFields
		result2 error
	}
	AssignQuotaToOrgStub        func(orgGuid, quotaGuid string) error
	assignQuotaToOrgMutex       sync.RWMutex
	assignQuotaToOrgArgsForCall []struct {
		orgGuid   string
		quotaGuid string
	}
	assignQuotaToOrgReturns struct {
		result1 error
	}
	CreateStub        func(quota models.QuotaFields) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		quota models.QuotaFields
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(quota models.QuotaFields) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		quota models.QuotaFields
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(quotaGuid string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		quotaGuid string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeQuotaRepository) FindAll() (quotas []models.QuotaFields, apiErr error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.findAllArgsForCall = append(fake.findAllArgsForCall, struct{}{})
	if fake.FindAllStub != nil {
		return fake.FindAllStub()
	} else {
		return fake.findAllReturns.result1, fake.findAllReturns.result2
	}
}

func (fake *FakeQuotaRepository) FindAllCallCount() int {
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	return len(fake.findAllArgsForCall)
}

func (fake *FakeQuotaRepository) FindAllReturns(result1 []models.QuotaFields, result2 error) {
	fake.findAllReturns = struct {
		result1 []models.QuotaFields
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaRepository) FindByName(name string) (quota models.QuotaFields, apiErr error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeQuotaRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeQuotaRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeQuotaRepository) FindByNameReturns(result1 models.QuotaFields, result2 error) {
	fake.findByNameReturns = struct {
		result1 models.QuotaFields
		result2 error
	}{result1, result2}
}

func (fake *FakeQuotaRepository) AssignQuotaToOrg(orgGuid string, quotaGuid string) error {
	fake.assignQuotaToOrgMutex.Lock()
	defer fake.assignQuotaToOrgMutex.Unlock()
	fake.assignQuotaToOrgArgsForCall = append(fake.assignQuotaToOrgArgsForCall, struct {
		orgGuid   string
		quotaGuid string
	}{orgGuid, quotaGuid})
	if fake.AssignQuotaToOrgStub != nil {
		return fake.AssignQuotaToOrgStub(orgGuid, quotaGuid)
	} else {
		return fake.assignQuotaToOrgReturns.result1
	}
}

func (fake *FakeQuotaRepository) AssignQuotaToOrgCallCount() int {
	fake.assignQuotaToOrgMutex.RLock()
	defer fake.assignQuotaToOrgMutex.RUnlock()
	return len(fake.assignQuotaToOrgArgsForCall)
}

func (fake *FakeQuotaRepository) AssignQuotaToOrgArgsForCall(i int) (string, string) {
	fake.assignQuotaToOrgMutex.RLock()
	defer fake.assignQuotaToOrgMutex.RUnlock()
	return fake.assignQuotaToOrgArgsForCall[i].orgGuid, fake.assignQuotaToOrgArgsForCall[i].quotaGuid
}

func (fake *FakeQuotaRepository) AssignQuotaToOrgReturns(result1 error) {
	fake.assignQuotaToOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQuotaRepository) Create(quota models.QuotaFields) error {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		quota models.QuotaFields
	}{quota})
	if fake.CreateStub != nil {
		return fake.CreateStub(quota)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeQuotaRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeQuotaRepository) CreateArgsForCall(i int) models.QuotaFields {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].quota
}

func (fake *FakeQuotaRepository) CreateReturns(result1 error) {
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQuotaRepository) Update(quota models.QuotaFields) error {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		quota models.QuotaFields
	}{quota})
	if fake.UpdateStub != nil {
		return fake.UpdateStub(quota)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeQuotaRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeQuotaRepository) UpdateArgsForCall(i int) models.QuotaFields {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].quota
}

func (fake *FakeQuotaRepository) UpdateReturns(result1 error) {
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQuotaRepository) Delete(quotaGuid string) error {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		quotaGuid string
	}{quotaGuid})
	if fake.DeleteStub != nil {
		return fake.DeleteStub(quotaGuid)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeQuotaRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeQuotaRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].quotaGuid
}

func (fake *FakeQuotaRepository) DeleteReturns(result1 error) {
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ QuotaRepository = new(FakeQuotaRepository)
