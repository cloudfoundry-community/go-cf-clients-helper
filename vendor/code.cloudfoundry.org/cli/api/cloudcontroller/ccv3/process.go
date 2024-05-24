package ccv3

import (
	"encoding/json"
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/types"
)

type Process struct {
	GUID string
	Type string
	// Command is the process start command. Note: This value will be obfuscated when obtained from listing.
	Command                      types.FilteredString
	HealthCheckType              constant.HealthCheckType
	HealthCheckEndpoint          string
	HealthCheckInvocationTimeout int64
	HealthCheckTimeout           int64
	Instances                    types.NullInt
	MemoryInMB                   types.NullUint64
	DiskInMB                     types.NullUint64
	AppGUID                      string
}

func (p Process) MarshalJSON() ([]byte, error) {
	var ccProcess marshalProcess

	marshalCommand(p, &ccProcess)
	marshalInstances(p, &ccProcess)
	marshalMemory(p, &ccProcess)
	marshalDisk(p, &ccProcess)
	marshalHealthCheck(p, &ccProcess)

	return json.Marshal(ccProcess)
}

func (p *Process) UnmarshalJSON(data []byte) error {
	var ccProcess struct {
		Command       types.FilteredString    `json:"command"`
		DiskInMB      types.NullUint64        `json:"disk_in_mb"`
		GUID          string                  `json:"guid"`
		Instances     types.NullInt           `json:"instances"`
		MemoryInMB    types.NullUint64        `json:"memory_in_mb"`
		Type          string                  `json:"type"`
		Relationships resources.Relationships `json:"relationships"`

		HealthCheck struct {
			Type constant.HealthCheckType `json:"type"`
			Data struct {
				Endpoint          string `json:"endpoint"`
				InvocationTimeout int64  `json:"invocation_timeout"`
				Timeout           int64  `json:"timeout"`
			} `json:"data"`
		} `json:"health_check"`
	}

	err := cloudcontroller.DecodeJSON(data, &ccProcess)
	if err != nil {
		return err
	}

	p.Command = ccProcess.Command
	p.DiskInMB = ccProcess.DiskInMB
	p.GUID = ccProcess.GUID
	p.HealthCheckEndpoint = ccProcess.HealthCheck.Data.Endpoint
	p.HealthCheckInvocationTimeout = ccProcess.HealthCheck.Data.InvocationTimeout
	p.HealthCheckTimeout = ccProcess.HealthCheck.Data.Timeout
	p.HealthCheckType = ccProcess.HealthCheck.Type
	p.Instances = ccProcess.Instances
	p.MemoryInMB = ccProcess.MemoryInMB
	p.Type = ccProcess.Type
	p.AppGUID = ccProcess.Relationships[constant.RelationshipTypeApplication].GUID

	return nil
}

// CreateApplicationProcessScale updates process instances count, memory or disk
func (client *Client) CreateApplicationProcessScale(appGUID string, process Process) (Process, Warnings, error) {
	var responseBody Process

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostApplicationProcessActionScaleRequest,
		URIParams:    internal.Params{"app_guid": appGUID, "type": process.Type},
		RequestBody:  process,
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// GetApplicationProcessByType returns application process of specified type
func (client *Client) GetApplicationProcessByType(appGUID string, processType string) (Process, Warnings, error) {
	var responseBody Process

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetApplicationProcessRequest,
		URIParams:    internal.Params{"app_guid": appGUID, "type": processType},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// GetApplicationProcesses lists processes for a given application. **Note**:
// Due to security, the API obfuscates certain values such as `command`.
func (client *Client) GetApplicationProcesses(appGUID string) ([]Process, Warnings, error) {
	var resources []Process

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationProcessesRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		ResponseBody: Process{},
		AppendToList: func(item interface{}) error {
			resources = append(resources, item.(Process))
			return nil
		},
	})

	return resources, warnings, err
}

