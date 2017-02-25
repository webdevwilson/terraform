// +build !core

//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//
package command

import (
	alicloudprovider "github.com/hashicorp/terraform/builtin/providers/alicloud"
	archiveprovider "github.com/hashicorp/terraform/builtin/providers/archive"
	artifactoryprovider "github.com/hashicorp/terraform/builtin/providers/artifactory"
	arukasprovider "github.com/hashicorp/terraform/builtin/providers/arukas"
	atlasprovider "github.com/hashicorp/terraform/builtin/providers/atlas"
	awsprovider "github.com/hashicorp/terraform/builtin/providers/aws"
	azureprovider "github.com/hashicorp/terraform/builtin/providers/azure"
	azurermprovider "github.com/hashicorp/terraform/builtin/providers/azurerm"
	bitbucketprovider "github.com/hashicorp/terraform/builtin/providers/bitbucket"
	chefprovider "github.com/hashicorp/terraform/builtin/providers/chef"
	clcprovider "github.com/hashicorp/terraform/builtin/providers/clc"
	cloudflareprovider "github.com/hashicorp/terraform/builtin/providers/cloudflare"
	cloudstackprovider "github.com/hashicorp/terraform/builtin/providers/cloudstack"
	cobblerprovider "github.com/hashicorp/terraform/builtin/providers/cobbler"
	consulprovider "github.com/hashicorp/terraform/builtin/providers/consul"
	datadogprovider "github.com/hashicorp/terraform/builtin/providers/datadog"
	digitaloceanprovider "github.com/hashicorp/terraform/builtin/providers/digitalocean"
	dmeprovider "github.com/hashicorp/terraform/builtin/providers/dme"
	dnsprovider "github.com/hashicorp/terraform/builtin/providers/dns"
	dnsimpleprovider "github.com/hashicorp/terraform/builtin/providers/dnsimple"
	dockerprovider "github.com/hashicorp/terraform/builtin/providers/docker"
	dynprovider "github.com/hashicorp/terraform/builtin/providers/dyn"
	externalprovider "github.com/hashicorp/terraform/builtin/providers/external"
	fastlyprovider "github.com/hashicorp/terraform/builtin/providers/fastly"
	githubprovider "github.com/hashicorp/terraform/builtin/providers/github"
	googleprovider "github.com/hashicorp/terraform/builtin/providers/google"
	grafanaprovider "github.com/hashicorp/terraform/builtin/providers/grafana"
	herokuprovider "github.com/hashicorp/terraform/builtin/providers/heroku"
	icinga2provider "github.com/hashicorp/terraform/builtin/providers/icinga2"
	ignitionprovider "github.com/hashicorp/terraform/builtin/providers/ignition"
	influxdbprovider "github.com/hashicorp/terraform/builtin/providers/influxdb"
	libratoprovider "github.com/hashicorp/terraform/builtin/providers/librato"
	logentriesprovider "github.com/hashicorp/terraform/builtin/providers/logentries"
	mailgunprovider "github.com/hashicorp/terraform/builtin/providers/mailgun"
	mysqlprovider "github.com/hashicorp/terraform/builtin/providers/mysql"
	newrelicprovider "github.com/hashicorp/terraform/builtin/providers/newrelic"
	nomadprovider "github.com/hashicorp/terraform/builtin/providers/nomad"
	ns1provider "github.com/hashicorp/terraform/builtin/providers/ns1"
	nullprovider "github.com/hashicorp/terraform/builtin/providers/null"
	openstackprovider "github.com/hashicorp/terraform/builtin/providers/openstack"
	opsgenieprovider "github.com/hashicorp/terraform/builtin/providers/opsgenie"
	packetprovider "github.com/hashicorp/terraform/builtin/providers/packet"
	pagerdutyprovider "github.com/hashicorp/terraform/builtin/providers/pagerduty"
	postgresqlprovider "github.com/hashicorp/terraform/builtin/providers/postgresql"
	powerdnsprovider "github.com/hashicorp/terraform/builtin/providers/powerdns"
	profitbricksprovider "github.com/hashicorp/terraform/builtin/providers/profitbricks"
	rabbitmqprovider "github.com/hashicorp/terraform/builtin/providers/rabbitmq"
	rancherprovider "github.com/hashicorp/terraform/builtin/providers/rancher"
	randomprovider "github.com/hashicorp/terraform/builtin/providers/random"
	rundeckprovider "github.com/hashicorp/terraform/builtin/providers/rundeck"
	scalewayprovider "github.com/hashicorp/terraform/builtin/providers/scaleway"
	softlayerprovider "github.com/hashicorp/terraform/builtin/providers/softlayer"
	spotinstprovider "github.com/hashicorp/terraform/builtin/providers/spotinst"
	statuscakeprovider "github.com/hashicorp/terraform/builtin/providers/statuscake"
	templateprovider "github.com/hashicorp/terraform/builtin/providers/template"
	terraformprovider "github.com/hashicorp/terraform/builtin/providers/terraform"
	testprovider "github.com/hashicorp/terraform/builtin/providers/test"
	tlsprovider "github.com/hashicorp/terraform/builtin/providers/tls"
	tritonprovider "github.com/hashicorp/terraform/builtin/providers/triton"
	ultradnsprovider "github.com/hashicorp/terraform/builtin/providers/ultradns"
	vaultprovider "github.com/hashicorp/terraform/builtin/providers/vault"
	vcdprovider "github.com/hashicorp/terraform/builtin/providers/vcd"
	vsphereprovider "github.com/hashicorp/terraform/builtin/providers/vsphere"
	fileprovisioner "github.com/hashicorp/terraform/builtin/provisioners/file"
	localexecprovisioner "github.com/hashicorp/terraform/builtin/provisioners/local-exec"
	remoteexecprovisioner "github.com/hashicorp/terraform/builtin/provisioners/remote-exec"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	// Legacy, will remove once it conforms with new structure
	chefprovisioner "github.com/hashicorp/terraform/builtin/provisioners/chef"
)

