// Code generated by counterfeiter. DO NOT EDIT.
package grootfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
)

type FakeBaseImagePuller struct {
	FetchBaseImageInfoStub        func(lager.Logger) (groot.BaseImageInfo, error)
	fetchBaseImageInfoMutex       sync.RWMutex
	fetchBaseImageInfoArgsForCall []struct {
		arg1 lager.Logger
	}
	fetchBaseImageInfoReturns struct {
		result1 groot.BaseImageInfo
		result2 error
	}
	fetchBaseImageInfoReturnsOnCall map[int]struct {
		result1 groot.BaseImageInfo
		result2 error
	}
	PullStub        func(lager.Logger, groot.BaseImageInfo, groot.BaseImageSpec) error
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		arg1 lager.Logger
		arg2 groot.BaseImageInfo
		arg3 groot.BaseImageSpec
	}
	pullReturns struct {
		result1 error
	}
	pullReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfo(arg1 lager.Logger) (groot.BaseImageInfo, error) {
	fake.fetchBaseImageInfoMutex.Lock()
	ret, specificReturn := fake.fetchBaseImageInfoReturnsOnCall[len(fake.fetchBaseImageInfoArgsForCall)]
	fake.fetchBaseImageInfoArgsForCall = append(fake.fetchBaseImageInfoArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("FetchBaseImageInfo", []interface{}{arg1})
	fake.fetchBaseImageInfoMutex.Unlock()
	if fake.FetchBaseImageInfoStub != nil {
		return fake.FetchBaseImageInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchBaseImageInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfoCallCount() int {
	fake.fetchBaseImageInfoMutex.RLock()
	defer fake.fetchBaseImageInfoMutex.RUnlock()
	return len(fake.fetchBaseImageInfoArgsForCall)
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfoCalls(stub func(lager.Logger) (groot.BaseImageInfo, error)) {
	fake.fetchBaseImageInfoMutex.Lock()
	defer fake.fetchBaseImageInfoMutex.Unlock()
	fake.FetchBaseImageInfoStub = stub
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfoArgsForCall(i int) lager.Logger {
	fake.fetchBaseImageInfoMutex.RLock()
	defer fake.fetchBaseImageInfoMutex.RUnlock()
	argsForCall := fake.fetchBaseImageInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfoReturns(result1 groot.BaseImageInfo, result2 error) {
	fake.fetchBaseImageInfoMutex.Lock()
	defer fake.fetchBaseImageInfoMutex.Unlock()
	fake.FetchBaseImageInfoStub = nil
	fake.fetchBaseImageInfoReturns = struct {
		result1 groot.BaseImageInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseImagePuller) FetchBaseImageInfoReturnsOnCall(i int, result1 groot.BaseImageInfo, result2 error) {
	fake.fetchBaseImageInfoMutex.Lock()
	defer fake.fetchBaseImageInfoMutex.Unlock()
	fake.FetchBaseImageInfoStub = nil
	if fake.fetchBaseImageInfoReturnsOnCall == nil {
		fake.fetchBaseImageInfoReturnsOnCall = make(map[int]struct {
			result1 groot.BaseImageInfo
			result2 error
		})
	}
	fake.fetchBaseImageInfoReturnsOnCall[i] = struct {
		result1 groot.BaseImageInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeBaseImagePuller) Pull(arg1 lager.Logger, arg2 groot.BaseImageInfo, arg3 groot.BaseImageSpec) error {
	fake.pullMutex.Lock()
	ret, specificReturn := fake.pullReturnsOnCall[len(fake.pullArgsForCall)]
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		arg1 lager.Logger
		arg2 groot.BaseImageInfo
		arg3 groot.BaseImageSpec
	}{arg1, arg2, arg3})
	fake.recordInvocation("Pull", []interface{}{arg1, arg2, arg3})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pullReturns
	return fakeReturns.result1
}

func (fake *FakeBaseImagePuller) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeBaseImagePuller) PullCalls(stub func(lager.Logger, groot.BaseImageInfo, groot.BaseImageSpec) error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = stub
}

func (fake *FakeBaseImagePuller) PullArgsForCall(i int) (lager.Logger, groot.BaseImageInfo, groot.BaseImageSpec) {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	argsForCall := fake.pullArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBaseImagePuller) PullReturns(result1 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBaseImagePuller) PullReturnsOnCall(i int, result1 error) {
	fake.pullMutex.Lock()
	defer fake.pullMutex.Unlock()
	fake.PullStub = nil
	if fake.pullReturnsOnCall == nil {
		fake.pullReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pullReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBaseImagePuller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchBaseImageInfoMutex.RLock()
	defer fake.fetchBaseImageInfoMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBaseImagePuller) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ groot.BaseImagePuller = new(FakeBaseImagePuller)
