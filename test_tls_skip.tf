terraform {
  required_providers {
    proxmox = {
      source = "local/terraform-provider-proxmox"
      version = "1.0.0"
    }
  }
}

provider "proxmox" {
  pm_api_url      = "https://your-proxmox-server:8006/api2/json"
  pm_user         = "terraform@pve"
  pm_password     = "your-password"
  pm_tls_insecure = true  # This enables TLS skip verification
  pm_debug        = true
}

# Test data source to verify connectivity
data "proxmox_pool" "default" {
  pool_id = "default"
}

output "pool_data" {
  value = data.proxmox_pool.default
  sensitive = true
}
