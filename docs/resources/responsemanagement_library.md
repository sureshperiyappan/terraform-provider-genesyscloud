---
page_title: "genesyscloud_responsemanagement_library Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud responsemanagement library
---
# genesyscloud_responsemanagement_library (Resource)

Genesys Cloud responsemanagement library

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [GET /api/v2/responsemanagement/libraries](https://developer.genesys.cloud/devapps/api-explorer#get-api-v2-responsemanagement-libraries)
* [POST /api/v2/responsemanagement/libraries](https://developer.genesys.cloud/devapps/api-explorer#post-api-v2-responsemanagement-libraries)
* [GET /api/v2/responsemanagement/libraries/{libraryId}](https://developer.genesys.cloud/devapps/api-explorer#get-api-v2-responsemanagement-libraries--libraryId-)
* [PUT /api/v2/responsemanagement/libraries/{libraryId}](https://developer.genesys.cloud/devapps/api-explorer#put-api-v2-responsemanagement-libraries--libraryId-)
* [DELETE /api/v2/responsemanagement/libraries/{libraryId}](https://developer.genesys.cloud/devapps/api-explorer#delete-api-v2-responsemanagement-libraries--libraryId-)

## Example Usage

```terraform
resource "genesyscloud_responsemanagement_library" "example_responsemanagement_library" {
  name = "Example library name"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The library name.

### Read-Only

- `id` (String) The ID of this resource.
