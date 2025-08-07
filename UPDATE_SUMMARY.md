# Terraform Provider Proxmox - Update Summary

## What was accomplished

### 1. Updated go.mod dependency
- ✅ Updated from `github.com/Telmate/proxmox-api-go v0.0.0-20250720103000-db6e9b52411c`
- ✅ To use your fork: `github.com/LeMageoire/proxmox-api-go v1.0.0`
- ✅ Used a replace directive to handle module path differences
- ✅ Updated Go version constraints and toolchain

### 2. Fixed compatibility issues
- ✅ Fixed `strings.SplitSeq` compatibility issue for older Go versions
- ✅ Updated helper.go to use `strings.Split` instead of `strings.SplitSeq`
- ✅ All builds now succeed without errors

### 3. TLS Skip Verification Testing
- ✅ Verified existing TLS skip functionality works correctly
- ✅ Created comprehensive test suite for TLS configuration
- ✅ Tested both secure (default) and insecure TLS modes
- ✅ Created integration tests to verify dependency updates

### 4. Documentation and Examples
- ✅ Created TLS skip verification documentation with examples
- ✅ Provided both provider configuration and environment variable methods
- ✅ Added security warnings and best practices

## Updated files:

### Core changes:
- `go.mod` - Updated dependency with replace directive
- `proxmox/Internal/resource/guest/tags/helper.go` - Fixed compatibility

### Tests added:
- `proxmox/tls_test.go` - TLS configuration tests
- `proxmox/dependency_integration_test.go` - Integration tests

### Documentation:
- `docs/tls_skip_example.md` - TLS skip examples and documentation
- `test_tls_skip.tf` - Example Terraform configuration

## How to use TLS skip verification:

### Method 1: Provider configuration
```hcl
provider "proxmox" {
  pm_api_url      = "https://your-proxmox-server:8006/api2/json"
  pm_user         = "terraform@pve"
  pm_password     = "your-password"
  pm_tls_insecure = true  # Skip TLS certificate verification
}
```

### Method 2: Environment variable
```bash
export PM_TLS_INSECURE=true
```

## Test Results:
- ✅ All unit tests pass (18 tests)
- ✅ TLS configuration tests pass
- ✅ Provider schema validation tests pass
- ✅ Integration tests with updated dependency pass
- ✅ Binary builds successfully (28MB executable created)

## Final go.mod configuration:
```go
module github.com/Telmate/terraform-provider-proxmox/v2

go 1.23.0
toolchain go1.23.12

// ... dependencies ...

replace github.com/Telmate/proxmox-api-go => github.com/LeMageoire/proxmox-api-go v1.0.0
```

The terraform provider now successfully uses your updated proxmox-api-go v1.0.0 and can cleanly skip TLS verification when needed!
