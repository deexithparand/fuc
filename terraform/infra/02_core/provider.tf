terraform {
  required_providers {
    docker = {
        source = "kreuzwerker/docker"
        version = "2.12.2"
    }
  }

  backend "local" {}
}

// configuration for each provider inside terraform - required provider - providers
provider "docker" {
  host = "unix:///var/run/docker.sock"
}