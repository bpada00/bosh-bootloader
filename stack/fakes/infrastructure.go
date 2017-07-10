// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-bootloader/aws/cloudformation"
)

type Infrastructure struct {
	UpdateStub        func(keyPairName string, azs []string, stackName, boshAZ, lbType, lbCertificateARN, envID string) (cloudformation.Stack, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		keyPairName      string
		azs              []string
		stackName        string
		boshAZ           string
		lbType           string
		lbCertificateARN string
		envID            string
	}
	updateReturns struct {
		result1 cloudformation.Stack
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 cloudformation.Stack
		result2 error
	}
	DeleteStub        func(stackName string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		stackName string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Infrastructure) Update(keyPairName string, azs []string, stackName string, boshAZ string, lbType string, lbCertificateARN string, envID string) (cloudformation.Stack, error) {
	var azsCopy []string
	if azs != nil {
		azsCopy = make([]string, len(azs))
		copy(azsCopy, azs)
	}
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		keyPairName      string
		azs              []string
		stackName        string
		boshAZ           string
		lbType           string
		lbCertificateARN string
		envID            string
	}{keyPairName, azsCopy, stackName, boshAZ, lbType, lbCertificateARN, envID})
	fake.recordInvocation("Update", []interface{}{keyPairName, azsCopy, stackName, boshAZ, lbType, lbCertificateARN, envID})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(keyPairName, azs, stackName, boshAZ, lbType, lbCertificateARN, envID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *Infrastructure) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *Infrastructure) UpdateArgsForCall(i int) (string, []string, string, string, string, string, string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].keyPairName, fake.updateArgsForCall[i].azs, fake.updateArgsForCall[i].stackName, fake.updateArgsForCall[i].boshAZ, fake.updateArgsForCall[i].lbType, fake.updateArgsForCall[i].lbCertificateARN, fake.updateArgsForCall[i].envID
}

func (fake *Infrastructure) UpdateReturns(result1 cloudformation.Stack, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *Infrastructure) UpdateReturnsOnCall(i int, result1 cloudformation.Stack, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 cloudformation.Stack
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *Infrastructure) Delete(stackName string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		stackName string
	}{stackName})
	fake.recordInvocation("Delete", []interface{}{stackName})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(stackName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *Infrastructure) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *Infrastructure) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].stackName
}

func (fake *Infrastructure) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *Infrastructure) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Infrastructure) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Infrastructure) recordInvocation(key string, args []interface{}) {
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