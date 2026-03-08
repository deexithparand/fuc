environment = "prod"

distro_config_set =  {
    ubuntu = {
        image_name = "ubuntu"
        image_tag = "latest"
        keep_locally = false

        container_instance_count = 2
        external_port_number     = 5000
        internal_port_number     = 80
    }
}