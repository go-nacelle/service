package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlayContainerGet(t *testing.T) {
	container := NewServiceContainer()
	container.Set("a", &IntWrapper{10})
	container.Set("b", &IntWrapper{20})
	container.Set("c", &IntWrapper{30})

	overlay := Overlay(container, map[string]interface{}{
		"a": &IntWrapper{40},
		"d": &IntWrapper{50},
	})

	value1, err1 := overlay.Get("a")
	assert.Nil(t, err1)
	assert.Equal(t, &IntWrapper{40}, value1)

	value2, err2 := overlay.Get("b")
	assert.Nil(t, err2)
	assert.Equal(t, &IntWrapper{20}, value2)

	value3, err3 := overlay.Get("c")
	assert.Nil(t, err3)
	assert.Equal(t, &IntWrapper{30}, value3)

	value4, err4 := overlay.Get("d")
	assert.Nil(t, err4)
	assert.Equal(t, &IntWrapper{50}, value4)
}

func TestOverlayContainerInject(t *testing.T) {
	container := NewServiceContainer()
	container.Set("a", &IntWrapper{10})
	container.Set("b", &IntWrapper{20})
	container.Set("c", &IntWrapper{30})

	overlay := Overlay(container, map[string]interface{}{
		"a": &IntWrapper{40},
		"d": &IntWrapper{50},
	})

	obj := &TestOverlayProcess{}
	err := overlay.Inject(obj)
	assert.Nil(t, err)
	assert.Equal(t, 40, obj.A.val)
	assert.Equal(t, 20, obj.B.val)
	assert.Equal(t, 30, obj.C.val)
	assert.Equal(t, 50, obj.D.val)
}
