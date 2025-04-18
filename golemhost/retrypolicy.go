package golemhost

import (
	"time"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
)

type RetryPolicy struct {
	MaxAttempts uint32
	MinDelay    time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
}

func NewRetryPolicy(policy host.RetryPolicy) RetryPolicy {
	return RetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    time.Duration(policy.MinDelay) * time.Nanosecond,
		MaxDelay:    time.Duration(policy.MaxDelay) * time.Nanosecond,
		Multiplier:  policy.Multiplier,
	}
}

func (policy RetryPolicy) ToBinding() host.RetryPolicy {
	return host.RetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    host.Duration(policy.MinDelay.Nanoseconds()),
		MaxDelay:    host.Duration(policy.MaxDelay.Nanoseconds()),
		Multiplier:  policy.Multiplier,
	}
}

func GetRetryPolicy() RetryPolicy {
	return NewRetryPolicy(host.GetRetryPolicy())
}

func SetRetryPolicy(policy RetryPolicy) {
	host.SetRetryPolicy(policy.ToBinding())
}

func WithRetryPolicy[T any](policy RetryPolicy, f func() (T, error)) (T, error) {
	currentPolicy := GetRetryPolicy()
	defer SetRetryPolicy(currentPolicy)
	SetRetryPolicy(policy)
	return f()
}
