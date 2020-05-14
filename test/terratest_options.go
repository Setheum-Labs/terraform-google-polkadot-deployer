package test

import (
    "fmt"
    "os"
    "strings"
    "testing"

    "github.com/gruntwork-io/terratest/modules/random"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/test-structure"
    "github.com/stretchr/testify/require"
)

func createTerraformOptions(t *testing.T, terraformDir string) {
    nodeCount := 2
    servicePort := 30100
    uniqueID := strings.ToLower(random.UniqueId())
    clusterName := fmt.Sprintf("test-polkadot-%s", uniqueID)

    gcpProjectId := os.Getenv("GCP_PROJECT_ID")
    require.NotEmpty(t, gcpProjectId, "GCP_PROJECT_ID variable is not set")

    terraformOptions := &terraform.Options{
        TerraformDir: terraformDir,
        Vars: map[string]interface{} {
            "cluster_name":   clusterName,
            "gcp_project_id": gcpProjectId,
            "location":       "europe-west4-b",
            "machine_type":   "n1-standard-1",
            "node_count":     nodeCount,
        },
        NoColor: true,
    }

    test_structure.SaveInt(t, terraformDir, "nodeCount", nodeCount)
    test_structure.SaveInt(t, terraformDir, "nodePort", servicePort)
    test_structure.SaveString(t, terraformDir, "clusterName", clusterName)
    test_structure.SaveString(t, terraformDir, "uniqueID", uniqueID)
    test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
}

func createHelmOptions(t *testing.T, terraformDir string) {
    helmValues := map[string]string{
        "image.repo":   "nginx",
        "image.tag":    "1.8",
        "service.type": "NodePort",
        "service.port": "30100",
    }

    helmValuesFile := test_structure.FormatTestDataPath(terraformDir, "HelmValues.json")
    test_structure.SaveString(t, terraformDir, "helmValuesFile", helmValuesFile)
    test_structure.SaveTestData(t, helmValuesFile, helmValues)
}
