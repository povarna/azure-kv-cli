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

	return &AzureKeyVault{client: client}, nil
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

func (azureKeyVault *AzureKeyVault) SetSecret(secretKey string, secretValue string) {
	params := azsecrets.SetSecretParameters{Value: &secretValue}
	_, err := azureKeyVault.client.SetSecret(context.TODO(), secretKey, params, nil)
	if err != nil {
		log.Fatalf("Failed to create secret: %s. Error: %+v", secretKey, err)
	}
	fmt.Printf("Secret: %s created successfully\n", secretKey)
}

func (azureKeyVault *AzureKeyVault) ListSecrets() {
	pager := azureKeyVault.client.NewListSecretsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("Unable to get the next secret page. Error: %v", err)
		}

		for _, secret := range page.Value {
			secretId := string(*secret.ID)
			key := secretId[strings.LastIndex(secretId, "/")+1:]

			fmt.Printf("%s\n", key)
		}
	}
}
