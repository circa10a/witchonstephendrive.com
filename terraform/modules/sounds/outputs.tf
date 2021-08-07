output "sound_play_response" {
  value = jsondecode(data.httpclient_request.sound_play.response_body)
}

output "supported_sounds" {
  value = jsondecode(data.httpclient_request.supported_sounds.response_body).supportedSounds
}

