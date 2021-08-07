terraform {
  required_providers {
    httpclient = {
      source = "dmachard/http-client"
    }
  }
}

data "httpclient_request" "color_change" {
  url            = "${var.api_base_url}/${var.colors_endpoint}/${var.color}"
  request_method = "POST"
}

data "httpclient_request" "supported_colors" {
  url = "${var.api_base_url}/${var.supported_colors_endpoint}"
}
