#  dynamics of module - ( fed from root folder )
variable "container_instance_count" {
  description = "suffix of container count - n'th container"
  type = number
}

variable "internal_port_number" {
  description = "port mapped inside the container to the host machine"
  type = number
}

variable "external_port_number" {
  description = "port of the host machine where the appplication is running"
  type = number
}

# output from the docker images created
variable "docker_image_id" {
  description = "container will be created from the image with the respective image id"
  type = string
}

variable "docker_image_name" {
  description = "name of the created container"
  type = string
}

variable "environment" {
  description = "environment for which the infra will get created"
  type = string
}

#  constants for module - ( default value specified )
variable "container_infinity_command" {
    description = "commands that run inside any container that is provisioned"
    type = list(string)
    default = ["tail","-f", "/dev/null"]
}