var InternalProviders = map[string]plugin.ProviderFunc{
	"alicloud":     alicloudprovider.Provider,
	"archive":      archiveprovider.Provider,
	"artifactory":  artifactoryprovider.Provider,
	"arukas":       arukasprovider.Provider,
	"atlas":        atlasprovider.Provider,
	"aws":          awsprovider.Provider,
	"azure":        azureprovider.Provider,
	"azurerm":      azurermprovider.Provider,
	"bitbucket":    bitbucketprovider.Provider,
	"chef":         chefprovider.Provider,
	"clc":          clcprovider.Provider,
	"cloudflare":   cloudflareprovider.Provider,
	"cloudstack":   cloudstackprovider.Provider,
	"cobbler":      cobblerprovider.Provider,
	"consul":       consulprovider.Provider,
	"datadog":      datadogprovider.Provider,
	"digitalocean": digitaloceanprovider.Provider,
	"dme":          dmeprovider.Provider,
	"dns":          dnsprovider.Provider,
	"dnsimple":     dnsimpleprovider.Provider,
	"docker":       dockerprovider.Provider,
	"dyn":          dynprovider.Provider,
	"external":     externalprovider.Provider,
	"fastly":       fastlyprovider.Provider,
	"github":       githubprovider.Provider,
	"google":       googleprovider.Provider,
	"grafana":      grafanaprovider.Provider,
	"heroku":       herokuprovider.Provider,
	"icinga2":      icinga2provider.Provider,
	"ignition":     ignitionprovider.Provider,
	"influxdb":     influxdbprovider.Provider,
	"librato":      libratoprovider.Provider,
	"logentries":   logentriesprovider.Provider,
	"mailgun":      mailgunprovider.Provider,
	"mysql":        mysqlprovider.Provider,
	"newrelic":     newrelicprovider.Provider,
	"nomad":        nomadprovider.Provider,
	"ns1":          ns1provider.Provider,
	"null":         nullprovider.Provider,
	"openstack":    openstackprovider.Provider,
	"opsgenie":     opsgenieprovider.Provider,
	"packet":       packetprovider.Provider,
	"pagerduty":    pagerdutyprovider.Provider,
	"postgresql":   postgresqlprovider.Provider,
	"powerdns":     powerdnsprovider.Provider,
	"profitbricks": profitbricksprovider.Provider,
	"rabbitmq":     rabbitmqprovider.Provider,
	"rancher":      rancherprovider.Provider,
	"random":       randomprovider.Provider,
	"rundeck":      rundeckprovider.Provider,
	"scaleway":     scalewayprovider.Provider,
	"softlayer":    softlayerprovider.Provider,
	"spotinst":     spotinstprovider.Provider,
	"statuscake":   statuscakeprovider.Provider,
	"template":     templateprovider.Provider,
	"terraform":    terraformprovider.Provider,
	"test":         testprovider.Provider,
	"tls":          tlsprovider.Provider,
	"triton":       tritonprovider.Provider,
	"ultradns":     ultradnsprovider.Provider,
	"vault":        vaultprovider.Provider,
	"vcd":          vcdprovider.Provider,
	"vsphere":      vsphereprovider.Provider,
}

var InternalProvisioners = map[string]plugin.ProvisionerFunc{
	"file":        fileprovisioner.Provisioner,
	"local-exec":  localexecprovisioner.Provisioner,
	"remote-exec": remoteexecprovisioner.Provisioner,
}

func init() {
	// Legacy provisioners that don't match our heuristics for auto-finding
	// built-in provisioners.
	InternalProvisioners["chef"] = func() terraform.ResourceProvisioner { return new(chefprovisioner.ResourceProvisioner) }
}
