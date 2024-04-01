// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/polshe-v/microservices_chat_server/internal/model"
)

// MessagesRepositoryMock implements repository.MessagesRepository
type MessagesRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, chatID string, message *model.Message) (err error)
	inspectFuncCreate   func(ctx context.Context, chatID string, message *model.Message)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mMessagesRepositoryMockCreate

	funcGetMessages          func(ctx context.Context, chatID string) (mpa1 []*model.Message, err error)
	inspectFuncGetMessages   func(ctx context.Context, chatID string)
	afterGetMessagesCounter  uint64
	beforeGetMessagesCounter uint64
	GetMessagesMock          mMessagesRepositoryMockGetMessages
}

// NewMessagesRepositoryMock returns a mock for repository.MessagesRepository
func NewMessagesRepositoryMock(t minimock.Tester) *MessagesRepositoryMock {
	m := &MessagesRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mMessagesRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*MessagesRepositoryMockCreateParams{}

	m.GetMessagesMock = mMessagesRepositoryMockGetMessages{mock: m}
	m.GetMessagesMock.callArgs = []*MessagesRepositoryMockGetMessagesParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessagesRepositoryMockCreate struct {
	mock               *MessagesRepositoryMock
	defaultExpectation *MessagesRepositoryMockCreateExpectation
	expectations       []*MessagesRepositoryMockCreateExpectation

	callArgs []*MessagesRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// MessagesRepositoryMockCreateExpectation specifies expectation struct of the MessagesRepository.Create
type MessagesRepositoryMockCreateExpectation struct {
	mock    *MessagesRepositoryMock
	params  *MessagesRepositoryMockCreateParams
	results *MessagesRepositoryMockCreateResults
	Counter uint64
}

// MessagesRepositoryMockCreateParams contains parameters of the MessagesRepository.Create
type MessagesRepositoryMockCreateParams struct {
	ctx     context.Context
	chatID  string
	message *model.Message
}

// MessagesRepositoryMockCreateResults contains results of the MessagesRepository.Create
type MessagesRepositoryMockCreateResults struct {
	err error
}

// Expect sets up expected params for MessagesRepository.Create
func (mmCreate *mMessagesRepositoryMockCreate) Expect(ctx context.Context, chatID string, message *model.Message) *mMessagesRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessagesRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MessagesRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &MessagesRepositoryMockCreateParams{ctx, chatID, message}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the MessagesRepository.Create
func (mmCreate *mMessagesRepositoryMockCreate) Inspect(f func(ctx context.Context, chatID string, message *model.Message)) *mMessagesRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for MessagesRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by MessagesRepository.Create
func (mmCreate *mMessagesRepositoryMockCreate) Return(err error) *MessagesRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessagesRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MessagesRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &MessagesRepositoryMockCreateResults{err}
	return mmCreate.mock
}

// Set uses given function f to mock the MessagesRepository.Create method
func (mmCreate *mMessagesRepositoryMockCreate) Set(f func(ctx context.Context, chatID string, message *model.Message) (err error)) *MessagesRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the MessagesRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the MessagesRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the MessagesRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mMessagesRepositoryMockCreate) When(ctx context.Context, chatID string, message *model.Message) *MessagesRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessagesRepositoryMock.Create mock is already set by Set")
	}

	expectation := &MessagesRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &MessagesRepositoryMockCreateParams{ctx, chatID, message},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up MessagesRepository.Create return parameters for the expectation previously defined by the When method
func (e *MessagesRepositoryMockCreateExpectation) Then(err error) *MessagesRepositoryMock {
	e.results = &MessagesRepositoryMockCreateResults{err}
	return e.mock
}

// Create implements repository.MessagesRepository
func (mmCreate *MessagesRepositoryMock) Create(ctx context.Context, chatID string, message *model.Message) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, chatID, message)
	}

	mm_params := MessagesRepositoryMockCreateParams{ctx, chatID, message}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := MessagesRepositoryMockCreateParams{ctx, chatID, message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("MessagesRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the MessagesRepositoryMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, chatID, message)
	}
	mmCreate.t.Fatalf("Unexpected call to MessagesRepositoryMock.Create. %v %v %v", ctx, chatID, message)
	return
}

// CreateAfterCounter returns a count of finished MessagesRepositoryMock.Create invocations
func (mmCreate *MessagesRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of MessagesRepositoryMock.Create invocations
func (mmCreate *MessagesRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to MessagesRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mMessagesRepositoryMockCreate) Calls() []*MessagesRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*MessagesRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *MessagesRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *MessagesRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessagesRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessagesRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to MessagesRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to MessagesRepositoryMock.Create")
	}
}

