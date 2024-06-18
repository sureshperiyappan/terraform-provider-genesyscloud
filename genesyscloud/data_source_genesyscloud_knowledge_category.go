package genesyscloud

import (
	"context"
	"fmt"
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v131/platformclientv2"
)

func dataSourceKnowledgeCategory() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Genesys Cloud Knowledge Base Category. Select a category by name.",
		ReadContext: provider.ReadWithPooledClient(dataSourceKnowledgeCategoryRead),
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Knowledge base category name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"knowledge_base_name": {
				Description: "Knowledge base name",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceKnowledgeCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkConfig := m.(*provider.ProviderMeta).ClientConfig
	knowledgeAPI := platformclientv2.NewKnowledgeApiWithConfig(sdkConfig)

	name := d.Get("name").(string)
	knowledgeBaseName := d.Get("knowledge_base_name").(string)

	return util.WithRetries(ctx, 15*time.Second, func() *retry.RetryError {
		const pageSize = 100
		publishedKnowledgeBases, publishedResp, getPublishedErr := knowledgeAPI.GetKnowledgeKnowledgebases("", "", "", fmt.Sprintf("%v", pageSize), knowledgeBaseName, "", true, "", "")
		unpublishedKnowledgeBases, unpublishedResp, getUnpublishedErr := knowledgeAPI.GetKnowledgeKnowledgebases("", "", "", fmt.Sprintf("%v", pageSize), knowledgeBaseName, "", false, "", "")

		if getPublishedErr != nil {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError("genesyscloud_knowledge_category", fmt.Sprintf("Failed to get knowledge base %s | error: %s", knowledgeBaseName, getPublishedErr), publishedResp))
		}
		if getUnpublishedErr != nil {
			return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError("genesyscloud_knowledge_category", fmt.Sprintf("Failed to get knowledge base %s | error: %s", knowledgeBaseName, getUnpublishedErr), unpublishedResp))
		}

		noPublishedEntities := publishedKnowledgeBases.Entities == nil || len(*publishedKnowledgeBases.Entities) == 0
		noUnpublishedEntities := unpublishedKnowledgeBases.Entities == nil || len(*unpublishedKnowledgeBases.Entities) == 0
		if noPublishedEntities && noUnpublishedEntities {

			return retry.RetryableError(util.BuildWithRetriesApiDiagnosticError("genesyscloud_knowledge_category", fmt.Sprintf("no knowledge bases found with name %s", knowledgeBaseName), publishedResp))
		}

		// prefer published knowledge base
		for _, knowledgeBase := range *publishedKnowledgeBases.Entities {
			if knowledgeBase.Name != nil && *knowledgeBase.Name == knowledgeBaseName {
				knowledgeCategories, resp, getErr := knowledgeAPI.GetKnowledgeKnowledgebaseCategories(*knowledgeBase.Id, "", "", fmt.Sprintf("%v", pageSize), "", false, name, "", "", false)

				if getErr != nil {
					return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError("genesyscloud_knowledge_category", fmt.Sprintf("Failed to get knowledge category %s | error: %s", name, getErr), resp))
				}

				for _, knowledgeCategory := range *knowledgeCategories.Entities {
					if *knowledgeCategory.Name == name {
						id := fmt.Sprintf("%s,%s", *knowledgeCategory.Id, *knowledgeCategory.KnowledgeBase.Id)
						d.SetId(id)
						return nil
					}
				}
			}
		}
		// use unpublished knowledge base if unpublished doesn't exist
		for _, knowledgeBase := range *unpublishedKnowledgeBases.Entities {
			if knowledgeBase.Name != nil && *knowledgeBase.Name == knowledgeBaseName {
				knowledgeCategories, resp, getErr := knowledgeAPI.GetKnowledgeKnowledgebaseCategories(*knowledgeBase.Id, "", "", fmt.Sprintf("%v", pageSize), "", false, name, "", "", false)

				if getErr != nil {
					return retry.NonRetryableError(util.BuildWithRetriesApiDiagnosticError("genesyscloud_knowledge_category", fmt.Sprintf("Failed to get knowledge category %s | error: %s", name, getErr), resp))
				}

				for _, knowledgeCategory := range *knowledgeCategories.Entities {
					if *knowledgeCategory.Name == name {
						id := fmt.Sprintf("%s,%s", *knowledgeCategory.Id, *knowledgeCategory.KnowledgeBase.Id)
						d.SetId(id)
						return nil
					}
				}
			}
		}
		return nil
	})
}
