package cmd

import (
	"log"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/povarna/azure-kv-cli/env"
	"github.com/spf13/cobra"
)

var readSecretsCmd = &cobra.Command{
	Use:   `readSecrets`,
	Short: `Read Key Vault Secrets`,
	Long: `Read Azure Key Vault list of secrets from the provided host 
For example:
azure_key_vault readSecrets -s <secret_key>,<secret_key>`,

	Run: readKeyVaultSecrets,
}

func init() {
	rootCmd.AddCommand(readSecretsCmd)
	readSecretsCmd.Flags().StringSliceP("secretKeys", "s", []string{}, "A list of key vault secrets separated by comma")
}

func readKeyVaultSecrets(cmd *cobra.Command, args []string) {
	vaultUrl := env.GetKeyVaultUrl("AZURE_KEY_VAULT_URL")

	secrets, _ := cmd.Flags().GetStringSlice("secretKeys")
	if len(secrets) == 0 {
		log.Fatal("No secrets key provided")
	}

	azureClient, err := azure_key_vault.InitAzureClient(vaultUrl)

	if err != nil {
		log.Fatalf("Unable to obtain azure connection")
	}

	azureClient.ReadSecrets(secrets)
}
