package cmd

import (
	"log"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/povarna/azure-kv-cli/env"
	"github.com/spf13/cobra"
)

var listSecretsCmd = &cobra.Command{
	Use:   `listSecrets`,
	Short: `List all Key Vault Secrets`,
	Long: `List all Azure Key Vault secret keys from the provided host 
For example:
azure_key_vault listSecrets`,

	Run: listKeyVaultSecrets,
}

func init() {
	rootCmd.AddCommand(listSecretsCmd)
}

func listKeyVaultSecrets(cmd *cobra.Command, args []string) {
	vaultUrl := env.GetKeyVaultUrl("AZURE_KEY_VAULT_URL")
	azureClient, err := azure_key_vault.InitAzureClient(vaultUrl)

	if err != nil {
		log.Fatalf("Unable to obtain azure connection")
	}

	azureClient.ListSecrets()
}
