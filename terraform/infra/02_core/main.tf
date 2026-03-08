# assumed command

# loc up --env dev --distro apline --tag latest --count 3 --port 8000:80
                # ✅            ✅            ✅           ✅         ✅  ✅
# loc down --env dev
# loc show ( show the stat of the containers )

# pull image from docker image module
module "docker_images" {
  source = "./modules/docker-images"

  for_each = var.distro_config_set

  # from tfvars
  environment = var.environment 

  # from the for_each
  image_name = each.value.image_name
  image_tag = each.value.image_tag
  keep_locally = each.value.keep_locally
}


module "docker_container" {
  source = "./modules/docker-containers"

  for_each = var.distro_config_set

  # from tfvars
  environment = var.environment 

  # from output of the docker images
  docker_image_id = module.docker_images[each.key].image_id
  docker_image_name = module.docker_images[each.key].image_name

  # from the for_each
  container_instance_count = each.value.container_instance_count
  external_port_number = each.value.external_port_number
  internal_port_number = each.value.internal_port_number  
}