variable "api_base_url" {
  type        = string
  description = "API Base URL, example: http://localhost:8080/api/v1"
  default     = "http://localhost:8080/api/v1"
}

variable "color" {
  type        = string
  description = "color to set lights"
  default     = ""
}

variable "sound" {
  type        = string
  description = "sound to play"
  default     = ""
}
