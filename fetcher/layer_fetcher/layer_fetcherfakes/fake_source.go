// Code generated by counterfeiter. DO NOT EDIT.
package layer_fetcherfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher/layer_fetcher"
	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
	"github.com/containers/image/types"
)

type FakeSource struct {
	BlobStub        func(lager.Logger, groot.LayerInfo) (string, int64, error)
	blobMutex       sync.RWMutex
	blobArgsForCall []struct {
		arg1 lager.Logger
		arg2 groot.LayerInfo
	}
	blobReturns struct {
		result1 string
		result2 int64
		result3 error
	}
	blobReturnsOnCall map[int]struct {
		result1 string
		result2 int64
		result3 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ManifestStub        func(lager.Logger) (types.Image, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
		arg1 lager.Logger
	}
	manifestReturns struct {
		result1 types.Image
		result2 error
	}
	manifestReturnsOnCall map[int]struct {
		result1 types.Image
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSource) Blob(arg1 lager.Logger, arg2 groot.LayerInfo) (string, int64, error) {
	fake.blobMutex.Lock()
	ret, specificReturn := fake.blobReturnsOnCall[len(fake.blobArgsForCall)]
	fake.blobArgsForCall = append(fake.blobArgsForCall, struct {
		arg1 lager.Logger
		arg2 groot.LayerInfo
	}{arg1, arg2})
	fake.recordInvocation("Blob", []interface{}{arg1, arg2})
	fake.blobMutex.Unlock()
	if fake.BlobStub != nil {
		return fake.BlobStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.blobReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSource) BlobCallCount() int {
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	return len(fake.blobArgsForCall)
}

func (fake *FakeSource) BlobCalls(stub func(lager.Logger, groot.LayerInfo) (string, int64, error)) {
	fake.blobMutex.Lock()
	defer fake.blobMutex.Unlock()
	fake.BlobStub = stub
}

func (fake *FakeSource) BlobArgsForCall(i int) (lager.Logger, groot.LayerInfo) {
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	argsForCall := fake.blobArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSource) BlobReturns(result1 string, result2 int64, result3 error) {
	fake.blobMutex.Lock()
	defer fake.blobMutex.Unlock()
	fake.BlobStub = nil
	fake.blobReturns = struct {
		result1 string
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSource) BlobReturnsOnCall(i int, result1 string, result2 int64, result3 error) {
	fake.blobMutex.Lock()
	defer fake.blobMutex.Unlock()
	fake.BlobStub = nil
	if fake.blobReturnsOnCall == nil {
		fake.blobReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int64
			result3 error
		})
	}
	fake.blobReturnsOnCall[i] = struct {
		result1 string
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSource) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeSource) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSource) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSource) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSource) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSource) Manifest(arg1 lager.Logger) (types.Image, error) {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Manifest", []interface{}{arg1})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.manifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSource) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeSource) ManifestCalls(stub func(lager.Logger) (types.Image, error)) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = stub
}

func (fake *FakeSource) ManifestArgsForCall(i int) lager.Logger {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	argsForCall := fake.manifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSource) ManifestReturns(result1 types.Image, result2 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ManifestReturnsOnCall(i int, result1 types.Image, result2 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 types.Image
			result2 error
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSource) recordInvocation(key string, args []interface{}) {
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

var _ layer_fetcher.Source = new(FakeSource)
