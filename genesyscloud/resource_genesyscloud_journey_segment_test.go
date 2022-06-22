package genesyscloud

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mypurecloud/platform-client-sdk-go/v72/platformclientv2"
)

func TestAccResourceJourneySegmentSession(t *testing.T) {
	runTestCase(t, "./test/data/journey_segment/basic_session_attributes")
}

func TestAccResourceJourneySegmentCustomer(t *testing.T) {
	runTestCase(t, "./test/data/journey_segment/basic_customer_attributes")
}

func TestAccResourceJourneySegmentContextOnly(t *testing.T) {
	runTestCase(t, "./test/data/journey_segment/context_only_to_journey_only")
}

func runTestCase(t *testing.T, folder string) {
	resourcePrefix, journeySegmentIdPrefix := setup(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps:             generateTestSteps(folder, resourcePrefix, journeySegmentIdPrefix),
		CheckDestroy:      testVerifyJourneySegmentsDestroyed,
	})
}

func setup(t *testing.T) (string, string) {
	const resourcePrefix = "genesyscloud_journey_segment."
	const journeySegmentIdPrefix = "terraform_test_"

	err := authorizeSdk()
	if err != nil {
		t.Fatal(err)
	}

	cleanupJourneySegments(journeySegmentIdPrefix)
	return resourcePrefix, journeySegmentIdPrefix
}

func generateTestSteps(folder string, resourcePrefix string, journeySegmentIdPrefix string) []resource.TestStep {
	var testSteps []resource.TestStep

	_, testCaseName := filepath.Split(folder)
	dirEntries, _ := os.ReadDir(folder)

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			resourceTf, _ := os.ReadFile(filepath.Join(folder, dirEntry.Name()))
			config := strings.Replace(string(resourceTf), "test_case", testCaseName, 1)
			testSteps = append(testSteps, resource.TestStep{Config: config})
		}
	}
	log.Printf("Generated %d test steps for %s testcase", len(dirEntries), testCaseName)

	testSteps = append(testSteps, resource.TestStep{
		ResourceName:      resourcePrefix + journeySegmentIdPrefix + testCaseName,
		ImportState:       true,
		ImportStateVerify: true,
	})

	return testSteps
}

func testVerifyJourneySegmentsDestroyed(state *terraform.State) error {
	journeyApi := platformclientv2.NewJourneyApiWithConfig(sdkConfig)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "genesyscloud_journey_segment" {
			continue
		}

		journeySegment, resp, err := journeyApi.GetJourneySegment(rs.Primary.ID)
		if journeySegment != nil {
			return fmt.Errorf("journey segment (%s) still exists", rs.Primary.ID)
		}

		if isStatus404(resp) {
			// Journey segment not found as expected
			continue
		}

		// Unexpected error
		return fmt.Errorf("unexpected error: %s", err)
	}
	// Success. All Journey segment destroyed
	return nil
}

func cleanupJourneySegments(journeySegmentIdPrefix string) {
	journeyApi := platformclientv2.NewJourneyApiWithConfig(sdkConfig)

	pageCount := 1 // Needed because of broken journey common paging
	for pageNum := 1; pageNum <= pageCount; pageNum++ {
		const pageSize = 100
		journeySegments, _, getErr := journeyApi.GetJourneySegments("", pageSize, pageNum, true, nil, nil, "")
		if getErr != nil {
			return
		}

		if journeySegments.Entities == nil || len(*journeySegments.Entities) == 0 {
			break
		}

		for _, journeySegment := range *journeySegments.Entities {
			if journeySegment.DisplayName != nil && strings.HasPrefix(*journeySegment.DisplayName, journeySegmentIdPrefix) {
				_, delErr := journeyApi.DeleteJourneySegment(*journeySegment.Id)
				if delErr != nil {
					diag.Errorf("failed to delete journey segment %s (%s): %s", *journeySegment.Id, *journeySegment.DisplayName, delErr)
					return
				}
				log.Printf("Deleted journey segment %s (%s)", *journeySegment.Id, *journeySegment.DisplayName)
			}
		}

		pageCount = *journeySegments.PageCount
	}
}
