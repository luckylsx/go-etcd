package registry

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type Registry struct {
	clientV3  *clientv3.Client
	clientKv  clientv3.KV
	lease     clientv3.Lease
	ctx       context.Context
	namespace string
	ttl       time.Duration
}

type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`
	// Name is the service name as registered.
	Name string `json:"name"`
	// Version is the version of the compiled.
	Version string `json:"version"`
	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`
	// Endpoints is endpoint addresses of the service instance.
	Endpoints []string `json:"endpoints"`
}

type Option func(opt *options)

const DefaultDuration = 10 * time.Second

type options struct {
	timeout     time.Duration
	username    string
	password    string
	dialTimeout time.Duration
	cxt         context.Context
	namespace   string
	ttl         time.Duration
}

func WithTimeout(timeout time.Duration) Option {
	return func(opt *options) {
		opt.timeout = timeout
	}
}

func WithUserName(username string) Option {
	return func(opt *options) {
		opt.username = username
	}
}

func WithPassword(password string) Option {
	return func(opt *options) {
		opt.password = password
	}
}

func WithDiaTimeout(diaTimeout time.Duration) Option {
	return func(opt *options) {
		opt.dialTimeout = diaTimeout
	}
}

func WithCtx(ctx context.Context) Option {
	return func(opt *options) {
		opt.cxt = ctx
	}
}

func WithTTl(ttl time.Duration) Option {
	return func(opt *options) {
		opt.ttl = ttl
	}
}

func WithNameSpace(namespace string) Option {
	return func(opt *options) {
		opt.namespace = namespace
	}
}

func New(endpoints []string, opts ...Option) (Registry, error) {
	options := &options{
		timeout:     DefaultDuration,
		username:    "",
		password:    "",
		dialTimeout: DefaultDuration,
		cxt:         context.Background(),
	}
	for _, opt := range opts {
		opt(options)
	}
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: options.dialTimeout,
		Username:    options.username,
		Password:    options.password,
	}
	client, err := clientv3.New(config)
	if err != nil {
		return Registry{}, err
	}
	return Registry{
		clientV3:  client,
		clientKv:  clientv3.NewKV(client),
		lease:     clientv3.NewLease(client),
		ctx:       options.cxt,
		namespace: options.namespace,
		ttl:       options.ttl,
	}, nil
}
