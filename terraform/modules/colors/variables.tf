variable "api_base_url" {
  type        = string
  description = "API Base URL, example: http://localhost:8080/api/v1"
}

variable "colors_endpoint" {
  type        = string
  description = "color change context path"
  default     = "color"
}

variable "color" {
  type        = string
  description = "color to set lights"
}

variable "supported_colors_endpoint" {
  type        = string
  description = "supported colors context path"
  default     = "colors"
}

