package azure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/golang/glog"
)

var environments = map[string]Environment{
	"AZURECHINACLOUD":        ChinaCloud,
	"AZUREGERMANCLOUD":       GermanCloud,
	"AZUREPUBLICCLOUD":       PublicCloud,
	"AZUREUSGOVERNMENTCLOUD": USGovernmentCloud,
	"AZURESTACKCLOUD":        AzurestackCloud,
}

// Environment represents a set of endpoints for each of Azure's Clouds.
type Environment struct {
	Name                         string `json:"name"`
	ManagementPortalURL          string `json:"managementPortalURL"`
	PublishSettingsURL           string `json:"publishSettingsURL"`
	ServiceManagementEndpoint    string `json:"serviceManagementEndpoint"`
	ResourceManagerEndpoint      string `json:"resourceManagerEndpoint"`
	ActiveDirectoryEndpoint      string `json:"activeDirectoryEndpoint"`
	GalleryEndpoint              string `json:"galleryEndpoint"`
	KeyVaultEndpoint             string `json:"keyVaultEndpoint"`
	GraphEndpoint                string `json:"graphEndpoint"`
	StorageEndpointSuffix        string `json:"storageEndpointSuffix"`
	SQLDatabaseDNSSuffix         string `json:"sqlDatabaseDNSSuffix"`
	TrafficManagerDNSSuffix      string `json:"trafficManagerDNSSuffix"`
	KeyVaultDNSSuffix            string `json:"keyVaultDNSSuffix"`
	ServiceBusEndpointSuffix     string `json:"serviceBusEndpointSuffix"`
	ServiceManagementVMDNSSuffix string `json:"serviceManagementVMDNSSuffix"`
	ResourceManagerVMDNSSuffix   string `json:"resourceManagerVMDNSSuffix"`
	ContainerRegistryDNSSuffix   string `json:"containerRegistryDNSSuffix"`
}

