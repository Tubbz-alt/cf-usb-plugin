// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/go-openapi/runtime"
	"github.com/SUSE/cf-usb-plugin/lib"
	operations "github.com/SUSE/cf-usb-plugin/lib/client/operations"
	"github.com/SUSE/cf-usb-plugin/lib/models"
)

type FakeUsbClientInterface struct {
	RegisterDriverEndpointStub        func(params *operations.RegisterDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.RegisterDriverEndpointCreated, error)
	registerDriverEndpointMutex       sync.RWMutex
	registerDriverEndpointArgsForCall []struct {
		params   *operations.RegisterDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}
	registerDriverEndpointReturns struct {
		result1 *operations.RegisterDriverEndpointCreated
		result2 error
	}
	UnregisterDriverEndpointStub        func(params *operations.UnregisterDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UnregisterDriverInstanceNoContent, error)
	unregisterDriverEndpointMutex       sync.RWMutex
	unregisterDriverEndpointArgsForCall []struct {
		params   *operations.UnregisterDriverInstanceParams
		authInfo runtime.ClientAuthInfoWriter
	}
	unregisterDriverEndpointReturns struct {
		result1 *operations.UnregisterDriverInstanceNoContent
		result2 error
	}
	GetDriverEndpointStub        func(params *operations.GetDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointOK, error)
	getDriverEndpointMutex       sync.RWMutex
	getDriverEndpointArgsForCall []struct {
		params   *operations.GetDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}
	getDriverEndpointReturns struct {
		result1 *operations.GetDriverEndpointOK
		result2 error
	}
	GetDriverEndpointByNameStub        func(instanceName string, authInfo runtime.ClientAuthInfoWriter) (*models.DriverEndpoint, error)
	getDriverEndpointByNameMutex       sync.RWMutex
	getDriverEndpointByNameArgsForCall []struct {
		instanceName string
		authInfo     runtime.ClientAuthInfoWriter
	}
	getDriverEndpointByNameReturns struct {
		result1 *models.DriverEndpoint
		result2 error
	}
	GetDriverEndpointsStub        func(params *operations.GetDriverEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointsOK, error)
	getDriverEndpointsMutex       sync.RWMutex
	getDriverEndpointsArgsForCall []struct {
		params   *operations.GetDriverEndpointsParams
		authInfo runtime.ClientAuthInfoWriter
	}
	getDriverEndpointsReturns struct {
		result1 *operations.GetDriverEndpointsOK
		result2 error
	}
	GetInfoStub        func(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct {
		params   *operations.GetInfoParams
		authInfo runtime.ClientAuthInfoWriter
	}
	getInfoReturns struct {
		result1 *operations.GetInfoOK
		result2 error
	}
	PingEndpointStub        func(params *operations.PingDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverEndpointOK, error)
	pingEndpointMutex       sync.RWMutex
	pingEndpointArgsForCall []struct {
		params   *operations.PingDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}
	pingEndpointReturns struct {
		result1 *operations.PingDriverEndpointOK
		result2 error
	}
	UpdateCatalogStub        func(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error)
	updateCatalogMutex       sync.RWMutex
	updateCatalogArgsForCall []struct {
		params   *operations.UpdateCatalogParams
		authInfo runtime.ClientAuthInfoWriter
	}
	updateCatalogReturns struct {
		result1 *operations.UpdateCatalogOK
		result2 error
	}
	UpdateDriverEndpointStub        func(params *operations.UpdateDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverEndpointOK, error)
	updateDriverEndpointMutex       sync.RWMutex
	updateDriverEndpointArgsForCall []struct {
		params   *operations.UpdateDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}
	updateDriverEndpointReturns struct {
		result1 *operations.UpdateDriverEndpointOK
		result2 error
	}
	SetTransportStub        func(transport runtime.ClientTransport)
	setTransportMutex       sync.RWMutex
	setTransportArgsForCall []struct {
		transport runtime.ClientTransport
	}
}

func (fake *FakeUsbClientInterface) RegisterDriverEndpoint(params *operations.RegisterDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.RegisterDriverEndpointCreated, error) {
	fake.registerDriverEndpointMutex.Lock()
	fake.registerDriverEndpointArgsForCall = append(fake.registerDriverEndpointArgsForCall, struct {
		params   *operations.RegisterDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.registerDriverEndpointMutex.Unlock()
	if fake.RegisterDriverEndpointStub != nil {
		return fake.RegisterDriverEndpointStub(params, authInfo)
	} else {
		return fake.registerDriverEndpointReturns.result1, fake.registerDriverEndpointReturns.result2
	}
}

func (fake *FakeUsbClientInterface) RegisterDriverEndpointCallCount() int {
	fake.registerDriverEndpointMutex.RLock()
	defer fake.registerDriverEndpointMutex.RUnlock()
	return len(fake.registerDriverEndpointArgsForCall)
}

func (fake *FakeUsbClientInterface) RegisterDriverEndpointArgsForCall(i int) (*operations.RegisterDriverEndpointParams, runtime.ClientAuthInfoWriter) {
	fake.registerDriverEndpointMutex.RLock()
	defer fake.registerDriverEndpointMutex.RUnlock()
	return fake.registerDriverEndpointArgsForCall[i].params, fake.registerDriverEndpointArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) RegisterDriverEndpointReturns(result1 *operations.RegisterDriverEndpointCreated, result2 error) {
	fake.RegisterDriverEndpointStub = nil
	fake.registerDriverEndpointReturns = struct {
		result1 *operations.RegisterDriverEndpointCreated
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) UnregisterDriverEndpoint(params *operations.UnregisterDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UnregisterDriverInstanceNoContent, error) {
	fake.unregisterDriverEndpointMutex.Lock()
	fake.unregisterDriverEndpointArgsForCall = append(fake.unregisterDriverEndpointArgsForCall, struct {
		params   *operations.UnregisterDriverInstanceParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.unregisterDriverEndpointMutex.Unlock()
	if fake.UnregisterDriverEndpointStub != nil {
		return fake.UnregisterDriverEndpointStub(params, authInfo)
	} else {
		return fake.unregisterDriverEndpointReturns.result1, fake.unregisterDriverEndpointReturns.result2
	}
}

func (fake *FakeUsbClientInterface) UnregisterDriverEndpointCallCount() int {
	fake.unregisterDriverEndpointMutex.RLock()
	defer fake.unregisterDriverEndpointMutex.RUnlock()
	return len(fake.unregisterDriverEndpointArgsForCall)
}

func (fake *FakeUsbClientInterface) UnregisterDriverEndpointArgsForCall(i int) (*operations.UnregisterDriverInstanceParams, runtime.ClientAuthInfoWriter) {
	fake.unregisterDriverEndpointMutex.RLock()
	defer fake.unregisterDriverEndpointMutex.RUnlock()
	return fake.unregisterDriverEndpointArgsForCall[i].params, fake.unregisterDriverEndpointArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) UnregisterDriverEndpointReturns(result1 *operations.UnregisterDriverInstanceNoContent, result2 error) {
	fake.UnregisterDriverEndpointStub = nil
	fake.unregisterDriverEndpointReturns = struct {
		result1 *operations.UnregisterDriverInstanceNoContent
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) GetDriverEndpoint(params *operations.GetDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointOK, error) {
	fake.getDriverEndpointMutex.Lock()
	fake.getDriverEndpointArgsForCall = append(fake.getDriverEndpointArgsForCall, struct {
		params   *operations.GetDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.getDriverEndpointMutex.Unlock()
	if fake.GetDriverEndpointStub != nil {
		return fake.GetDriverEndpointStub(params, authInfo)
	} else {
		return fake.getDriverEndpointReturns.result1, fake.getDriverEndpointReturns.result2
	}
}

func (fake *FakeUsbClientInterface) GetDriverEndpointCallCount() int {
	fake.getDriverEndpointMutex.RLock()
	defer fake.getDriverEndpointMutex.RUnlock()
	return len(fake.getDriverEndpointArgsForCall)
}

func (fake *FakeUsbClientInterface) GetDriverEndpointArgsForCall(i int) (*operations.GetDriverEndpointParams, runtime.ClientAuthInfoWriter) {
	fake.getDriverEndpointMutex.RLock()
	defer fake.getDriverEndpointMutex.RUnlock()
	return fake.getDriverEndpointArgsForCall[i].params, fake.getDriverEndpointArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) GetDriverEndpointReturns(result1 *operations.GetDriverEndpointOK, result2 error) {
	fake.GetDriverEndpointStub = nil
	fake.getDriverEndpointReturns = struct {
		result1 *operations.GetDriverEndpointOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) GetDriverEndpointByName(instanceName string, authInfo runtime.ClientAuthInfoWriter) (*models.DriverEndpoint, error) {
	fake.getDriverEndpointByNameMutex.Lock()
	fake.getDriverEndpointByNameArgsForCall = append(fake.getDriverEndpointByNameArgsForCall, struct {
		instanceName string
		authInfo     runtime.ClientAuthInfoWriter
	}{instanceName, authInfo})
	fake.getDriverEndpointByNameMutex.Unlock()
	if fake.GetDriverEndpointByNameStub != nil {
		return fake.GetDriverEndpointByNameStub(instanceName, authInfo)
	} else {
		return fake.getDriverEndpointByNameReturns.result1, fake.getDriverEndpointByNameReturns.result2
	}
}

func (fake *FakeUsbClientInterface) GetDriverEndpointByNameCallCount() int {
	fake.getDriverEndpointByNameMutex.RLock()
	defer fake.getDriverEndpointByNameMutex.RUnlock()
	return len(fake.getDriverEndpointByNameArgsForCall)
}

func (fake *FakeUsbClientInterface) GetDriverEndpointByNameArgsForCall(i int) (string, runtime.ClientAuthInfoWriter) {
	fake.getDriverEndpointByNameMutex.RLock()
	defer fake.getDriverEndpointByNameMutex.RUnlock()
	return fake.getDriverEndpointByNameArgsForCall[i].instanceName, fake.getDriverEndpointByNameArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) GetDriverEndpointByNameReturns(result1 *models.DriverEndpoint, result2 error) {
	fake.GetDriverEndpointByNameStub = nil
	fake.getDriverEndpointByNameReturns = struct {
		result1 *models.DriverEndpoint
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) GetDriverEndpoints(params *operations.GetDriverEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointsOK, error) {
	fake.getDriverEndpointsMutex.Lock()
	fake.getDriverEndpointsArgsForCall = append(fake.getDriverEndpointsArgsForCall, struct {
		params   *operations.GetDriverEndpointsParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.getDriverEndpointsMutex.Unlock()
	if fake.GetDriverEndpointsStub != nil {
		return fake.GetDriverEndpointsStub(params, authInfo)
	} else {
		return fake.getDriverEndpointsReturns.result1, fake.getDriverEndpointsReturns.result2
	}
}

func (fake *FakeUsbClientInterface) GetDriverEndpointsCallCount() int {
	fake.getDriverEndpointsMutex.RLock()
	defer fake.getDriverEndpointsMutex.RUnlock()
	return len(fake.getDriverEndpointsArgsForCall)
}

func (fake *FakeUsbClientInterface) GetDriverEndpointsArgsForCall(i int) (*operations.GetDriverEndpointsParams, runtime.ClientAuthInfoWriter) {
	fake.getDriverEndpointsMutex.RLock()
	defer fake.getDriverEndpointsMutex.RUnlock()
	return fake.getDriverEndpointsArgsForCall[i].params, fake.getDriverEndpointsArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) GetDriverEndpointsReturns(result1 *operations.GetDriverEndpointsOK, result2 error) {
	fake.GetDriverEndpointsStub = nil
	fake.getDriverEndpointsReturns = struct {
		result1 *operations.GetDriverEndpointsOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error) {
	fake.getInfoMutex.Lock()
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct {
		params   *operations.GetInfoParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub(params, authInfo)
	} else {
		return fake.getInfoReturns.result1, fake.getInfoReturns.result2
	}
}

func (fake *FakeUsbClientInterface) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeUsbClientInterface) GetInfoArgsForCall(i int) (*operations.GetInfoParams, runtime.ClientAuthInfoWriter) {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return fake.getInfoArgsForCall[i].params, fake.getInfoArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) GetInfoReturns(result1 *operations.GetInfoOK, result2 error) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 *operations.GetInfoOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) PingEndpoint(params *operations.PingDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverEndpointOK, error) {
	fake.pingEndpointMutex.Lock()
	fake.pingEndpointArgsForCall = append(fake.pingEndpointArgsForCall, struct {
		params   *operations.PingDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.pingEndpointMutex.Unlock()
	if fake.PingEndpointStub != nil {
		return fake.PingEndpointStub(params, authInfo)
	} else {
		return fake.pingEndpointReturns.result1, fake.pingEndpointReturns.result2
	}
}

func (fake *FakeUsbClientInterface) PingEndpointCallCount() int {
	fake.pingEndpointMutex.RLock()
	defer fake.pingEndpointMutex.RUnlock()
	return len(fake.pingEndpointArgsForCall)
}

func (fake *FakeUsbClientInterface) PingEndpointArgsForCall(i int) (*operations.PingDriverEndpointParams, runtime.ClientAuthInfoWriter) {
	fake.pingEndpointMutex.RLock()
	defer fake.pingEndpointMutex.RUnlock()
	return fake.pingEndpointArgsForCall[i].params, fake.pingEndpointArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) PingEndpointReturns(result1 *operations.PingDriverEndpointOK, result2 error) {
	fake.PingEndpointStub = nil
	fake.pingEndpointReturns = struct {
		result1 *operations.PingDriverEndpointOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error) {
	fake.updateCatalogMutex.Lock()
	fake.updateCatalogArgsForCall = append(fake.updateCatalogArgsForCall, struct {
		params   *operations.UpdateCatalogParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.updateCatalogMutex.Unlock()
	if fake.UpdateCatalogStub != nil {
		return fake.UpdateCatalogStub(params, authInfo)
	} else {
		return fake.updateCatalogReturns.result1, fake.updateCatalogReturns.result2
	}
}

func (fake *FakeUsbClientInterface) UpdateCatalogCallCount() int {
	fake.updateCatalogMutex.RLock()
	defer fake.updateCatalogMutex.RUnlock()
	return len(fake.updateCatalogArgsForCall)
}

func (fake *FakeUsbClientInterface) UpdateCatalogArgsForCall(i int) (*operations.UpdateCatalogParams, runtime.ClientAuthInfoWriter) {
	fake.updateCatalogMutex.RLock()
	defer fake.updateCatalogMutex.RUnlock()
	return fake.updateCatalogArgsForCall[i].params, fake.updateCatalogArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) UpdateCatalogReturns(result1 *operations.UpdateCatalogOK, result2 error) {
	fake.UpdateCatalogStub = nil
	fake.updateCatalogReturns = struct {
		result1 *operations.UpdateCatalogOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) UpdateDriverEndpoint(params *operations.UpdateDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverEndpointOK, error) {
	fake.updateDriverEndpointMutex.Lock()
	fake.updateDriverEndpointArgsForCall = append(fake.updateDriverEndpointArgsForCall, struct {
		params   *operations.UpdateDriverEndpointParams
		authInfo runtime.ClientAuthInfoWriter
	}{params, authInfo})
	fake.updateDriverEndpointMutex.Unlock()
	if fake.UpdateDriverEndpointStub != nil {
		return fake.UpdateDriverEndpointStub(params, authInfo)
	} else {
		return fake.updateDriverEndpointReturns.result1, fake.updateDriverEndpointReturns.result2
	}
}

func (fake *FakeUsbClientInterface) UpdateDriverEndpointCallCount() int {
	fake.updateDriverEndpointMutex.RLock()
	defer fake.updateDriverEndpointMutex.RUnlock()
	return len(fake.updateDriverEndpointArgsForCall)
}

func (fake *FakeUsbClientInterface) UpdateDriverEndpointArgsForCall(i int) (*operations.UpdateDriverEndpointParams, runtime.ClientAuthInfoWriter) {
	fake.updateDriverEndpointMutex.RLock()
	defer fake.updateDriverEndpointMutex.RUnlock()
	return fake.updateDriverEndpointArgsForCall[i].params, fake.updateDriverEndpointArgsForCall[i].authInfo
}

func (fake *FakeUsbClientInterface) UpdateDriverEndpointReturns(result1 *operations.UpdateDriverEndpointOK, result2 error) {
	fake.UpdateDriverEndpointStub = nil
	fake.updateDriverEndpointReturns = struct {
		result1 *operations.UpdateDriverEndpointOK
		result2 error
	}{result1, result2}
}

func (fake *FakeUsbClientInterface) SetTransport(transport runtime.ClientTransport) {
	fake.setTransportMutex.Lock()
	fake.setTransportArgsForCall = append(fake.setTransportArgsForCall, struct {
		transport runtime.ClientTransport
	}{transport})
	fake.setTransportMutex.Unlock()
	if fake.SetTransportStub != nil {
		fake.SetTransportStub(transport)
	}
}

func (fake *FakeUsbClientInterface) SetTransportCallCount() int {
	fake.setTransportMutex.RLock()
	defer fake.setTransportMutex.RUnlock()
	return len(fake.setTransportArgsForCall)
}

func (fake *FakeUsbClientInterface) SetTransportArgsForCall(i int) runtime.ClientTransport {
	fake.setTransportMutex.RLock()
	defer fake.setTransportMutex.RUnlock()
	return fake.setTransportArgsForCall[i].transport
}

var _ lib.UsbClientInterface = new(FakeUsbClientInterface)
