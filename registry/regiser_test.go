package registry

import (
	"context"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestRegister(t *testing.T) {
	points := []string{"http://127.0.0.1:2379"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	registry, err := New(points, WithNameSpace("microservice"), WithCtx(ctx))
	instance := ServiceInstance{
		ID:   "1",
		Name: "test",
	}
	if err != nil {
		assert.NilError(t, err)
	}
	assert.NilError(t, err)
	err = registry.Register(instance)
	if err != nil {
		t.Error(err)
	}
	cancel()
	assert.NilError(t, err)
}

func TestRegistry_GetService(t *testing.T) {
	points := []string{"http://127.0.0.1:2379"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	registry, err := New(points, WithNameSpace("microservice"), WithCtx(ctx))
	instance := ServiceInstance{
		ID:   "1",
		Name: "test",
	}
	if err != nil {
		assert.NilError(t, err)
	}
	assert.NilError(t, err)
	res, err := registry.GetService(instance)
	if err != nil {
		t.Error(err)
	}
	cancel()
	t.Logf("res = %+v", res)
	assert.NilError(t, err)
}
