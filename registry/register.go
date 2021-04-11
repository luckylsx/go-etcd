package registry

import (
	"encoding/json"
	"fmt"

	"go.etcd.io/etcd/clientv3"
)

// Register register data to etcd
func (r Registry) Register(instance ServiceInstance) error {
	key := r.getKey(instance)
	value, err := json.Marshal(instance)
	if err != nil {
		return err
	}
	if r.ttl > 0 {
		grant, _ := r.lease.Grant(r.ctx, int64(r.ttl))
		_, err = r.clientV3.Put(r.ctx, key, string(value), clientv3.WithLease(grant.ID))
	} else {
		_, err = r.clientV3.Put(r.ctx, key, string(value))
	}
	if err != nil {
		return err
	}
	return nil
}

// DeRegister unregister from etcd
func (r Registry) DeRegister(key string, instance ServiceInstance) error {
	defer func() {
		if r.lease != nil {
			_ = r.lease.Close()
		}
	}()
	key = r.getKey(instance)
	_, err := r.clientV3.Delete(r.ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (r Registry) getKey(instance ServiceInstance) string {
	return fmt.Sprintf("%s/%s/%s", r.namespace, instance.Name, instance.ID)
}

// GetService getService form etcd
func (r Registry) GetService(instance ServiceInstance) ([]string, error) {
	key := r.getKey(instance)
	res, err := r.clientV3.Get(r.ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	itmes := make([]string, 0, len(res.Kvs))
	for _, v := range res.Kvs {
		itmes = append(itmes, string(v.Value))
	}
	return itmes, nil
}
