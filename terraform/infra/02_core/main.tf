# pulls the image
resource "docker_image" "alpine" {
  name = "alpine:latest"
}

# Create a container
resource "docker_container" "linux_container_01" {
  image = docker_image.alpine.latest
  name  = "linux-container-01"
  ports {
    internal = 80
    external = 3000
  }
  command = [ "tail", "-f", "/dev/null"]
}