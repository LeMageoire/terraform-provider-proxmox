# TLS Skip Verification Example

This example demonstrates how to configure the Terraform Proxmox provider to skip TLS certificate verification. This is useful when connecting to Proxmox servers with self-signed certificates or when you don't have the CA certificate that issued the Proxmox API certificate.

## Configuration

### Method 1: Using provider configuration

```hcl
terraform {
  required_providers {
    proxmox = {
      source = "telmate/proxmox"
      version = ">= 2.0.0"
    }
  }
}

provider "proxmox" {
  pm_api_url      = "https://your-proxmox-server:8006/api2/json"
  pm_user         = "terraform@pve"
  pm_password     = "your-password"
  pm_tls_insecure = true  # Skip TLS certificate verification
}
```

### Method 2: Using environment variables

```bash
export PM_API_URL="https://your-proxmox-server:8006/api2/json"
export PM_USER="terraform@pve"
export PM_PASSWORD="your-password"
export PM_TLS_INSECURE=true  # Skip TLS certificate verification
```

```hcl
provider "proxmox" {
  # Configuration will be read from environment variables
}
```

## Security Note

⚠️ **Warning**: Setting `pm_tls_insecure = true` or `PM_TLS_INSECURE=true` disables TLS certificate verification, which reduces security. Only use this in test environments or when you understand the security implications.

For production environments, it's recommended to:
1. Use proper SSL certificates signed by a trusted CA
2. Configure your system to trust the CA that signed the Proxmox certificate
3. Use certificate pinning if possible

## Testing the Configuration

You can test the TLS configuration by creating a simple resource:

```hcl
data "proxmox_version" "version" {}

output "proxmox_version" {
  value = data.proxmox_version.version.version
}
```

Run `terraform plan` to verify that the provider can connect to your Proxmox server with the TLS configuration.
