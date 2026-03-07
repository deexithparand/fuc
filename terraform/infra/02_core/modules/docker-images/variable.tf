#  dynamics of module - ( fed from root folder )
variable "image_name" {
    description = "name of the image to be pulled"
    type = string
}

variable "image_tag" {
    description = "tag of the respective image that will be pulled"
    type = string
}

variable "keep_locally" {
    description = "if enabled to true, images won't get cleared on terraform destroy"
    type = bool
}

# fed from tfvars
variable "environment" {
  description = "environment where the infrastructure is provisioned"
  type = string
}

#  constants for module - ( default value specified )
variable "platform" {
    description = "for running on my mac ( deexith's mac = constant )"
    type = string
    default = "linux/arm64"
}