// GetNewApplicationProcesses gets processes for an application in the middle of a deployment.
// The app's processes will include a web process that will be removed when the deployment completes,
// so exclude that soon-to-be-removed process from the result.
func (client *Client) GetNewApplicationProcesses(appGUID string, deploymentGUID string) ([]Process, Warnings, error) {
	var allWarnings Warnings

	deployment, warnings, err := client.GetDeployment(deploymentGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return nil, allWarnings, err
	}

	allProcesses, warnings, err := client.GetApplicationProcesses(appGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return nil, allWarnings, err
	}

	var newWebProcessGUID string
	for _, process := range deployment.NewProcesses {
		if process.Type == constant.ProcessTypeWeb {
			newWebProcessGUID = process.GUID
		}
	}

	var processesList []Process
	for _, process := range allProcesses {
		if process.Type == constant.ProcessTypeWeb {
			if process.GUID == newWebProcessGUID {
				processesList = append(processesList, process)
			}
		} else {
			processesList = append(processesList, process)
		}
	}

	return processesList, allWarnings, nil
}

// GetProcess returns a process with the given guid
func (client *Client) GetProcess(processGUID string) (Process, Warnings, error) {
	var responseBody Process

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetProcessRequest,
		URIParams:    internal.Params{"process_guid": processGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client Client) GetProcesses(query ...Query) ([]Process, Warnings, error) {
	var resources []Process

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetProcessesRequest,
		Query:        query,
		ResponseBody: Process{},
		AppendToList: func(item interface{}) error {
			resources = append(resources, item.(Process))
			return nil
		},
	})

	return resources, warnings, err
}

// UpdateProcess updates the process's command or health check settings. GUID
// is always required; HealthCheckType is only required when updating health
// check settings.
func (client *Client) UpdateProcess(process Process) (Process, Warnings, error) {
	var responseBody Process

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.PatchProcessRequest,
		URIParams:   internal.Params{"process_guid": process.GUID},
		RequestBody: Process{
			Command:                      process.Command,
			HealthCheckType:              process.HealthCheckType,
			HealthCheckEndpoint:          process.HealthCheckEndpoint,
			HealthCheckTimeout:           process.HealthCheckTimeout,
			HealthCheckInvocationTimeout: process.HealthCheckInvocationTimeout,
		},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

type healthCheck struct {
	Type constant.HealthCheckType `json:"type,omitempty"`
	Data struct {
		Endpoint          interface{} `json:"endpoint,omitempty"`
		InvocationTimeout int64       `json:"invocation_timeout,omitempty"`
		Timeout           int64       `json:"timeout,omitempty"`
	} `json:"data"`
}

type marshalProcess struct {
	Command    interface{} `json:"command,omitempty"`
	Instances  json.Number `json:"instances,omitempty"`
	MemoryInMB json.Number `json:"memory_in_mb,omitempty"`
	DiskInMB   json.Number `json:"disk_in_mb,omitempty"`

	HealthCheck *healthCheck `json:"health_check,omitempty"`
}

func marshalCommand(p Process, ccProcess *marshalProcess) {
	if p.Command.IsSet {
		ccProcess.Command = &p.Command
	}
}

func marshalDisk(p Process, ccProcess *marshalProcess) {
	if p.DiskInMB.IsSet {
		ccProcess.DiskInMB = json.Number(fmt.Sprint(p.DiskInMB.Value))
	}
}

func marshalHealthCheck(p Process, ccProcess *marshalProcess) {
	if p.HealthCheckType != "" || p.HealthCheckEndpoint != "" || p.HealthCheckInvocationTimeout != 0 || p.HealthCheckTimeout != 0 {
		ccProcess.HealthCheck = new(healthCheck)
		ccProcess.HealthCheck.Type = p.HealthCheckType
		ccProcess.HealthCheck.Data.InvocationTimeout = p.HealthCheckInvocationTimeout
		ccProcess.HealthCheck.Data.Timeout = p.HealthCheckTimeout
		if p.HealthCheckEndpoint != "" {
			ccProcess.HealthCheck.Data.Endpoint = p.HealthCheckEndpoint
		}
	}
}

func marshalInstances(p Process, ccProcess *marshalProcess) {
	if p.Instances.IsSet {
		ccProcess.Instances = json.Number(fmt.Sprint(p.Instances.Value))
	}
}

func marshalMemory(p Process, ccProcess *marshalProcess) {
	if p.MemoryInMB.IsSet {
		ccProcess.MemoryInMB = json.Number(fmt.Sprint(p.MemoryInMB.Value))
	}
}
