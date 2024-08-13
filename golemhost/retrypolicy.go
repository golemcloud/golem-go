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

func newRetryPolicy(policy binding.GolemApi0_2_0_HostRetryPolicy) RetryPolicy {
	return RetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    time.Duration(policy.MinDelay) * time.Nanosecond,
		MaxDelay:    time.Duration(policy.MaxDelay) * time.Nanosecond,
		Multiplier:  policy.Multiplier,
	}
}

func (policy RetryPolicy) toBinding() binding.GolemApi0_2_0_HostRetryPolicy {
	return binding.GolemApi0_2_0_HostRetryPolicy{
		MaxAttempts: policy.MaxAttempts,
		MinDelay:    binding.GolemApi0_2_0_HostDuration(policy.MinDelay.Nanoseconds()),
		MaxDelay:    binding.GolemApi0_2_0_HostDuration(policy.MaxDelay.Nanoseconds()),
		Multiplier:  policy.Multiplier,
	}
}

func GetRetryPolicy() RetryPolicy {
	return newRetryPolicy(binding.GolemApi0_2_0_HostGetRetryPolicy())
}

func SetRetryPolicy(policy RetryPolicy) {
	binding.GolemApi0_2_0_HostSetRetryPolicy(policy.toBinding())
}

func WithRetryPolicy[T any](policy RetryPolicy, f func() (T, error)) (T, error) {
	currentPolicy := GetRetryPolicy()
	defer SetRetryPolicy(currentPolicy)
	SetRetryPolicy(policy)
	return f()
}
