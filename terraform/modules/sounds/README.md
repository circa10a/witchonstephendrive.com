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
| [httpclient_request.sound_play](https://registry.terraform.io/providers/dmachard/http-client/latest/docs/data-sources/request) | data source |
| [httpclient_request.supported_sounds](https://registry.terraform.io/providers/dmachard/http-client/latest/docs/data-sources/request) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_base_url"></a> [api\_base\_url](#input\_api\_base\_url) | API Base URL, example: https://witchonstephendrive.com/api/v1 | `string` | n/a | yes |
| <a name="input_sound"></a> [sound](#input\_sound) | sound to play | `string` | n/a | yes |
| <a name="input_sound_play_endpoint"></a> [sound\_play\_endpoint](#input\_sound\_play\_endpoint) | sound play context path | `string` | `"sound"` | no |
| <a name="input_supported_sounds_endpoint"></a> [supported\_sounds\_endpoint](#input\_supported\_sounds\_endpoint) | supported sounds context path | `string` | `"sounds"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_sound_play_response"></a> [sound\_play\_response](#output\_sound\_play\_response) | n/a |
| <a name="output_supported_sounds"></a> [supported\_sounds](#output\_supported\_sounds) | n/a |
