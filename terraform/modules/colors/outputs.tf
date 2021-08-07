output "color_change_response" {
  value = jsondecode(data.httpclient_request.color_change.response_body)
}

output "supported_colors" {
  value = jsondecode(data.httpclient_request.supported_colors.response_body).supportedColors
}
