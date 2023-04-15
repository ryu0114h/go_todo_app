// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/ryu0114h/go_todo_app/entity"
	"sync"
)

// Ensure, that ListTasksServiceMock does implement ListTasksService.
// If this is not the case, regenerate this file with moq.
var _ ListTasksService = &ListTasksServiceMock{}

// ListTasksServiceMock is a mock implementation of ListTasksService.
//
//	func TestSomethingThatUsesListTasksService(t *testing.T) {
//
//		// make and configure a mocked ListTasksService
//		mockedListTasksService := &ListTasksServiceMock{
//			ListTasksFunc: func(ctx context.Context) (entity.Tasks, error) {
//				panic("mock out the ListTasks method")
//			},
//		}
//
//		// use mockedListTasksService in code that requires ListTasksService
//		// and then make assertions.
//
//	}
type ListTasksServiceMock struct {
	// ListTasksFunc mocks the ListTasks method.
	ListTasksFunc func(ctx context.Context) (entity.Tasks, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListTasks holds details about calls to the ListTasks method.
		ListTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockListTasks sync.RWMutex
}

// ListTasks calls ListTasksFunc.
func (mock *ListTasksServiceMock) ListTasks(ctx context.Context) (entity.Tasks, error) {
	if mock.ListTasksFunc == nil {
		panic("ListTasksServiceMock.ListTasksFunc: method is nil but ListTasksService.ListTasks was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockListTasks.Lock()
	mock.calls.ListTasks = append(mock.calls.ListTasks, callInfo)
	mock.lockListTasks.Unlock()
	return mock.ListTasksFunc(ctx)
}

// ListTasksCalls gets all the calls that were made to ListTasks.
// Check the length with:
//
//	len(mockedListTasksService.ListTasksCalls())
func (mock *ListTasksServiceMock) ListTasksCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockListTasks.RLock()
	calls = mock.calls.ListTasks
	mock.lockListTasks.RUnlock()
	return calls
}

// Ensure, that AddTaskServiceMock does implement AddTaskService.
// If this is not the case, regenerate this file with moq.
var _ AddTaskService = &AddTaskServiceMock{}

// AddTaskServiceMock is a mock implementation of AddTaskService.
//
//	func TestSomethingThatUsesAddTaskService(t *testing.T) {
//
//		// make and configure a mocked AddTaskService
//		mockedAddTaskService := &AddTaskServiceMock{
//			AddTaskFunc: func(ctx context.Context, title string) (*entity.Task, error) {
//				panic("mock out the AddTask method")
//			},
//		}
//
//		// use mockedAddTaskService in code that requires AddTaskService
//		// and then make assertions.
//
//	}
type AddTaskServiceMock struct {
	// AddTaskFunc mocks the AddTask method.
	AddTaskFunc func(ctx context.Context, title string) (*entity.Task, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddTask holds details about calls to the AddTask method.
		AddTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Title is the title argument value.
			Title string
		}
	}
	lockAddTask sync.RWMutex
}

// AddTask calls AddTaskFunc.
func (mock *AddTaskServiceMock) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	if mock.AddTaskFunc == nil {
		panic("AddTaskServiceMock.AddTaskFunc: method is nil but AddTaskService.AddTask was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Title string
	}{
		Ctx:   ctx,
		Title: title,
	}
	mock.lockAddTask.Lock()
	mock.calls.AddTask = append(mock.calls.AddTask, callInfo)
	mock.lockAddTask.Unlock()
	return mock.AddTaskFunc(ctx, title)
}

// AddTaskCalls gets all the calls that were made to AddTask.
// Check the length with:
//
//	len(mockedAddTaskService.AddTaskCalls())
func (mock *AddTaskServiceMock) AddTaskCalls() []struct {
	Ctx   context.Context
	Title string
} {
	var calls []struct {
		Ctx   context.Context
		Title string
	}
	mock.lockAddTask.RLock()
	calls = mock.calls.AddTask
	mock.lockAddTask.RUnlock()
	return calls
}

// Ensure, that RegisterUserServiceMock does implement RegisterUserService.
// If this is not the case, regenerate this file with moq.
var _ RegisterUserService = &RegisterUserServiceMock{}

// RegisterUserServiceMock is a mock implementation of RegisterUserService.
//
//	func TestSomethingThatUsesRegisterUserService(t *testing.T) {
//
//		// make and configure a mocked RegisterUserService
//		mockedRegisterUserService := &RegisterUserServiceMock{
//			RegisterUserFunc: func(ctx context.Context, name string, password string, role string) (*entity.User, error) {
//				panic("mock out the RegisterUser method")
//			},
//		}
//
//		// use mockedRegisterUserService in code that requires RegisterUserService
//		// and then make assertions.
//
//	}
type RegisterUserServiceMock struct {
	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, name string, password string, role string) (*entity.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Password is the password argument value.
			Password string
			// Role is the role argument value.
			Role string
		}
	}
	lockRegisterUser sync.RWMutex
}

