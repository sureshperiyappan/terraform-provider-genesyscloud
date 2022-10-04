---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_outbound_sequence Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Data source for Genesys Cloud Outbound Sequence. Select a Outbound Sequence by name.
---

# genesyscloud_outbound_sequence (Data Source)

Data source for Genesys Cloud Outbound Sequence. Select a Outbound Sequence by name.

## Example Usage

```terraform
data "genesyscloud_outbound_sequence" "example_outbound_sequence" {
  name = "Example Sequence"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String) Outbound Sequence name.

### Read-Only

- `id` (String) The ID of this resource.

