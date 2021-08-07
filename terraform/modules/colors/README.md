## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_httpclient"></a> [httpclient](#provider\_httpclient) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [httpclient_request.color_change](https://registry.terraform.io/providers/dmachard/http-client/latest/docs/data-sources/request) | data source |
| [httpclient_request.supported_colors](https://registry.terraform.io/providers/dmachard/http-client/latest/docs/data-sources/request) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_base_url"></a> [api\_base\_url](#input\_api\_base\_url) | API Base URL, example: https://witchonstephendrive.com/api/v1 | `string` | n/a | yes |
| <a name="input_color"></a> [color](#input\_color) | color to set lights | `string` | n/a | yes |
| <a name="input_colors_endpoint"></a> [colors\_endpoint](#input\_colors\_endpoint) | color change context path | `string` | `"color"` | no |
| <a name="input_supported_colors_endpoint"></a> [supported\_colors\_endpoint](#input\_supported\_colors\_endpoint) | supported colors context path | `string` | `"colors"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_color_change_response"></a> [color\_change\_response](#output\_color\_change\_response) | n/a |
| <a name="output_supported_colors"></a> [supported\_colors](#output\_supported\_colors) | n/a |
