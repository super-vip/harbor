// Code generated by mockery v2.53.3. DO NOT EDIT.

package task

import (
	context "context"

	dao "github.com/goharbor/harbor/src/pkg/task/dao"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// mockExecutionDAO is an autogenerated mock type for the ExecutionDAO type
type mockExecutionDAO struct {
	mock.Mock
}

// AsyncRefreshStatus provides a mock function with given fields: ctx, id, vendor
func (_m *mockExecutionDAO) AsyncRefreshStatus(ctx context.Context, id int64, vendor string) error {
	ret := _m.Called(ctx, id, vendor)

	if len(ret) == 0 {
		panic("no return value specified for AsyncRefreshStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, id, vendor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx, query
func (_m *mockExecutionDAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, execution
func (_m *mockExecutionDAO) Create(ctx context.Context, execution *dao.Execution) (int64, error) {
	ret := _m.Called(ctx, execution)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Execution) (int64, error)); ok {
		return rf(ctx, execution)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Execution) int64); ok {
		r0 = rf(ctx, execution)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dao.Execution) error); ok {
		r1 = rf(ctx, execution)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) Get(ctx context.Context, id int64) (*dao.Execution, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *dao.Execution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*dao.Execution, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dao.Execution); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.Execution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetrics provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) GetMetrics(ctx context.Context, id int64) (*dao.Metrics, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMetrics")
	}

	var r0 *dao.Metrics
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*dao.Metrics, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dao.Metrics); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.Metrics)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *mockExecutionDAO) List(ctx context.Context, query *q.Query) ([]*dao.Execution, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*dao.Execution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*dao.Execution, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*dao.Execution); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.Execution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshStatus provides a mock function with given fields: ctx, id
func (_m *mockExecutionDAO) RefreshStatus(ctx context.Context, id int64) (bool, string, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RefreshStatus")
	}

	var r0 bool
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (bool, string, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) string); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, execution, props
func (_m *mockExecutionDAO) Update(ctx context.Context, execution *dao.Execution, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, execution)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dao.Execution, ...string) error); ok {
		r0 = rf(ctx, execution, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockExecutionDAO creates a new instance of mockExecutionDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockExecutionDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockExecutionDAO {
	mock := &mockExecutionDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
