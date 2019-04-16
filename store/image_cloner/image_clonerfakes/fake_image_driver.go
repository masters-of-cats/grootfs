// Code generated by counterfeiter. DO NOT EDIT.
package image_clonerfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/grootfs/store/image_cloner"
	"code.cloudfoundry.org/lager"
)

type FakeImageDriver struct {
	CreateImageStub        func(lager.Logger, image_cloner.ImageDriverSpec) (groot.MountInfo, error)
	createImageMutex       sync.RWMutex
	createImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 image_cloner.ImageDriverSpec
	}
	createImageReturns struct {
		result1 groot.MountInfo
		result2 error
	}
	createImageReturnsOnCall map[int]struct {
		result1 groot.MountInfo
		result2 error
	}
	DestroyImageStub        func(lager.Logger, string) error
	destroyImageMutex       sync.RWMutex
	destroyImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	destroyImageReturns struct {
		result1 error
	}
	destroyImageReturnsOnCall map[int]struct {
		result1 error
	}
	FetchStatsStub        func(lager.Logger, string) (groot.VolumeStats, error)
	fetchStatsMutex       sync.RWMutex
	fetchStatsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	fetchStatsReturns struct {
		result1 groot.VolumeStats
		result2 error
	}
	fetchStatsReturnsOnCall map[int]struct {
		result1 groot.VolumeStats
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageDriver) CreateImage(arg1 lager.Logger, arg2 image_cloner.ImageDriverSpec) (groot.MountInfo, error) {
	fake.createImageMutex.Lock()
	ret, specificReturn := fake.createImageReturnsOnCall[len(fake.createImageArgsForCall)]
	fake.createImageArgsForCall = append(fake.createImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 image_cloner.ImageDriverSpec
	}{arg1, arg2})
	fake.recordInvocation("CreateImage", []interface{}{arg1, arg2})
	fake.createImageMutex.Unlock()
	if fake.CreateImageStub != nil {
		return fake.CreateImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageDriver) CreateImageCallCount() int {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	return len(fake.createImageArgsForCall)
}

func (fake *FakeImageDriver) CreateImageCalls(stub func(lager.Logger, image_cloner.ImageDriverSpec) (groot.MountInfo, error)) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = stub
}

func (fake *FakeImageDriver) CreateImageArgsForCall(i int) (lager.Logger, image_cloner.ImageDriverSpec) {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	argsForCall := fake.createImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageDriver) CreateImageReturns(result1 groot.MountInfo, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	fake.createImageReturns = struct {
		result1 groot.MountInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImageDriver) CreateImageReturnsOnCall(i int, result1 groot.MountInfo, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	if fake.createImageReturnsOnCall == nil {
		fake.createImageReturnsOnCall = make(map[int]struct {
			result1 groot.MountInfo
			result2 error
		})
	}
	fake.createImageReturnsOnCall[i] = struct {
		result1 groot.MountInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImageDriver) DestroyImage(arg1 lager.Logger, arg2 string) error {
	fake.destroyImageMutex.Lock()
	ret, specificReturn := fake.destroyImageReturnsOnCall[len(fake.destroyImageArgsForCall)]
	fake.destroyImageArgsForCall = append(fake.destroyImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DestroyImage", []interface{}{arg1, arg2})
	fake.destroyImageMutex.Unlock()
	if fake.DestroyImageStub != nil {
		return fake.DestroyImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyImageReturns
	return fakeReturns.result1
}

func (fake *FakeImageDriver) DestroyImageCallCount() int {
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	return len(fake.destroyImageArgsForCall)
}

func (fake *FakeImageDriver) DestroyImageCalls(stub func(lager.Logger, string) error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = stub
}

func (fake *FakeImageDriver) DestroyImageArgsForCall(i int) (lager.Logger, string) {
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	argsForCall := fake.destroyImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageDriver) DestroyImageReturns(result1 error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = nil
	fake.destroyImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageDriver) DestroyImageReturnsOnCall(i int, result1 error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = nil
	if fake.destroyImageReturnsOnCall == nil {
		fake.destroyImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageDriver) FetchStats(arg1 lager.Logger, arg2 string) (groot.VolumeStats, error) {
	fake.fetchStatsMutex.Lock()
	ret, specificReturn := fake.fetchStatsReturnsOnCall[len(fake.fetchStatsArgsForCall)]
	fake.fetchStatsArgsForCall = append(fake.fetchStatsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FetchStats", []interface{}{arg1, arg2})
	fake.fetchStatsMutex.Unlock()
	if fake.FetchStatsStub != nil {
		return fake.FetchStatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchStatsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageDriver) FetchStatsCallCount() int {
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	return len(fake.fetchStatsArgsForCall)
}

func (fake *FakeImageDriver) FetchStatsCalls(stub func(lager.Logger, string) (groot.VolumeStats, error)) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = stub
}

func (fake *FakeImageDriver) FetchStatsArgsForCall(i int) (lager.Logger, string) {
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	argsForCall := fake.fetchStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageDriver) FetchStatsReturns(result1 groot.VolumeStats, result2 error) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = nil
	fake.fetchStatsReturns = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeImageDriver) FetchStatsReturnsOnCall(i int, result1 groot.VolumeStats, result2 error) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = nil
	if fake.fetchStatsReturnsOnCall == nil {
		fake.fetchStatsReturnsOnCall = make(map[int]struct {
			result1 groot.VolumeStats
			result2 error
		})
	}
	fake.fetchStatsReturnsOnCall[i] = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeImageDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageDriver) recordInvocation(key string, args []interface{}) {
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

var _ image_cloner.ImageDriver = new(FakeImageDriver)
