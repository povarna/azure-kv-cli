package cmd

import (
	"log"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/povarna/azure-kv-cli/env"
	"github.com/spf13/cobra"
)

var setSecretCmd = &cobra.Command{
	Use:   `setSecret`,
	Short: `Set Key Vault Secret`,
	Long: `Set Azure Key Vault secret on the provided host 
For example:
azure_key_vault setSecret -k <secret_key> -v <secret_value>`,

	Run: setKeyVaultSecret,
}

func init() {
	rootCmd.AddCommand(setSecretCmd)
	setSecretCmd.Flags().StringP("key", "k", "", "Azure key vault secret key")
	setSecretCmd.Flags().StringP("value", "v", "", "Azure key vault secret value")
}

func setKeyVaultSecret(cmd *cobra.Command, args []string) {
	vaultUrl := env.GetKeyVaultUrl("AZURE_KEY_VAULT_URL")

	secretKey, _ := cmd.Flags().GetString("key")
	if len(secretKey) == 0 {
		log.Fatal("No secrets key provided")
	}

	secretValue, _ := cmd.Flags().GetString("value")
	if len(secretValue) == 0 {
		log.Fatal("No secrets value provided")
	}

	azureClient, err := azure_key_vault.InitAzureClient(vaultUrl)

	if err != nil {
		log.Fatalf("Unable to obtain azure connection")
	}

	azureClient.SetSecret(secretKey, secretValue)
}
