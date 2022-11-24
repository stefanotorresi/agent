// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/trento-project/agent/pkg/factsengine/entities"

	mock "github.com/stretchr/testify/mock"
)

// FactGatherer is an autogenerated mock type for the FactGatherer type
type FactGatherer struct {
	mock.Mock
}

// Gather provides a mock function with given fields: factsRequests
func (_m *FactGatherer) Gather(factsRequests []entities.FactRequest) ([]entities.Fact, error) {
	ret := _m.Called(factsRequests)

	var r0 []entities.Fact
	if rf, ok := ret.Get(0).(func([]entities.FactRequest) []entities.Fact); ok {
		r0 = rf(factsRequests)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Fact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]entities.FactRequest) error); ok {
		r1 = rf(factsRequests)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFactGatherer interface {
	mock.TestingT
	Cleanup(func())
}

// NewFactGatherer creates a new instance of FactGatherer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFactGatherer(t mockConstructorTestingTNewFactGatherer) *FactGatherer {
	mock := &FactGatherer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
