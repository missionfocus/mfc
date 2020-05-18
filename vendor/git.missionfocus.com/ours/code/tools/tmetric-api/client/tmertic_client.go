// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"git.missionfocus.com/ours/code/tools/tmetric-api/client/account_members"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/accounts"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/activity"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/api_version"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/clients"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/demo_data"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/integrations"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/invoices"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/projects"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/recent_tasks"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_detailed"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_long_timers"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_projects"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_summary_activity"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_tasks"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/report_team"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/synchronized_versions"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/tags"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/tasks"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/time_entries"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/time_off_balances"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/time_off_policies"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/time_off_requests"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/timeline"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/timer"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/user_groups"
	"git.missionfocus.com/ours/code/tools/tmetric-api/client/user_profile"
)

// Default tmertic HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "app.tmetric.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new tmertic HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Tmertic {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new tmertic HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Tmertic {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new tmertic client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Tmertic {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Tmertic)
	cli.Transport = transport
	cli.AccountMembers = account_members.New(transport, formats)
	cli.Accounts = accounts.New(transport, formats)
	cli.Activity = activity.New(transport, formats)
	cli.APIVersion = api_version.New(transport, formats)
	cli.Clients = clients.New(transport, formats)
	cli.DemoData = demo_data.New(transport, formats)
	cli.Integrations = integrations.New(transport, formats)
	cli.Invoices = invoices.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.RecentTasks = recent_tasks.New(transport, formats)
	cli.ReportDetailed = report_detailed.New(transport, formats)
	cli.ReportLongTimers = report_long_timers.New(transport, formats)
	cli.ReportProjects = report_projects.New(transport, formats)
	cli.ReportSummaryActivity = report_summary_activity.New(transport, formats)
	cli.ReportTasks = report_tasks.New(transport, formats)
	cli.ReportTeam = report_team.New(transport, formats)
	cli.SynchronizedVersions = synchronized_versions.New(transport, formats)
	cli.Tags = tags.New(transport, formats)
	cli.Tasks = tasks.New(transport, formats)
	cli.TimeEntries = time_entries.New(transport, formats)
	cli.TimeOffBalances = time_off_balances.New(transport, formats)
	cli.TimeOffPolicies = time_off_policies.New(transport, formats)
	cli.TimeOffRequests = time_off_requests.New(transport, formats)
	cli.Timeline = timeline.New(transport, formats)
	cli.Timer = timer.New(transport, formats)
	cli.UserGroups = user_groups.New(transport, formats)
	cli.UserProfile = user_profile.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Tmertic is a client for tmertic
type Tmertic struct {
	AccountMembers account_members.ClientService

	Accounts accounts.ClientService

	Activity activity.ClientService

	APIVersion api_version.ClientService

	Clients clients.ClientService

	DemoData demo_data.ClientService

	Integrations integrations.ClientService

	Invoices invoices.ClientService

	Projects projects.ClientService

	RecentTasks recent_tasks.ClientService

	ReportDetailed report_detailed.ClientService

	ReportLongTimers report_long_timers.ClientService

	ReportProjects report_projects.ClientService

	ReportSummaryActivity report_summary_activity.ClientService

	ReportTasks report_tasks.ClientService

	ReportTeam report_team.ClientService

	SynchronizedVersions synchronized_versions.ClientService

	Tags tags.ClientService

	Tasks tasks.ClientService

	TimeEntries time_entries.ClientService

	TimeOffBalances time_off_balances.ClientService

	TimeOffPolicies time_off_policies.ClientService

	TimeOffRequests time_off_requests.ClientService

	Timeline timeline.ClientService

	Timer timer.ClientService

	UserGroups user_groups.ClientService

	UserProfile user_profile.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Tmertic) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccountMembers.SetTransport(transport)
	c.Accounts.SetTransport(transport)
	c.Activity.SetTransport(transport)
	c.APIVersion.SetTransport(transport)
	c.Clients.SetTransport(transport)
	c.DemoData.SetTransport(transport)
	c.Integrations.SetTransport(transport)
	c.Invoices.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.RecentTasks.SetTransport(transport)
	c.ReportDetailed.SetTransport(transport)
	c.ReportLongTimers.SetTransport(transport)
	c.ReportProjects.SetTransport(transport)
	c.ReportSummaryActivity.SetTransport(transport)
	c.ReportTasks.SetTransport(transport)
	c.ReportTeam.SetTransport(transport)
	c.SynchronizedVersions.SetTransport(transport)
	c.Tags.SetTransport(transport)
	c.Tasks.SetTransport(transport)
	c.TimeEntries.SetTransport(transport)
	c.TimeOffBalances.SetTransport(transport)
	c.TimeOffPolicies.SetTransport(transport)
	c.TimeOffRequests.SetTransport(transport)
	c.Timeline.SetTransport(transport)
	c.Timer.SetTransport(transport)
	c.UserGroups.SetTransport(transport)
	c.UserProfile.SetTransport(transport)
}
