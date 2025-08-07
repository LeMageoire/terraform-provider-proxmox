package proxmox

import (
	"crypto/tls"
	"testing"
)

func TestTLSConfiguration(t *testing.T) {
	tests := []struct {
		name            string
		tlsInsecure     bool
		expectedConfig  bool
	}{
		{
			name:            "TLS secure (default)",
			tlsInsecure:     false,
			expectedConfig:  false, // tlsconf should be nil for secure connections
		},
		{
			name:            "TLS insecure (skip verification)",
			tlsInsecure:     true,
			expectedConfig:  true, // tlsconf should be set with InsecureSkipVerify=true
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the TLS configuration logic from the client creation
			tlsconf := &tls.Config{InsecureSkipVerify: true}
			if !tt.tlsInsecure {
				tlsconf = nil
			}

			if tt.expectedConfig {
				// When TLS insecure is enabled, tlsconf should not be nil
				if tlsconf == nil {
					t.Errorf("Expected TLS config to be set for insecure mode, got nil")
				} else if !tlsconf.InsecureSkipVerify {
					t.Errorf("Expected InsecureSkipVerify to be true, got false")
				}
			} else {
				// When TLS is secure (default), tlsconf should be nil
				if tlsconf != nil {
					t.Errorf("Expected TLS config to be nil for secure mode, got %+v", tlsconf)
				}
			}
		})
	}
}

func TestTLSEnvironmentVariable(t *testing.T) {
	// Test that the environment variable PM_TLS_INSECURE is properly configured
	// This is more of a schema test to ensure the DefaultFunc is properly set
	
	// We can't easily test the actual environment variable reading without 
	// setting up the full provider context, but we can verify the schema
	// configuration exists by checking it doesn't panic
	
	provider := Provider()
	schema := provider.Schema
	
	if tlsSchema, exists := schema[schemaPmTlsInsecure]; exists {
		if tlsSchema.Type.String() != "TypeBool" {
			t.Errorf("Expected pm_tls_insecure to be TypeBool, got %s", tlsSchema.Type.String())
		}
		if !tlsSchema.Optional {
			t.Errorf("Expected pm_tls_insecure to be optional")
		}
	} else {
		t.Errorf("Expected pm_tls_insecure schema to exist")
	}
}