// RegisterUser calls RegisterUserFunc.
func (mock *RegisterUserServiceMock) RegisterUser(ctx context.Context, name string, password string, role string) (*entity.User, error) {
	if mock.RegisterUserFunc == nil {
		panic("RegisterUserServiceMock.RegisterUserFunc: method is nil but RegisterUserService.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Name     string
		Password string
		Role     string
	}{
		Ctx:      ctx,
		Name:     name,
		Password: password,
		Role:     role,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, name, password, role)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedRegisterUserService.RegisterUserCalls())
func (mock *RegisterUserServiceMock) RegisterUserCalls() []struct {
	Ctx      context.Context
	Name     string
	Password string
	Role     string
} {
	var calls []struct {
		Ctx      context.Context
		Name     string
		Password string
		Role     string
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// Ensure, that LoginServiceMock does implement LoginService.
// If this is not the case, regenerate this file with moq.
var _ LoginService = &LoginServiceMock{}

// LoginServiceMock is a mock implementation of LoginService.
//
//	func TestSomethingThatUsesLoginService(t *testing.T) {
//
//		// make and configure a mocked LoginService
//		mockedLoginService := &LoginServiceMock{
//			LoginFunc: func(ctx context.Context, name string, pw string) (string, error) {
//				panic("mock out the Login method")
//			},
//		}
//
//		// use mockedLoginService in code that requires LoginService
//		// and then make assertions.
//
//	}
type LoginServiceMock struct {
	// LoginFunc mocks the Login method.
	LoginFunc func(ctx context.Context, name string, pw string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Login holds details about calls to the Login method.
		Login []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Pw is the pw argument value.
			Pw string
		}
	}
	lockLogin sync.RWMutex
}

// Login calls LoginFunc.
func (mock *LoginServiceMock) Login(ctx context.Context, name string, pw string) (string, error) {
	if mock.LoginFunc == nil {
		panic("LoginServiceMock.LoginFunc: method is nil but LoginService.Login was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
		Pw   string
	}{
		Ctx:  ctx,
		Name: name,
		Pw:   pw,
	}
	mock.lockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	mock.lockLogin.Unlock()
	return mock.LoginFunc(ctx, name, pw)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//
//	len(mockedLoginService.LoginCalls())
func (mock *LoginServiceMock) LoginCalls() []struct {
	Ctx  context.Context
	Name string
	Pw   string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
		Pw   string
	}
	mock.lockLogin.RLock()
	calls = mock.calls.Login
	mock.lockLogin.RUnlock()
	return calls
}

// Ensure, that DeleteTaskServiceMock does implement DeleteTaskService.
// If this is not the case, regenerate this file with moq.
var _ DeleteTaskService = &DeleteTaskServiceMock{}

// DeleteTaskServiceMock is a mock implementation of DeleteTaskService.
//
//	func TestSomethingThatUsesDeleteTaskService(t *testing.T) {
//
//		// make and configure a mocked DeleteTaskService
//		mockedDeleteTaskService := &DeleteTaskServiceMock{
//			DeleteTaskFunc: func(ctx context.Context, id entity.TaskID) error {
//				panic("mock out the DeleteTask method")
//			},
//		}
//
//		// use mockedDeleteTaskService in code that requires DeleteTaskService
//		// and then make assertions.
//
//	}
type DeleteTaskServiceMock struct {
	// DeleteTaskFunc mocks the DeleteTask method.
	DeleteTaskFunc func(ctx context.Context, id entity.TaskID) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteTask holds details about calls to the DeleteTask method.
		DeleteTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID entity.TaskID
		}
	}
	lockDeleteTask sync.RWMutex
}

// DeleteTask calls DeleteTaskFunc.
func (mock *DeleteTaskServiceMock) DeleteTask(ctx context.Context, id entity.TaskID) error {
	if mock.DeleteTaskFunc == nil {
		panic("DeleteTaskServiceMock.DeleteTaskFunc: method is nil but DeleteTaskService.DeleteTask was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  entity.TaskID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteTask.Lock()
	mock.calls.DeleteTask = append(mock.calls.DeleteTask, callInfo)
	mock.lockDeleteTask.Unlock()
	return mock.DeleteTaskFunc(ctx, id)
}

// DeleteTaskCalls gets all the calls that were made to DeleteTask.
// Check the length with:
//
//	len(mockedDeleteTaskService.DeleteTaskCalls())
func (mock *DeleteTaskServiceMock) DeleteTaskCalls() []struct {
	Ctx context.Context
	ID  entity.TaskID
} {
	var calls []struct {
		Ctx context.Context
		ID  entity.TaskID
	}
	mock.lockDeleteTask.RLock()
	calls = mock.calls.DeleteTask
	mock.lockDeleteTask.RUnlock()
	return calls
}
