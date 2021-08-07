module "colors" {
    source = "./modules/colors"
    count = length(var.color) > 0 ? 1 :0
    api_base_url = var.api_base_url
    color = var.color
}

module "sounds" {
    source = "./modules/sounds"
    count = length(var.color) > 0 ? 1 :0
    api_base_url = var.api_base_url
    sound = var.sound
}