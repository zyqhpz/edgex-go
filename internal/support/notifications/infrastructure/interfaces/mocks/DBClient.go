// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddNotification provides a mock function with given fields: n
func (_m *DBClient) AddNotification(n models.Notification) (models.Notification, errors.EdgeX) {
	ret := _m.Called(n)

	var r0 models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Notification) (models.Notification, errors.EdgeX)); ok {
		return rf(n)
	}
	if rf, ok := ret.Get(0).(func(models.Notification) models.Notification); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.Notification)
	}

	if rf, ok := ret.Get(1).(func(models.Notification) errors.EdgeX); ok {
		r1 = rf(n)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddSubscription provides a mock function with given fields: e
func (_m *DBClient) AddSubscription(e models.Subscription) (models.Subscription, errors.EdgeX) {
	ret := _m.Called(e)

	var r0 models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Subscription) (models.Subscription, errors.EdgeX)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(models.Subscription) models.Subscription); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(models.Subscription)
	}

	if rf, ok := ret.Get(1).(func(models.Subscription) errors.EdgeX); ok {
		r1 = rf(e)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddTransmission provides a mock function with given fields: trans
func (_m *DBClient) AddTransmission(trans models.Transmission) (models.Transmission, errors.EdgeX) {
	ret := _m.Called(trans)

	var r0 models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Transmission) (models.Transmission, errors.EdgeX)); ok {
		return rf(trans)
	}
	if rf, ok := ret.Get(0).(func(models.Transmission) models.Transmission); ok {
		r0 = rf(trans)
	} else {
		r0 = ret.Get(0).(models.Transmission)
	}

	if rf, ok := ret.Get(1).(func(models.Transmission) errors.EdgeX); ok {
		r1 = rf(trans)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllSubscriptions provides a mock function with given fields: offset, limit
func (_m *DBClient) AllSubscriptions(offset int, limit int) ([]models.Subscription, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	var r0 []models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) ([]models.Subscription, errors.EdgeX)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []models.Subscription); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllTransmissions provides a mock function with given fields: offset, limit
func (_m *DBClient) AllTransmissions(offset int, limit int) ([]models.Transmission, errors.EdgeX) {
	ret := _m.Called(offset, limit)

	var r0 []models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) ([]models.Transmission, errors.EdgeX)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []models.Transmission); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CleanupNotificationsByAge provides a mock function with given fields: age
func (_m *DBClient) CleanupNotificationsByAge(age int64) errors.EdgeX {
	ret := _m.Called(age)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64) errors.EdgeX); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteNotificationById provides a mock function with given fields: id
func (_m *DBClient) DeleteNotificationById(id string) errors.EdgeX {
	ret := _m.Called(id)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteProcessedNotificationsByAge provides a mock function with given fields: age
func (_m *DBClient) DeleteProcessedNotificationsByAge(age int64) errors.EdgeX {
	ret := _m.Called(age)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64) errors.EdgeX); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteProcessedTransmissionsByAge provides a mock function with given fields: age
func (_m *DBClient) DeleteProcessedTransmissionsByAge(age int64) errors.EdgeX {
	ret := _m.Called(age)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int64) errors.EdgeX); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteSubscriptionByName provides a mock function with given fields: name
func (_m *DBClient) DeleteSubscriptionByName(name string) errors.EdgeX {
	ret := _m.Called(name)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) errors.EdgeX); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// LatestNotificationByOffset provides a mock function with given fields: offset