type mMessagesRepositoryMockGetMessages struct {
	mock               *MessagesRepositoryMock
	defaultExpectation *MessagesRepositoryMockGetMessagesExpectation
	expectations       []*MessagesRepositoryMockGetMessagesExpectation

	callArgs []*MessagesRepositoryMockGetMessagesParams
	mutex    sync.RWMutex
}

// MessagesRepositoryMockGetMessagesExpectation specifies expectation struct of the MessagesRepository.GetMessages
type MessagesRepositoryMockGetMessagesExpectation struct {
	mock    *MessagesRepositoryMock
	params  *MessagesRepositoryMockGetMessagesParams
	results *MessagesRepositoryMockGetMessagesResults
	Counter uint64
}

// MessagesRepositoryMockGetMessagesParams contains parameters of the MessagesRepository.GetMessages
type MessagesRepositoryMockGetMessagesParams struct {
	ctx    context.Context
	chatID string
}

// MessagesRepositoryMockGetMessagesResults contains results of the MessagesRepository.GetMessages
type MessagesRepositoryMockGetMessagesResults struct {
	mpa1 []*model.Message
	err  error
}

// Expect sets up expected params for MessagesRepository.GetMessages
func (mmGetMessages *mMessagesRepositoryMockGetMessages) Expect(ctx context.Context, chatID string) *mMessagesRepositoryMockGetMessages {
	if mmGetMessages.mock.funcGetMessages != nil {
		mmGetMessages.mock.t.Fatalf("MessagesRepositoryMock.GetMessages mock is already set by Set")
	}

	if mmGetMessages.defaultExpectation == nil {
		mmGetMessages.defaultExpectation = &MessagesRepositoryMockGetMessagesExpectation{}
	}

	mmGetMessages.defaultExpectation.params = &MessagesRepositoryMockGetMessagesParams{ctx, chatID}
	for _, e := range mmGetMessages.expectations {
		if minimock.Equal(e.params, mmGetMessages.defaultExpectation.params) {
			mmGetMessages.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetMessages.defaultExpectation.params)
		}
	}

	return mmGetMessages
}

// Inspect accepts an inspector function that has same arguments as the MessagesRepository.GetMessages
func (mmGetMessages *mMessagesRepositoryMockGetMessages) Inspect(f func(ctx context.Context, chatID string)) *mMessagesRepositoryMockGetMessages {
	if mmGetMessages.mock.inspectFuncGetMessages != nil {
		mmGetMessages.mock.t.Fatalf("Inspect function is already set for MessagesRepositoryMock.GetMessages")
	}

	mmGetMessages.mock.inspectFuncGetMessages = f

	return mmGetMessages
}

// Return sets up results that will be returned by MessagesRepository.GetMessages
func (mmGetMessages *mMessagesRepositoryMockGetMessages) Return(mpa1 []*model.Message, err error) *MessagesRepositoryMock {
	if mmGetMessages.mock.funcGetMessages != nil {
		mmGetMessages.mock.t.Fatalf("MessagesRepositoryMock.GetMessages mock is already set by Set")
	}

	if mmGetMessages.defaultExpectation == nil {
		mmGetMessages.defaultExpectation = &MessagesRepositoryMockGetMessagesExpectation{mock: mmGetMessages.mock}
	}
	mmGetMessages.defaultExpectation.results = &MessagesRepositoryMockGetMessagesResults{mpa1, err}
	return mmGetMessages.mock
}

// Set uses given function f to mock the MessagesRepository.GetMessages method
func (mmGetMessages *mMessagesRepositoryMockGetMessages) Set(f func(ctx context.Context, chatID string) (mpa1 []*model.Message, err error)) *MessagesRepositoryMock {
	if mmGetMessages.defaultExpectation != nil {
		mmGetMessages.mock.t.Fatalf("Default expectation is already set for the MessagesRepository.GetMessages method")
	}

	if len(mmGetMessages.expectations) > 0 {
		mmGetMessages.mock.t.Fatalf("Some expectations are already set for the MessagesRepository.GetMessages method")
	}

	mmGetMessages.mock.funcGetMessages = f
	return mmGetMessages.mock
}

// When sets expectation for the MessagesRepository.GetMessages which will trigger the result defined by the following
// Then helper
func (mmGetMessages *mMessagesRepositoryMockGetMessages) When(ctx context.Context, chatID string) *MessagesRepositoryMockGetMessagesExpectation {
	if mmGetMessages.mock.funcGetMessages != nil {
		mmGetMessages.mock.t.Fatalf("MessagesRepositoryMock.GetMessages mock is already set by Set")
	}

	expectation := &MessagesRepositoryMockGetMessagesExpectation{
		mock:   mmGetMessages.mock,
		params: &MessagesRepositoryMockGetMessagesParams{ctx, chatID},
	}
	mmGetMessages.expectations = append(mmGetMessages.expectations, expectation)
	return expectation
}

