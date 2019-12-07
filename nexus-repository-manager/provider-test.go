package nexus

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"nexus": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("NEXUS_REPOSITORY_MANAGER_URL"); v == "" {
		t.Fatal("NEXUS_REPOSITORY_MANAGER_URL must be set for acceptance tests")
	}
	if v := os.Getenv("NEXUS_REPOSITORY_MANAGER_AUTH_USERNAME"); v == "" {
		t.Fatal("NEXUS_REPOSITORY_MANAGER_AUTH_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("NEXUS_REPOSITORY_MANAGER_AUTH_PASSWORD"); v == "" {
		t.Fatal("NEXUS_REPOSITORY_MANAGER_AUTH_PASSWORD must be set for acceptance tests")
	}
}
