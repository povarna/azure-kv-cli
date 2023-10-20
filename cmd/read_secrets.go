package cmd

import (
	"log"
	"os"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/spf13/cobra"
)

var readSecretsCmd = &cobra.Command{
	Use:   `readSecrets`,
	Short: `Read Key Vault Secret`,
	Long: `Read Azure Key Vault secret from the provided host 
For example:
azure_key_vault readSecret -s <secret_key>, <secret_key>`,

	Run: readKeyVaultSecrets,
}

func init() {
	rootCmd.AddCommand(readSecretsCmd)
	readSecretsCmd.Flags().StringSliceP("secrets", "s", []string{}, "A list of key vault secrets separated by comma")
}

func readKeyVaultSecrets(cmd *cobra.Command, args []string) {
	vaultUrl := os.Getenv("AZURE_KEY_VAULT_URL")

	if len(vaultUrl) == 0 {
		log.Fatal("AZURE_KEY_VAULT_URL env variable is empty!")
	}

	secrets, _ := cmd.Flags().GetStringSlice("secrets")
	if len(secrets) == 0 {
		log.Fatal("No secrets key provided")
	}

	azureClient, err := azure_key_vault.InitAzureClient(vaultUrl)

	if err != nil {
		log.Fatalf("Unable to obtain azure connection")
	}

	azureClient.ReadSecrets(secrets)
}
