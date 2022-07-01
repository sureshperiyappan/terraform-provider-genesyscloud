---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "genesyscloud_journey_action_map Data Source - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Data source for Genesys Cloud Action Map. Select a journey action map by name
---

# genesyscloud_journey_action_map (Data Source)

Data source for Genesys Cloud Action Map. Select a journey action map by name

## Example Usage

```terraform
data "genesyscloud_journey_action_map" "example_journey_action_map_data" {
  name = "example journey action map name"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Journey Action Map name.

### Read-Only

- `id` (String) The ID of this resource.

