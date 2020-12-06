package service

import (
	"fmt"
	"reflect"
	"sync"
)

// ServiceContainer is a wrapper around services indexed by a unique
// name. Services can be retrieved by name, or injected into a struct
// by reading tagged fields.
type ServiceContainer interface {
	// Get retrieves the service registered to the given key. It is an
	// error for a service not to be registered to this key.
	Get(key interface{}) (interface{}, error)

	// MustGet calls Get and panics on error.
	MustGet(service interface{}) interface{}

	// Set registers a service with the given key. It is an error for
	// a service to already be registered to this key.
	Set(key interface{}, service interface{}) error

	// MustSet calls Set and panics on error.
	MustSet(service interface{}, value interface{})

	// Inject will attempt to populate the given type with values from
	// the service container based on the value's struct tags. An error
	// may occur if a service has not been registered, a service has a
	// different type than expected, or struct tags are malformed.
	Inject(obj interface{}) error
}

// InjectableServiceKey is an optional interface for service keys.
//
// Non-string service key values should implement this interface if they
// intend to be injected via struct tags.
type InjectableServiceKey interface {
	// Tag returns the string version of the key. Two distinct key values
	// that return the same value from this method should be considered
	// equivalent within a single service container.
	Tag() string
}

type serviceContainer struct {
	services  map[interface{}]interface{}
	keysByTag map[string]interface{}
	mutex     sync.RWMutex
}

// NewServiceContainer creates an empty service container.
func NewServiceContainer() ServiceContainer {
	return &serviceContainer{
		services:  map[interface{}]interface{}{},
		keysByTag: map[string]interface{}{},
	}
}

// Get retrieves the service registered to the given key. It is an
// error for a service not to be registered to this key.
func (c *serviceContainer) Get(key interface{}) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if service, ok := c.services[key]; ok {
		return service, nil
	}
	if tag, ok := tagForKey(key); ok {
		if key, ok := c.keysByTag[tag]; ok {
			if service, ok := c.services[key]; ok {
				return service, nil
			}
		}
	}

	return nil, fmt.Errorf("no service registered to key %s", prettyKey(key))
}

// MustGet calls Get and panics on error.
func (c *serviceContainer) MustGet(service interface{}) interface{} {
	value, err := c.Get(service)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Set registers a service with the given key. It is an error for
// a service to already be registered to this key.
func (c *serviceContainer) Set(key interface{}, service interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.services[key]; ok {
		return fmt.Errorf(`duplicate service key %s`, prettyKey(key))
	}

	if tag, ok := tagForKey(key); ok {
		if _, ok := c.keysByTag[tag]; ok {
			return fmt.Errorf(`duplicate service key %s`, prettyKey(key))
		}

		c.keysByTag[tag] = key
	}
	c.services[key] = service

	return nil
}

// MustSet calls Set and panics on error.
func (c *serviceContainer) MustSet(service interface{}, value interface{}) {
	if err := c.Set(service, value); err != nil {
		panic(err.Error())
	}
}

// Inject will attempt to populate the given type with values from
// the service container based on the value's struct tags. An error
// may occur if a service has not been registered, a service has a
// different type than expected, or struct tags are malformed.
func (c *serviceContainer) Inject(obj interface{}) error {
	_, err := inject(c, obj, nil, nil)
	return err
}

// TODO - document
// TODO - test
func prettyKey(key interface{}) string {
	if tag, ok := tagForKey(key); ok {
		if _, ok := key.(string); ok {
			return fmt.Sprintf(`"%s"`, key)
		}

		return fmt.Sprintf(`%s ("%s")`, reflect.TypeOf(key).Name(), tag)
	}

	return reflect.TypeOf(key).Name()
}

// tagForKey returns the string version of the given struct key value
// and a boolean flag indicating such a string's existence.
func tagForKey(key interface{}) (string, bool) {
	if k, ok := key.(string); ok {
		return k, true
	}

	if k, ok := key.(InjectableServiceKey); ok {
		return k.Tag(), true
	}

	return "", false
}
