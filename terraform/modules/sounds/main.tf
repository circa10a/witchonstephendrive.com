terraform {
  required_providers {
    httpclient = {
      source = "dmachard/http-client"
    }
  }
}

data "httpclient_request" "sound_play" {
  url            = "${var.api_base_url}/${var.sound_play_endpoint}/${var.sound}"
  request_method = "POST"
}

data "httpclient_request" "supported_sounds" {
  url = "${var.api_base_url}/${var.supported_sounds_endpoint}"
}
