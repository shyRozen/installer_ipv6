locals {
  cidr_prefix = element(split("/", var.machine_cidr), 1)
 # base64ign = base64encode(var.ignition)
}

data "external" "gzip_base64" {
  program = ["bash", "-c", <<EOT
    # Create a temporary file for the large Ignition content
    tmpfile=$(mktemp)
    
    # Write the content of var.ignition to the temporary file using stdin
    cat > $tmpfile
    
    # Gzip and base64 encode the temporary file's content
    encoded_output=$(gzip -9 -c $tmpfile | base64 -w0)
    
    # Cleanup
    #rm -f $tmpfile
    
    # Output the result in JSON format
    echo "{\"encoded\": \"$encoded_output\"}"
  EOT
  ]
  
  query = {
    content = var.ignition
  }
}


resource "vsphere_virtual_machine" "vm" {
  name = var.vmname

  resource_pool_id = var.resource_pool_id
  datastore_id     = var.datastore_id
  num_cpus         = var.num_cpus
  memory           = var.memory
  guest_id         = var.guest_id
  folder           = var.folder_id
  enable_disk_uuid = "true"

  wait_for_guest_net_timeout  = "0"
  wait_for_guest_net_routable = "false"

  network_interface {
    network_id = var.network_id
  }

  disk {
    label            = "disk0"
    size             = 60
    thin_provisioned = var.disk_thin_provisioned
  }

  clone {
    template_uuid = var.template_uuid
  }

  
  extra_config = {
    "guestinfo.ignition.config.data"           = data.external.gzip_base64.result.encoded
    "guestinfo.ignition.config.data.encoding"  = "gzip+base64"
    "guestinfo.afterburn.initrd.network-kargs" = "ip=[${var.ipaddress}]::[${cidrhost(var.machine_cidr, 1)}]:${local.cidr_prefix}:${var.vmname}:ens192:none:${join(":", [for dns in var.dns_addresses : format("[%s]", dns)])}"
    "stealclock.enable"                        = "TRUE"
  }
}

