package config

import (
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/discovery"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/driver/cluster/balancer"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/meta/credentials"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
	"time"
)

// Config contains driver configuration options.
type Config struct {
	// Database is a required database name.
	Database string

	// Credentials is an ydb client credentials.
	// In most cases Credentials are required.
	Credentials credentials.Credentials

	// Trace contains driver tracing options.
	Trace trace.DriverTrace

	// RequestTimeout is the maximum amount of time a Call() will wait for an
	// operation to complete.
	// If RequestTimeout is zero then no timeout is used.
	RequestTimeout time.Duration

	// StreamTimeout is the maximum amount of time a StreamRead() will wait for
	// an operation to complete.
	// If StreamTimeout is zero then no timeout is used.
	StreamTimeout time.Duration

	// OperationTimeout is the maximum amount of time a YDB server will process
	// an operation. After timeout exceeds YDB will try to cancel operation and
	// regardless of the cancellation appropriate error will be returned to
	// the client.
	// If OperationTimeout is zero then no timeout is used.
	OperationTimeout time.Duration

	// OperationCancelAfter is the maximum amount of time a YDB server will process an
	// operation. After timeout exceeds YDB will try to cancel operation and if
	// it succeeds appropriate error will be returned to the client; otherwise
	// processing will be continued.
	// If OperationCancelAfter is zero then no timeout is used.
	OperationCancelAfter time.Duration

	// DiscoveryInterval is the frequency of background tasks of ydb endpoints
	// discovery.
	// If DiscoveryInterval is zero then the DefaultDiscoveryInterval is used.
	// If DiscoveryInterval is negative, then no background discovery prepared.
	DiscoveryInterval time.Duration

	// ConnectionTTL is a time to live duration for close idle grpc connection
	// If ConnectionTTL less or equal zero grpc connections will always stay
	// online. If ConnectionTTL is positive^ then grpc connections will be used
	// with lazy dialing and idle connections will be automatically closed after
	// expiration ConnectionTTL
	ConnectionTTL time.Duration

	// BalancingConfig is an optional configuration related to selected
	// BalancingMethod. That is, some balancing methods allow to be configured.
	BalancingConfig balancer.Config

	// PreferLocalEndpoints adds endpoint selection logic when local endpoints
	// are always used first.
	// When no alive local endpoints left other endpoints will be used.
	//
	// NOTE: some balancing methods (such as p2c) also may use knowledge of
	// endpoint's locality. Difference is that with PreferLocalEndpoints local
	// endpoints selected separately from others. That is, if there at least
	// one local endpoint it will be used regardless of its performance
	// indicators.
	//
	// NOTE: currently driver (and even ydb itself) does not track load factor
	// of each endpoint properly. Enabling this option may lead to the
	// situation, when all but one nodes in local datacenter become inactive
	// and all clients will overload this single instance very quickly. That
	// is, currently this option may be called as experimental.
	// You have been warned.
	PreferLocalEndpoints bool

	// RequestsType set an additional types hint to all requests.
	// It is needed only for debug purposes and advanced cases.
	RequestsType string

	// FastDial will make dialer return Driver as soon as 1st connection succeeds.
	// NB: it may be not the fastest node to serve requests.
	FastDial bool
}

func (d *Config) WithDefaults() (c Config) {
	if d != nil {
		c = *d
	}
	if c.DiscoveryInterval == 0 {
		c.DiscoveryInterval = discovery.DefaultDiscoveryInterval
	}
	return c
}