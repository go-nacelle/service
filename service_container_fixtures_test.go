package service

import "fmt"

type IntWrapper struct {
	val int
}

type FloatWrapper struct {
	val float64
}

type TestSimpleNonPointer struct {
	Value IntWrapper `service:"value"`
}

type TestSimpleProcess struct {
	Value *IntWrapper `service:"value"`
}

type TestAnonymousSimpleProcess struct {
	*TestSimpleProcess
}

type TestAnonymousNonPointerSimpleProcess struct {
	TestSimpleProcess
}

type TestAnonymousDeepNonPointerSimpleProcess struct {
	TestSimpleProcess
}

type TestAnonymousNoServiceTags struct {
	*IntWrapper
}

type TestAnonymousUnexportedProcess struct {
	*privateProcess
}

type privateProcess struct {
	Value *IntWrapper `service:"value"`
}

type TestUnsettableService struct {
	value *IntWrapper `service:"value"`
}

type TestNonPointerField struct {
	Value IntWrapper `service:"value"`
}

type TestOptionalServiceProcess struct {
	Value *IntWrapper `service:"value" optional:"true"`
}

type TestBadOptionalServiceProcess struct {
	Value *IntWrapper `service:"value" optional:"yup"`
}

type SimplePostInjectProcess struct {
	IValue *IntWrapper `service:"value"`
	FValue *FloatWrapper
}

func (p *SimplePostInjectProcess) PostInject() error {
	p.FValue = &FloatWrapper{float64(p.IValue.val)}
	return nil
}

type ErrorPostInjectProcess struct{}

func (p *ErrorPostInjectProcess) PostInject() error {
	return fmt.Errorf("oops")
}

type RootInjectProcess struct {
	Services ServiceContainer         `service:"services"`
	Child    *SimplePostInjectProcess `service:"process"`
}

func (p *RootInjectProcess) PostInject() error {
	return p.Services.Inject(p.Child)
}