// Then sets up MessagesRepository.GetMessages return parameters for the expectation previously defined by the When method
func (e *MessagesRepositoryMockGetMessagesExpectation) Then(mpa1 []*model.Message, err error) *MessagesRepositoryMock {
	e.results = &MessagesRepositoryMockGetMessagesResults{mpa1, err}
	return e.mock
}

// GetMessages implements repository.MessagesRepository
func (mmGetMessages *MessagesRepositoryMock) GetMessages(ctx context.Context, chatID string) (mpa1 []*model.Message, err error) {
	mm_atomic.AddUint64(&mmGetMessages.beforeGetMessagesCounter, 1)
	defer mm_atomic.AddUint64(&mmGetMessages.afterGetMessagesCounter, 1)

	if mmGetMessages.inspectFuncGetMessages != nil {
		mmGetMessages.inspectFuncGetMessages(ctx, chatID)
	}

	mm_params := MessagesRepositoryMockGetMessagesParams{ctx, chatID}

	// Record call args
	mmGetMessages.GetMessagesMock.mutex.Lock()
	mmGetMessages.GetMessagesMock.callArgs = append(mmGetMessages.GetMessagesMock.callArgs, &mm_params)
	mmGetMessages.GetMessagesMock.mutex.Unlock()

	for _, e := range mmGetMessages.GetMessagesMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.mpa1, e.results.err
		}
	}

	if mmGetMessages.GetMessagesMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetMessages.GetMessagesMock.defaultExpectation.Counter, 1)
		mm_want := mmGetMessages.GetMessagesMock.defaultExpectation.params
		mm_got := MessagesRepositoryMockGetMessagesParams{ctx, chatID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetMessages.t.Errorf("MessagesRepositoryMock.GetMessages got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetMessages.GetMessagesMock.defaultExpectation.results
		if mm_results == nil {
			mmGetMessages.t.Fatal("No results are set for the MessagesRepositoryMock.GetMessages")
		}
		return (*mm_results).mpa1, (*mm_results).err
	}
	if mmGetMessages.funcGetMessages != nil {
		return mmGetMessages.funcGetMessages(ctx, chatID)
	}
	mmGetMessages.t.Fatalf("Unexpected call to MessagesRepositoryMock.GetMessages. %v %v", ctx, chatID)
	return
}

// GetMessagesAfterCounter returns a count of finished MessagesRepositoryMock.GetMessages invocations
func (mmGetMessages *MessagesRepositoryMock) GetMessagesAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetMessages.afterGetMessagesCounter)
}

// GetMessagesBeforeCounter returns a count of MessagesRepositoryMock.GetMessages invocations
func (mmGetMessages *MessagesRepositoryMock) GetMessagesBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetMessages.beforeGetMessagesCounter)
}

// Calls returns a list of arguments used in each call to MessagesRepositoryMock.GetMessages.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetMessages *mMessagesRepositoryMockGetMessages) Calls() []*MessagesRepositoryMockGetMessagesParams {
	mmGetMessages.mutex.RLock()

	argCopy := make([]*MessagesRepositoryMockGetMessagesParams, len(mmGetMessages.callArgs))
	copy(argCopy, mmGetMessages.callArgs)

	mmGetMessages.mutex.RUnlock()

	return argCopy
}

// MinimockGetMessagesDone returns true if the count of the GetMessages invocations corresponds
// the number of defined expectations
func (m *MessagesRepositoryMock) MinimockGetMessagesDone() bool {
	for _, e := range m.GetMessagesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMessagesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetMessagesCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetMessages != nil && mm_atomic.LoadUint64(&m.afterGetMessagesCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetMessagesInspect logs each unmet expectation
func (m *MessagesRepositoryMock) MinimockGetMessagesInspect() {
	for _, e := range m.GetMessagesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessagesRepositoryMock.GetMessages with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMessagesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetMessagesCounter) < 1 {
		if m.GetMessagesMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessagesRepositoryMock.GetMessages")
		} else {
			m.t.Errorf("Expected call to MessagesRepositoryMock.GetMessages with params: %#v", *m.GetMessagesMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetMessages != nil && mm_atomic.LoadUint64(&m.afterGetMessagesCounter) < 1 {
		m.t.Error("Expected call to MessagesRepositoryMock.GetMessages")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessagesRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetMessagesInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessagesRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessagesRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetMessagesDone()
}