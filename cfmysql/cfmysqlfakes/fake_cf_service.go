// This file was generated by counterfeiter
package cfmysqlfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	modelsplugin "code.cloudfoundry.org/cli/plugin/models"
	"github.com/andreasf/cf-mysql-plugin/cfmysql"
)

type FakeCfService struct {
	GetMysqlServicesStub        func(cliConnection plugin.CliConnection) ([]cfmysql.MysqlService, error)
	getMysqlServicesMutex       sync.RWMutex
	getMysqlServicesArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getMysqlServicesReturns struct {
		result1 []cfmysql.MysqlService
		result2 error
	}
	GetStartedAppsStub        func(cliConnection plugin.CliConnection) ([]modelsplugin.GetAppsModel, error)
	getStartedAppsMutex       sync.RWMutex
	getStartedAppsArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getStartedAppsReturns struct {
		result1 []modelsplugin.GetAppsModel
		result2 error
	}
	OpenSshTunnelStub        func(cliConnection plugin.CliConnection, toService cfmysql.MysqlService, throughApp string, localPort int)
	openSshTunnelMutex       sync.RWMutex
	openSshTunnelArgsForCall []struct {
		cliConnection plugin.CliConnection
		toService     cfmysql.MysqlService
		throughApp    string
		localPort     int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfService) GetMysqlServices(cliConnection plugin.CliConnection) ([]cfmysql.MysqlService, error) {
	fake.getMysqlServicesMutex.Lock()
	fake.getMysqlServicesArgsForCall = append(fake.getMysqlServicesArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetMysqlServices", []interface{}{cliConnection})
	fake.getMysqlServicesMutex.Unlock()
	if fake.GetMysqlServicesStub != nil {
		return fake.GetMysqlServicesStub(cliConnection)
	} else {
		return fake.getMysqlServicesReturns.result1, fake.getMysqlServicesReturns.result2
	}
}

func (fake *FakeCfService) GetMysqlServicesCallCount() int {
	fake.getMysqlServicesMutex.RLock()
	defer fake.getMysqlServicesMutex.RUnlock()
	return len(fake.getMysqlServicesArgsForCall)
}

func (fake *FakeCfService) GetMysqlServicesArgsForCall(i int) plugin.CliConnection {
	fake.getMysqlServicesMutex.RLock()
	defer fake.getMysqlServicesMutex.RUnlock()
	return fake.getMysqlServicesArgsForCall[i].cliConnection
}

func (fake *FakeCfService) GetMysqlServicesReturns(result1 []cfmysql.MysqlService, result2 error) {
	fake.GetMysqlServicesStub = nil
	fake.getMysqlServicesReturns = struct {
		result1 []cfmysql.MysqlService
		result2 error
	}{result1, result2}
}

func (fake *FakeCfService) GetStartedApps(cliConnection plugin.CliConnection) ([]modelsplugin.GetAppsModel, error) {
	fake.getStartedAppsMutex.Lock()
	fake.getStartedAppsArgsForCall = append(fake.getStartedAppsArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetStartedApps", []interface{}{cliConnection})
	fake.getStartedAppsMutex.Unlock()
	if fake.GetStartedAppsStub != nil {
		return fake.GetStartedAppsStub(cliConnection)
	} else {
		return fake.getStartedAppsReturns.result1, fake.getStartedAppsReturns.result2
	}
}

func (fake *FakeCfService) GetStartedAppsCallCount() int {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return len(fake.getStartedAppsArgsForCall)
}

func (fake *FakeCfService) GetStartedAppsArgsForCall(i int) plugin.CliConnection {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return fake.getStartedAppsArgsForCall[i].cliConnection
}

func (fake *FakeCfService) GetStartedAppsReturns(result1 []modelsplugin.GetAppsModel, result2 error) {
	fake.GetStartedAppsStub = nil
	fake.getStartedAppsReturns = struct {
		result1 []modelsplugin.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCfService) OpenSshTunnel(cliConnection plugin.CliConnection, toService cfmysql.MysqlService, throughApp string, localPort int) {
	fake.openSshTunnelMutex.Lock()
	fake.openSshTunnelArgsForCall = append(fake.openSshTunnelArgsForCall, struct {
		cliConnection plugin.CliConnection
		toService     cfmysql.MysqlService
		throughApp    string
		localPort     int
	}{cliConnection, toService, throughApp, localPort})
	fake.recordInvocation("OpenSshTunnel", []interface{}{cliConnection, toService, throughApp, localPort})
	fake.openSshTunnelMutex.Unlock()
	if fake.OpenSshTunnelStub != nil {
		fake.OpenSshTunnelStub(cliConnection, toService, throughApp, localPort)
	}
}

func (fake *FakeCfService) OpenSshTunnelCallCount() int {
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return len(fake.openSshTunnelArgsForCall)
}

func (fake *FakeCfService) OpenSshTunnelArgsForCall(i int) (plugin.CliConnection, cfmysql.MysqlService, string, int) {
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return fake.openSshTunnelArgsForCall[i].cliConnection, fake.openSshTunnelArgsForCall[i].toService, fake.openSshTunnelArgsForCall[i].throughApp, fake.openSshTunnelArgsForCall[i].localPort
}

func (fake *FakeCfService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMysqlServicesMutex.RLock()
	defer fake.getMysqlServicesMutex.RUnlock()
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	fake.openSshTunnelMutex.RLock()
	defer fake.openSshTunnelMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCfService) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.CfService = new(FakeCfService)