func (_m *DBClient) LatestNotificationByOffset(offset uint32) (models.Notification, errors.EdgeX) {
	ret := _m.Called(offset)

	var r0 models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(uint32) (models.Notification, errors.EdgeX)); ok {
		return rf(offset)
	}
	if rf, ok := ret.Get(0).(func(uint32) models.Notification); ok {
		r0 = rf(offset)
	} else {
		r0 = ret.Get(0).(models.Notification)
	}

	if rf, ok := ret.Get(1).(func(uint32) errors.EdgeX); ok {
		r1 = rf(offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationById provides a mock function with given fields: id
func (_m *DBClient) NotificationById(id string) (models.Notification, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (models.Notification, errors.EdgeX)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Notification); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Notification)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationCountByCategoriesAndLabels provides a mock function with given fields: categories, labels
func (_m *DBClient) NotificationCountByCategoriesAndLabels(categories []string, labels []string) (uint32, errors.EdgeX) {
	ret := _m.Called(categories, labels)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func([]string, []string) (uint32, errors.EdgeX)); ok {
		return rf(categories, labels)
	}
	if rf, ok := ret.Get(0).(func([]string, []string) uint32); ok {
		r0 = rf(categories, labels)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func([]string, []string) errors.EdgeX); ok {
		r1 = rf(categories, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationCountByCategory provides a mock function with given fields: category
func (_m *DBClient) NotificationCountByCategory(category string) (uint32, errors.EdgeX) {
	ret := _m.Called(category)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(category)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationCountByLabel provides a mock function with given fields: label
func (_m *DBClient) NotificationCountByLabel(label string) (uint32, errors.EdgeX) {
	ret := _m.Called(label)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(label)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(label)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(label)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationCountByStatus provides a mock function with given fields: status
func (_m *DBClient) NotificationCountByStatus(status string) (uint32, errors.EdgeX) {
	ret := _m.Called(status)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationCountByTimeRange provides a mock function with given fields: start, end
func (_m *DBClient) NotificationCountByTimeRange(start int, end int) (uint32, errors.EdgeX) {
	ret := _m.Called(start, end)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) (uint32, errors.EdgeX)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(int, int) uint32); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationTotalCount provides a mock function with given fields:
func (_m *DBClient) NotificationTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func() (uint32, errors.EdgeX)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByCategoriesAndLabels provides a mock function with given fields: offset, limit, categories, labels
func (_m *DBClient) NotificationsByCategoriesAndLabels(offset int, limit int, categories []string, labels []string) ([]models.Notification, errors.EdgeX) {
	ret := _m.Called(offset, limit, categories, labels)

	var r0 []models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, []string, []string) ([]models.Notification, errors.EdgeX)); ok {
		return rf(offset, limit, categories, labels)
	}
	if rf, ok := ret.Get(0).(func(int, int, []string, []string) []models.Notification); ok {
		r0 = rf(offset, limit, categories, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, []string, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, categories, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByCategory provides a mock function with given fields: offset, limit, category
func (_m *DBClient) NotificationsByCategory(offset int, limit int, category string) ([]models.Notification, errors.EdgeX) {
	ret := _m.Called(offset, limit, category)

	var r0 []models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Notification, errors.EdgeX)); ok {
		return rf(offset, limit, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Notification); ok {
		r0 = rf(offset, limit, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByLabel provides a mock function with given fields: offset, limit, label
func (_m *DBClient) NotificationsByLabel(offset int, limit int, label string) ([]models.Notification, errors.EdgeX) {
	ret := _m.Called(offset, limit, label)

	var r0 []models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Notification, errors.EdgeX)); ok {
		return rf(offset, limit, label)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Notification); ok {
		r0 = rf(offset, limit, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, label)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByStatus provides a mock function with given fields: offset, limit, status
func (_m *DBClient) NotificationsByStatus(offset int, limit int, status string) ([]models.Notification, errors.EdgeX) {
	ret := _m.Called(offset, limit, status)

	var r0 []models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Notification, errors.EdgeX)); ok {
		return rf(offset, limit, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Notification); ok {
		r0 = rf(offset, limit, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NotificationsByTimeRange provides a mock function with given fields: start, end, offset, limit
func (_m *DBClient) NotificationsByTimeRange(start int, end int, offset int, limit int) ([]models.Notification, errors.EdgeX) {
	ret := _m.Called(start, end, offset, limit)

	var r0 []models.Notification
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, int, int) ([]models.Notification, errors.EdgeX)); ok {
		return rf(start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int, int, int) []models.Notification); ok {
		r0 = rf(start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, int, int) errors.EdgeX); ok {
		r1 = rf(start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionById provides a mock function with given fields: id
func (_m *DBClient) SubscriptionById(id string) (models.Subscription, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (models.Subscription, errors.EdgeX)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Subscription); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Subscription)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionByName provides a mock function with given fields: name
func (_m *DBClient) SubscriptionByName(name string) (models.Subscription, errors.EdgeX) {
	ret := _m.Called(name)

	var r0 models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (models.Subscription, errors.EdgeX)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.Subscription); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Subscription)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionCountByCategory provides a mock function with given fields: category
func (_m *DBClient) SubscriptionCountByCategory(category string) (uint32, errors.EdgeX) {
	ret := _m.Called(category)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(category)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionCountByLabel provides a mock function with given fields: label
func (_m *DBClient) SubscriptionCountByLabel(label string) (uint32, errors.EdgeX) {
	ret := _m.Called(label)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(label)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(label)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(label)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionCountByReceiver provides a mock function with given fields: receiver
func (_m *DBClient) SubscriptionCountByReceiver(receiver string) (uint32, errors.EdgeX) {
	ret := _m.Called(receiver)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(receiver)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(receiver)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(receiver)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionTotalCount provides a mock function with given fields:
func (_m *DBClient) SubscriptionTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func() (uint32, errors.EdgeX)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByCategoriesAndLabels provides a mock function with given fields: offset, limit, categories, labels
func (_m *DBClient) SubscriptionsByCategoriesAndLabels(offset int, limit int, categories []string, labels []string) ([]models.Subscription, errors.EdgeX) {
	ret := _m.Called(offset, limit, categories, labels)

	var r0 []models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, []string, []string) ([]models.Subscription, errors.EdgeX)); ok {
		return rf(offset, limit, categories, labels)
	}
	if rf, ok := ret.Get(0).(func(int, int, []string, []string) []models.Subscription); ok {
		r0 = rf(offset, limit, categories, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, []string, []string) errors.EdgeX); ok {
		r1 = rf(offset, limit, categories, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByCategory provides a mock function with given fields: offset, limit, category
func (_m *DBClient) SubscriptionsByCategory(offset int, limit int, category string) ([]models.Subscription, errors.EdgeX) {
	ret := _m.Called(offset, limit, category)

	var r0 []models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Subscription, errors.EdgeX)); ok {
		return rf(offset, limit, category)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Subscription); ok {
		r0 = rf(offset, limit, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, category)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByLabel provides a mock function with given fields: offset, limit, label
func (_m *DBClient) SubscriptionsByLabel(offset int, limit int, label string) ([]models.Subscription, errors.EdgeX) {
	ret := _m.Called(offset, limit, label)

	var r0 []models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Subscription, errors.EdgeX)); ok {
		return rf(offset, limit, label)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Subscription); ok {
		r0 = rf(offset, limit, label)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, label)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SubscriptionsByReceiver provides a mock function with given fields: offset, limit, receiver
func (_m *DBClient) SubscriptionsByReceiver(offset int, limit int, receiver string) ([]models.Subscription, errors.EdgeX) {
	ret := _m.Called(offset, limit, receiver)

	var r0 []models.Subscription
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Subscription, errors.EdgeX)); ok {
		return rf(offset, limit, receiver)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Subscription); ok {
		r0 = rf(offset, limit, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, receiver)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionById provides a mock function with given fields: id
func (_m *DBClient) TransmissionById(id string) (models.Transmission, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (models.Transmission, errors.EdgeX)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Transmission); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Transmission)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionCountByNotificationId provides a mock function with given fields: id
func (_m *DBClient) TransmissionCountByNotificationId(id string) (uint32, errors.EdgeX) {
	ret := _m.Called(id)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionCountByStatus provides a mock function with given fields: status
func (_m *DBClient) TransmissionCountByStatus(status string) (uint32, errors.EdgeX) {
	ret := _m.Called(status)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionCountBySubscriptionName provides a mock function with given fields: subscriptionName
func (_m *DBClient) TransmissionCountBySubscriptionName(subscriptionName string) (uint32, errors.EdgeX) {
	ret := _m.Called(subscriptionName)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(string) (uint32, errors.EdgeX)); ok {
		return rf(subscriptionName)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(subscriptionName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) errors.EdgeX); ok {
		r1 = rf(subscriptionName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionCountByTimeRange provides a mock function with given fields: start, end
func (_m *DBClient) TransmissionCountByTimeRange(start int, end int) (uint32, errors.EdgeX) {
	ret := _m.Called(start, end)

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int) (uint32, errors.EdgeX)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(int, int) uint32); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(int, int) errors.EdgeX); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionTotalCount provides a mock function with given fields:
func (_m *DBClient) TransmissionTotalCount() (uint32, errors.EdgeX) {
	ret := _m.Called()

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func() (uint32, errors.EdgeX)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() errors.EdgeX); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionsByNotificationId provides a mock function with given fields: offset, limit, id
func (_m *DBClient) TransmissionsByNotificationId(offset int, limit int, id string) ([]models.Transmission, errors.EdgeX) {
	ret := _m.Called(offset, limit, id)

	var r0 []models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Transmission, errors.EdgeX)); ok {
		return rf(offset, limit, id)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Transmission); ok {
		r0 = rf(offset, limit, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionsByStatus provides a mock function with given fields: offset, limit, status
func (_m *DBClient) TransmissionsByStatus(offset int, limit int, status string) ([]models.Transmission, errors.EdgeX) {
	ret := _m.Called(offset, limit, status)

	var r0 []models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Transmission, errors.EdgeX)); ok {
		return rf(offset, limit, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Transmission); ok {
		r0 = rf(offset, limit, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionsBySubscriptionName provides a mock function with given fields: offset, limit, subscriptionName
func (_m *DBClient) TransmissionsBySubscriptionName(offset int, limit int, subscriptionName string) ([]models.Transmission, errors.EdgeX) {
	ret := _m.Called(offset, limit, subscriptionName)

	var r0 []models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, string) ([]models.Transmission, errors.EdgeX)); ok {
		return rf(offset, limit, subscriptionName)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []models.Transmission); ok {
		r0 = rf(offset, limit, subscriptionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) errors.EdgeX); ok {
		r1 = rf(offset, limit, subscriptionName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// TransmissionsByTimeRange provides a mock function with given fields: start, end, offset, limit
func (_m *DBClient) TransmissionsByTimeRange(start int, end int, offset int, limit int) ([]models.Transmission, errors.EdgeX) {
	ret := _m.Called(start, end, offset, limit)

	var r0 []models.Transmission
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(int, int, int, int) ([]models.Transmission, errors.EdgeX)); ok {
		return rf(start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int, int, int) []models.Transmission); ok {
		r0 = rf(start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transmission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, int, int) errors.EdgeX); ok {
		r1 = rf(start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateNotification provides a mock function with given fields: s
func (_m *DBClient) UpdateNotification(s models.Notification) errors.EdgeX {
	ret := _m.Called(s)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Notification) errors.EdgeX); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// UpdateSubscription provides a mock function with given fields: s
func (_m *DBClient) UpdateSubscription(s models.Subscription) errors.EdgeX {
	ret := _m.Called(s)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Subscription) errors.EdgeX); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// UpdateTransmission provides a mock function with given fields: trans
func (_m *DBClient) UpdateTransmission(trans models.Transmission) errors.EdgeX {
	ret := _m.Called(trans)

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(models.Transmission) errors.EdgeX); ok {
		r0 = rf(trans)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

type mockConstructorTestingTNewDBClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBClient creates a new instance of DBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBClient(t mockConstructorTestingTNewDBClient) *DBClient {
	mock := &DBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
