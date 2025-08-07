package proxmox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUpdatedDependencyIntegration(t *testing.T) {
	// Test that the updated proxmox-api-go dependency is properly integrated
	// and that TLS configuration works correctly
	
	provider := Provider()
	
	// Test provider schema includes TLS configuration
	if provider.Schema == nil {
		t.Fatal("Provider schema is nil")
	}
	
	tlsSchema, exists := provider.Schema[schemaPmTlsInsecure]
	if !exists {
		t.Fatal("TLS insecure schema not found")
	}
	
	if tlsSchema.Type != schema.TypeBool {
		t.Errorf("Expected TLS insecure to be TypeBool, got %v", tlsSchema.Type)
	}
	
	// Test provider resources are available
	expectedResources := []string{
		"proxmox_vm_qemu",
		"proxmox_lxc", 
		"proxmox_pool",
		"proxmox_cloud_init_disk",
		"proxmox_storage_iso",
	}
	
	for _, resourceName := range expectedResources {
		if _, exists := provider.ResourcesMap[resourceName]; !exists {
			t.Errorf("Expected resource %s not found", resourceName)
		}
	}
	
	// Test provider data sources are available
	expectedDataSources := []string{
		"proxmox_ha_groups",
	}
	
	for _, dataSourceName := range expectedDataSources {
		if _, exists := provider.DataSourcesMap[dataSourceName]; !exists {
			t.Errorf("Expected data source %s not found", dataSourceName)
		}
	}
}

func TestProviderConfigurationWithTLS(t *testing.T) {
	// Test provider configuration with TLS settings
	provider := Provider()
	
	// Create a resource data with TLS insecure set to true
	rawConfig := map[string]interface{}{
		schemaPmApiUrl:      "https://test-server:8006/api2/json",
		schemaPmUser:        "test@pve",
		schemaPmPassword:    "testpass",
		schemaPmTlsInsecure: true,
		schemaPmDebug:       false,
		schemaPmTimeout:     300,
	}
	
	d := schema.TestResourceDataRaw(t, provider.Schema, rawConfig)
	
	// Verify the TLS insecure setting is properly read
	tlsInsecure := d.Get(schemaPmTlsInsecure).(bool)
	if !tlsInsecure {
		t.Errorf("Expected TLS insecure to be true, got false")
	}
	
	// Test with TLS secure (default)
	rawConfigSecure := map[string]interface{}{
		schemaPmApiUrl:   "https://test-server:8006/api2/json",
		schemaPmUser:     "test@pve", 
		schemaPmPassword: "testpass",
		// pm_tls_insecure not set, should default to false
	}
	
	dSecure := schema.TestResourceDataRaw(t, provider.Schema, rawConfigSecure)
	tlsSecure := dSecure.Get(schemaPmTlsInsecure).(bool)
	if tlsSecure {
		t.Errorf("Expected TLS insecure to be false by default, got true")
	}
}

func TestClientCreationParameters(t *testing.T) {
	// Test that client creation parameters are correctly passed
	// This tests the integration with the updated proxmox-api-go dependency
	
	testCases := []struct {
		name        string
		tlsInsecure bool
		expectError bool
	}{
		{
			name:        "TLS Secure",
			tlsInsecure: false,
			expectError: true, // Will error because of invalid credentials, but that's expected
		},
		{
			name:        "TLS Insecure", 
			tlsInsecure: true,
			expectError: true, // Will error because of invalid credentials, but that's expected
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test the client creation logic (without actually connecting)
			_, err := getClient(
				"https://test-server:8006/api2/json",
				"test@pve",
				"testpass",
				"", // no token ID
				"", // no token secret
				"", // no OTP
				tc.tlsInsecure,
				"", // no headers
				300, // timeout
				false, // no debug
				"", // no proxy
			)
			
			// We expect an error because we're using invalid credentials,
			// but the function should at least accept the parameters without panicking
			if err == nil {
				t.Errorf("Expected error due to invalid credentials")
			}
		})
	}
}
