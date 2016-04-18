// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeGardenWorkerDB struct {
	CreateContainerStub        func(db.Container, time.Duration) (db.SavedContainer, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 db.Container
		arg2 time.Duration
	}
	createContainerReturns struct {
		result1 db.SavedContainer
		result2 error
	}
	UpdateExpiresAtOnContainerStub        func(handle string, ttl time.Duration) error
	updateExpiresAtOnContainerMutex       sync.RWMutex
	updateExpiresAtOnContainerArgsForCall []struct {
		handle string
		ttl    time.Duration
	}
	updateExpiresAtOnContainerReturns struct {
		result1 error
	}
	InsertVolumeStub        func(db.Volume) error
	insertVolumeMutex       sync.RWMutex
	insertVolumeArgsForCall []struct {
		arg1 db.Volume
	}
	insertVolumeReturns struct {
		result1 error
	}
	GetVolumesByIdentifierStub        func(db.VolumeIdentifier) ([]db.SavedVolume, error)
	getVolumesByIdentifierMutex       sync.RWMutex
	getVolumesByIdentifierArgsForCall []struct {
		arg1 db.VolumeIdentifier
	}
	getVolumesByIdentifierReturns struct {
		result1 []db.SavedVolume
		result2 error
	}
}

func (fake *FakeGardenWorkerDB) CreateContainer(arg1 db.Container, arg2 time.Duration) (db.SavedContainer, error) {
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 db.Container
		arg2 time.Duration
	}{arg1, arg2})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(arg1, arg2)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) CreateContainerArgsForCall(i int) (db.Container, time.Duration) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].arg1, fake.createContainerArgsForCall[i].arg2
}

func (fake *FakeGardenWorkerDB) CreateContainerReturns(result1 db.SavedContainer, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 db.SavedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainer(handle string, ttl time.Duration) error {
	fake.updateExpiresAtOnContainerMutex.Lock()
	fake.updateExpiresAtOnContainerArgsForCall = append(fake.updateExpiresAtOnContainerArgsForCall, struct {
		handle string
		ttl    time.Duration
	}{handle, ttl})
	fake.updateExpiresAtOnContainerMutex.Unlock()
	if fake.UpdateExpiresAtOnContainerStub != nil {
		return fake.UpdateExpiresAtOnContainerStub(handle, ttl)
	} else {
		return fake.updateExpiresAtOnContainerReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerCallCount() int {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return len(fake.updateExpiresAtOnContainerArgsForCall)
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerArgsForCall(i int) (string, time.Duration) {
	fake.updateExpiresAtOnContainerMutex.RLock()
	defer fake.updateExpiresAtOnContainerMutex.RUnlock()
	return fake.updateExpiresAtOnContainerArgsForCall[i].handle, fake.updateExpiresAtOnContainerArgsForCall[i].ttl
}

func (fake *FakeGardenWorkerDB) UpdateExpiresAtOnContainerReturns(result1 error) {
	fake.UpdateExpiresAtOnContainerStub = nil
	fake.updateExpiresAtOnContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) InsertVolume(arg1 db.Volume) error {
	fake.insertVolumeMutex.Lock()
	fake.insertVolumeArgsForCall = append(fake.insertVolumeArgsForCall, struct {
		arg1 db.Volume
	}{arg1})
	fake.insertVolumeMutex.Unlock()
	if fake.InsertVolumeStub != nil {
		return fake.InsertVolumeStub(arg1)
	} else {
		return fake.insertVolumeReturns.result1
	}
}

func (fake *FakeGardenWorkerDB) InsertVolumeCallCount() int {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return len(fake.insertVolumeArgsForCall)
}

func (fake *FakeGardenWorkerDB) InsertVolumeArgsForCall(i int) db.Volume {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return fake.insertVolumeArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) InsertVolumeReturns(result1 error) {
	fake.InsertVolumeStub = nil
	fake.insertVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifier(arg1 db.VolumeIdentifier) ([]db.SavedVolume, error) {
	fake.getVolumesByIdentifierMutex.Lock()
	fake.getVolumesByIdentifierArgsForCall = append(fake.getVolumesByIdentifierArgsForCall, struct {
		arg1 db.VolumeIdentifier
	}{arg1})
	fake.getVolumesByIdentifierMutex.Unlock()
	if fake.GetVolumesByIdentifierStub != nil {
		return fake.GetVolumesByIdentifierStub(arg1)
	} else {
		return fake.getVolumesByIdentifierReturns.result1, fake.getVolumesByIdentifierReturns.result2
	}
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierCallCount() int {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return len(fake.getVolumesByIdentifierArgsForCall)
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierArgsForCall(i int) db.VolumeIdentifier {
	fake.getVolumesByIdentifierMutex.RLock()
	defer fake.getVolumesByIdentifierMutex.RUnlock()
	return fake.getVolumesByIdentifierArgsForCall[i].arg1
}

func (fake *FakeGardenWorkerDB) GetVolumesByIdentifierReturns(result1 []db.SavedVolume, result2 error) {
	fake.GetVolumesByIdentifierStub = nil
	fake.getVolumesByIdentifierReturns = struct {
		result1 []db.SavedVolume
		result2 error
	}{result1, result2}
}

var _ worker.GardenWorkerDB = new(FakeGardenWorkerDB)
