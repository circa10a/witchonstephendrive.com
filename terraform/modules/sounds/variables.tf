variable "api_base_url" {
  type        = string
  description = "API Base URL, example: http://localhost:8080/api/v1"
}

variable "sound_play_endpoint" {
  type        = string
  description = "sound play context path"
  default     = "sound"
}

variable "sound" {
  type        = string
  description = "sound to play"
}

variable "supported_sounds_endpoint" {
  type        = string
  description = "supported sounds context path"
  default     = "sounds"
}




