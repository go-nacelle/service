package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceContainerGetAndSet(t *testing.T) {
	container := NewServiceContainer()
	container.Set("a", &IntWrapper{10})
	container.Set("b", &FloatWrapper{3.14})
	container.Set("c", &IntWrapper{25})

	value1, err1 := container.Get("a")
	assert.Nil(t, err1)
	assert.Equal(t, &IntWrapper{10}, value1)

	value2, err2 := container.Get("b")
	assert.Nil(t, err2)
	assert.Equal(t, &FloatWrapper{3.14}, value2)

	value3, err3 := container.Get("c")
	assert.Nil(t, err3)
	assert.Equal(t, &IntWrapper{25}, value3)
}

func TestServiceContainerInject(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestSimpleProcess{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectNonPointer(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", IntWrapper{42})
	obj := &TestSimpleNonPointer{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectAnonymous(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousSimpleProcess{&TestSimpleProcess{}}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectAnonymousZeroValue(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousSimpleProcess{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectAnonymousNonPointer(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousNonPointerSimpleProcess{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectAnonymousDeepNonPointer(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousDeepNonPointerSimpleProcess{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectAnonymousZeroValueNoServiceTags(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousNoServiceTags{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Nil(t, obj.IntWrapper)
}

func TestServiceContainerInjectAnonymousUnexported(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	obj := &TestAnonymousUnexportedProcess{&privateProcess{}}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Nil(t, obj.privateProcess.Value)
}

func TestServiceContainerInjectNonStruct(t *testing.T) {
	container := NewServiceContainer()
	obj := func() error { return nil }
	err := container.Inject(obj)
	assert.Nil(t, err)
}

func TestServiceContainerInjectMissingService(t *testing.T) {
	container := NewServiceContainer()
	obj := &TestSimpleProcess{}
	err := container.Inject(obj)
	assert.EqualError(t, err, "no service registered to key `value`")
}

func TestServiceContainerInjectBadType(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &FloatWrapper{3.14})
	obj := &TestSimpleProcess{}
	err := container.Inject(obj)
	assert.EqualError(t, err, "field 'Value' cannot be assigned a value of type *service.FloatWrapper")
}

func TestServiceContainerInjectNil(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", nil)
	obj := &TestNonPointerField{}
	err := container.Inject(obj)
	assert.EqualError(t, err, "field 'Value' cannot be assigned a value of type nil")
}

func TestServiceContainerInjectOptional(t *testing.T) {
	container := NewServiceContainer()
	obj := &TestOptionalServiceProcess{}
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Nil(t, obj.Value)

	container.Set("value", &IntWrapper{42})
	err = container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42, obj.Value.val)
}

func TestServiceContainerInjectBadOptional(t *testing.T) {
	container := NewServiceContainer()
	obj := &TestBadOptionalServiceProcess{}
	err := container.Inject(obj)
	assert.EqualError(t, err, "field 'Value' has an invalid optional tag")
}

func TestServiceContainerUnsettableFields(t *testing.T) {
	container := NewServiceContainer()
	container.Set("value", &IntWrapper{42})
	err := container.Inject(&TestUnsettableService{})
	assert.EqualError(t, err, "field 'value' can not be set - it may be unexported")
}

func TestServiceContainerPostInject(t *testing.T) {
	container := NewServiceContainer()
	obj := &SimplePostInjectProcess{}
	container.Set("value", &IntWrapper{42})
	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42.0, obj.FValue.val)
}

func TestServiceContainerPostInjectError(t *testing.T) {
	container := NewServiceContainer()
	obj := &ErrorPostInjectProcess{}
	container.Set("value", &IntWrapper{42})
	err := container.Inject(obj)
	assert.EqualError(t, err, "oops")
}

func TestServiceContainerPostInjectChain(t *testing.T) {
	container := NewServiceContainer()
	obj := &RootInjectProcess{}
	process := &SimplePostInjectProcess{}

	container.Set("value", &IntWrapper{42})
	container.Set("process", process)
	container.Set("services", container)

	err := container.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 42.0, process.FValue.val)
}

func TestServiceContainerDuplicateRegistration(t *testing.T) {
	container := NewServiceContainer()
	err1 := container.Set("dup", struct{}{})
	err2 := container.Set("dup", struct{}{})
	assert.Nil(t, err1)
	assert.EqualError(t, err2, "duplicate service key `dup`")
}

func TestServiceContainerGetUnregisteredKey(t *testing.T) {
	container := NewServiceContainer()
	_, err := container.Get("unregistered")
	assert.EqualError(t, err, "no service registered to key `unregistered`")
}

func TestServiceContainerMustSetPanics(t *testing.T) {
	assert.Panics(t, func() {
		container := NewServiceContainer()
		container.MustSet("unregistered", struct{}{})
		container.MustSet("unregistered", struct{}{})
	})
}

func TestServiceContainerMustGetPanics(t *testing.T) {
	assert.Panics(t, func() {
		NewServiceContainer().MustGet("unregistered")
	})
}
