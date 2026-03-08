environment = "dev" 

distro_config_set =  {
    alpine = {
        image_name = "alpine"
        image_tag = "latest"
        keep_locally = false

        container_instance_count = 2
        external_port_number     = 8000
        internal_port_number     = 80
    }
}