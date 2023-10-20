package azure_key_vault

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzureKeyVault struct {
	client *azsecrets.Client
}

func InitAzureClient(vaultUrl string) (*AzureKeyVault, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain credentials: %v", err)
	}

	client, err := azsecrets.NewClient(vaultUrl, cred, nil)

	if err != nil {
		return new(AzureKeyVault), err
	}

	return &AzureKeyVault{client}, nil
}

func (azureKeyVault *AzureKeyVault) ReadSecret(secret string, version string) (azsecrets.GetSecretResponse, error) {
	return azureKeyVault.client.GetSecret(context.TODO(), secret, version, nil)
}

func (azureKeyVault *AzureKeyVault) ReadSecrets(secrets []string) {
	version := ""

	for _, secret := range secrets {
		resp, err := azureKeyVault.ReadSecret(strings.Trim(secret, " "), version)

		if err != nil {
			log.Printf("Unable to get secret: %v", err)
		}

		fmt.Printf("%s:%s\n", secret, *resp.Value)
	}
}
