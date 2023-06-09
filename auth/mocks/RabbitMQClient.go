// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	amqp091 "github.com/rabbitmq/amqp091-go"

	mock "github.com/stretchr/testify/mock"
)

// RabbitMQClient is an autogenerated mock type for the RabbitMQClient type
type RabbitMQClient struct {
	mock.Mock
}

// ChClose provides a mock function with given fields:
func (_m *RabbitMQClient) ChClose() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnClose provides a mock function with given fields:
func (_m *RabbitMQClient) ConnClose() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: ctx, exchangeName, exchangeType, contentType, routingKeyPub, dataBytes
func (_m *RabbitMQClient) Publish(ctx context.Context, exchangeName string, exchangeType string, contentType string, routingKeyPub string, dataBytes []byte) error {
	ret := _m.Called(ctx, exchangeName, exchangeType, contentType, routingKeyPub, dataBytes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []byte) error); ok {
		r0 = rf(ctx, exchangeName, exchangeType, contentType, routingKeyPub, dataBytes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *RabbitMQClient) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: exchangeName, exchangeType, queueName, routingKeyCon, autoAck, prefetchCount
func (_m *RabbitMQClient) Subscribe(exchangeName string, exchangeType string, queueName string, routingKeyCon string, autoAck bool, prefetchCount int) (<-chan amqp091.Delivery, error) {
	ret := _m.Called(exchangeName, exchangeType, queueName, routingKeyCon, autoAck, prefetchCount)

	var r0 <-chan amqp091.Delivery
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool, int) (<-chan amqp091.Delivery, error)); ok {
		return rf(exchangeName, exchangeType, queueName, routingKeyCon, autoAck, prefetchCount)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool, int) <-chan amqp091.Delivery); ok {
		r0 = rf(exchangeName, exchangeType, queueName, routingKeyCon, autoAck, prefetchCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp091.Delivery)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, bool, int) error); ok {
		r1 = rf(exchangeName, exchangeType, queueName, routingKeyCon, autoAck, prefetchCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRabbitMQClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewRabbitMQClient creates a new instance of RabbitMQClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRabbitMQClient(t mockConstructorTestingTNewRabbitMQClient) *RabbitMQClient {
	mock := &RabbitMQClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
