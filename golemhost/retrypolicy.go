package golemhost

import (
	"time"

	"github.com/golemcloud/golem-go/binding"
)

type RetryPolicy struct {
	MaxAttempts uint32
	MinDelay    time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
}

func NewRetryPolicy(policy binding.GolemApi1_1_6_HostRetryPolicy) RetryPolicy {
	return RetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    time.Duration(policy.MinDelay) * time.Nanosecond,
		MaxDelay:    time.Duration(policy.MaxDelay) * time.Nanosecond,
		Multiplier:  policy.Multiplier,
	}
}

func (policy RetryPolicy) ToBinding() binding.GolemApi1_1_6_HostRetryPolicy {
	return binding.GolemApi1_1_6_HostRetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    binding.GolemApi1_1_6_HostDuration(policy.MinDelay.Nanoseconds()),
		MaxDelay:    binding.GolemApi1_1_6_HostDuration(policy.MaxDelay.Nanoseconds()),
		Multiplier:  policy.Multiplier,
	}
}

func GetRetryPolicy() RetryPolicy {
	return NewRetryPolicy(binding.GolemApi1_1_6_HostGetRetryPolicy())
}

func SetRetryPolicy(policy RetryPolicy) {
	binding.GolemApi1_1_6_HostSetRetryPolicy(policy.ToBinding())
}

func WithRetryPolicy[T any](policy RetryPolicy, f func() (T, error)) (T, error) {
	currentPolicy := GetRetryPolicy()
	defer SetRetryPolicy(currentPolicy)
	SetRetryPolicy(policy)
	return f()
}
