package cmd

import (
	"log"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/povarna/azure-kv-cli/env"
	"github.com/spf13/cobra"
)

var readSecretCmd = &cobra.Command{
	Use:   `readSecret`,
	Short: `Read Key Vault Secret`,
	Long: `Read Azure Key Vault secret from the provided host 
For example:
azure_key_vault readSecret -s <secret_key> -v <secret_version>`,

	Run: readKeyVaultSecret,
}

func init() {
	rootCmd.AddCommand(readSecretCmd)
	readSecretCmd.Flags().StringP("secretKey", "k", "", "Secret Key")
	readSecretCmd.Flags().StringP("version", "v", "", "Secret Version")
}

func readKeyVaultSecret(cmd *cobra.Command, args []string) {
	vaultUrl := env.GetKeyVaultUrl("AZURE_KEY_VAULT_URL")

	secret, _ := cmd.Flags().GetString("secretKey")
	if len(secret) == 0 {
		log.Fatal("No secret key provided")
	}

	version, _ := cmd.Flags().GetString("version")

	azureClient, err := azure_key_vault.InitAzureClient(vaultUrl)

	if err != nil {
		log.Fatalf("Unable to obtain azure connection")
	}

	resp, err := azureClient.ReadSecret(secret, version)
	if err != nil {
		log.Printf("Unable to get secret: %v", err)
	}

	log.Printf("%s:%s\n", secret, *resp.Value)
}
