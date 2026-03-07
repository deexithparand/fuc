resource "docker_container" "container" {

    // count is a meta argument ( internally does for_each )
    count = var.container_instance_count

    image = var.docker_image_id
    name = "${replace(var.docker_image_name,":","-")}-${var.environment}-container-${count.index}"

    ports {
        internal = var.internal_port_number
        external = var.external_port_number + count.index
    }

    command = var.container_infinity_command
}

# Create a container
# resource "docker_container" "linux_container_01" {
#   image = docker_image.alpine.latest
#   name  = "linux-container-01"
#   ports {
#     internal = 80
#     external = 3000
#   }
#   command = ["tail", "-f", "/dev/null"]
# }