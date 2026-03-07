# assumed command

# fuc up --env dev --distro apline --tag latest --container container-name --count 3 --port 8000:80
# fuc down --env dev
# fuc show ( show the stat of the containers )

# pull image from docker image module
module "docker_images" {
  source = "./modules/docker-images"

  // values for each image are mapped as kv pairs and passed
  for_each = {

    alpine = {
      image_name = "alpine"
      image_tag = "latest"
      keep_locally = false
    }

    ubuntu = {
      image_name = "ubuntu"
      image_tag = "latest"
      keep_locally = false
    }

  }

  # from tfvars
  environment = var.environment 

  # from the for_each
  image_name = each.value.image_name
  image_tag = each.value.image_tag
  keep_locally = each.value.keep_locally
}


module "docker_container" {
  source = "./modules/docker-containers"

  for_each = {

    alpine = {

      docker_image_id = module.docker_images["alpine"].image_id
      docker_image_name = module.docker_images["alpine"].image_name
      container_instance_count = 2
      external_port_number = 8000 
      internal_port_number = 80

    }

    ubuntu = {
      docker_image_id = module.docker_images["ubuntu"].image_id
      docker_image_name = module.docker_images["ubuntu"].image_name
      container_instance_count = 1
      external_port_number = 5000 
      internal_port_number = 80  
    }
  }

  environment = var.environment 
  docker_image_id = each.value.docker_image_id
  docker_image_name = each.value.docker_image_name
  container_instance_count = each.value.container_instance_count
  external_port_number = each.value.external_port_number
  internal_port_number = each.value.internal_port_number  
}