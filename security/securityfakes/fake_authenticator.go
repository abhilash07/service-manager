// Code generated by counterfeiter. DO NOT EDIT.
package securityfakes

import (
	"net/http"
	"sync"

	"github.com/Peripli/service-manager/pkg/web"
	"github.com/Peripli/service-manager/security"
)

type FakeAuthenticator struct {
	AuthenticateStub        func(req *http.Request) (*web.User, security.AuthenticationDecision, error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		req *http.Request
	}
	authenticateReturns struct {
		result1 *web.User
		result2 security.AuthenticationDecision
		result3 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 *web.User
		result2 security.AuthenticationDecision
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticator) Authenticate(req *http.Request) (*web.User, security.AuthenticationDecision, error) {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		req *http.Request
	}{req})
	fake.recordInvocation("Authenticate", []interface{}{req})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(req)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.authenticateReturns.result1, fake.authenticateReturns.result2, fake.authenticateReturns.result3
}

func (fake *FakeAuthenticator) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeAuthenticator) AuthenticateArgsForCall(i int) *http.Request {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.authenticateArgsForCall[i].req
}

func (fake *FakeAuthenticator) AuthenticateReturns(result1 *web.User, result2 security.AuthenticationDecision, result3 error) {
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 *web.User
		result2 security.AuthenticationDecision
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAuthenticator) AuthenticateReturnsOnCall(i int, result1 *web.User, result2 security.AuthenticationDecision, result3 error) {
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 *web.User
			result2 security.AuthenticationDecision
			result3 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 *web.User
		result2 security.AuthenticationDecision
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAuthenticator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthenticator) recordInvocation(key string, args []interface{}) {
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

var _ security.Authenticator = new(FakeAuthenticator)