var (
	// PublicCloud is the default public Azure cloud environment
	PublicCloud = Environment{
		Name:                         "AzurePublicCloud",
		ManagementPortalURL:          "https://manage.windowsazure.com/",
		PublishSettingsURL:           "https://manage.windowsazure.com/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.windows.net/",
		ResourceManagerEndpoint:      "https://management.azure.com/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.com/",
		GalleryEndpoint:              "https://gallery.azure.com/",
		KeyVaultEndpoint:             "https://vault.azure.net/",
		GraphEndpoint:                "https://graph.windows.net/",
		StorageEndpointSuffix:        "core.windows.net",
		SQLDatabaseDNSSuffix:         "database.windows.net",
		TrafficManagerDNSSuffix:      "trafficmanager.net",
		KeyVaultDNSSuffix:            "vault.azure.net",
		ServiceBusEndpointSuffix:     "servicebus.azure.com",
		ServiceManagementVMDNSSuffix: "cloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.com",
		ContainerRegistryDNSSuffix:   "azurecr.io",
	}

	// USGovernmentCloud is the cloud environment for the US Government
	USGovernmentCloud = Environment{
		Name:                         "AzureUSGovernmentCloud",
		ManagementPortalURL:          "https://manage.windowsazure.us/",
		PublishSettingsURL:           "https://manage.windowsazure.us/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:      "https://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.com/",
		GalleryEndpoint:              "https://gallery.usgovcloudapi.net/",
		KeyVaultEndpoint:             "https://vault.usgovcloudapi.net/",
		GraphEndpoint:                "https://graph.usgovcloudapi.net/",
		StorageEndpointSuffix:        "core.usgovcloudapi.net",
		SQLDatabaseDNSSuffix:         "database.usgovcloudapi.net",
		TrafficManagerDNSSuffix:      "usgovtrafficmanager.net",
		KeyVaultDNSSuffix:            "vault.usgovcloudapi.net",
		ServiceBusEndpointSuffix:     "servicebus.usgovcloudapi.net",
		ServiceManagementVMDNSSuffix: "usgovcloudapp.net",
		ResourceManagerVMDNSSuffix:   "cloudapp.windowsazure.us",
		ContainerRegistryDNSSuffix:   "azurecr.io",
	}

	// ChinaCloud is the cloud environment operated in China
	ChinaCloud = Environment{
		Name:                         "AzureChinaCloud",
		ManagementPortalURL:          "https://manage.chinacloudapi.com/",
		PublishSettingsURL:           "https://manage.chinacloudapi.com/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:      "https://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:      "https://login.chinacloudapi.cn/",
		GalleryEndpoint:              "https://gallery.chinacloudapi.cn/",
		KeyVaultEndpoint:             "https://vault.azure.cn/",
		GraphEndpoint:                "https://graph.chinacloudapi.cn/",
		StorageEndpointSuffix:        "core.chinacloudapi.cn",
		SQLDatabaseDNSSuffix:         "database.chinacloudapi.cn",
		TrafficManagerDNSSuffix:      "trafficmanager.cn",
		KeyVaultDNSSuffix:            "vault.azure.cn",
		ServiceBusEndpointSuffix:     "servicebus.chinacloudapi.net",
		ServiceManagementVMDNSSuffix: "chinacloudapp.cn",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.cn",
		ContainerRegistryDNSSuffix:   "azurecr.io",
	}

	// GermanCloud is the cloud environment operated in Germany
	GermanCloud = Environment{
		Name:                         "AzureGermanCloud",
		ManagementPortalURL:          "http://portal.microsoftazure.de/",
		PublishSettingsURL:           "https://manage.microsoftazure.de/publishsettings/index",
		ServiceManagementEndpoint:    "https://management.core.cloudapi.de/",
		ResourceManagerEndpoint:      "https://management.microsoftazure.de/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.de/",
		GalleryEndpoint:              "https://gallery.cloudapi.de/",
		KeyVaultEndpoint:             "https://vault.microsoftazure.de/",
		GraphEndpoint:                "https://graph.cloudapi.de/",
		StorageEndpointSuffix:        "core.cloudapi.de",
		SQLDatabaseDNSSuffix:         "database.cloudapi.de",
		TrafficManagerDNSSuffix:      "azuretrafficmanager.de",
		KeyVaultDNSSuffix:            "vault.microsoftazure.de",
		ServiceBusEndpointSuffix:     "servicebus.cloudapi.de",
		ServiceManagementVMDNSSuffix: "azurecloudapp.de",
		ResourceManagerVMDNSSuffix:   "cloudapp.microsoftazure.de",
		ContainerRegistryDNSSuffix:   "azurecr.io",
	}

	//AzurestackCloud is the Azure enviornment running in customer datacenter
	AzurestackCloud = Environment{
		Name: "AzurestackCloud",
	}
)

// EnvironmentFromName returns an Environment based on the common name specified
func EnvironmentFromName(name string) (Environment, error) {
	name = strings.ToUpper(name)
	env, ok := environments[name]
	if !ok {
		return env, fmt.Errorf("autorest/azure: There is no cloud environment matching the name %q", name)
	}

	if strings.EqualFold(name, "AZURESTACKCLOUD") == true {

		//glog.Info("Reading Azure stack config from  /etc/kubernetes/azurestackcloud.json", dir)

		glog.Info("Reading Azure stack config from Directory", path.Join("etc", "kubernetes", "azurestackcloud.json"))

		fbytes, err := ioutil.ReadFile(path.Join("/etc", "kubernetes", "azurestackcloud.json"))
		if err != nil {
			return env, fmt.Errorf("Error opening Azure Stack Cloud Config file - azurestackcloud.json %q", err.Error())
		}

		err = json.Unmarshal(fbytes, &env)
		if err != nil {
			return env, fmt.Errorf("error parsing for Azure Stack Cloud Config %q", err.Error())
		}
	}
	return env, nil
}
