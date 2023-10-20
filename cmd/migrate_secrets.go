package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/povarna/azure-kv-cli/azure_key_vault"
	"github.com/povarna/azure-kv-cli/env"
	"github.com/spf13/cobra"
)

var migrateSecretsCmd = &cobra.Command{
	Use:   `migrateSecrets`,
	Short: `Migrate Key Vault Secrets`,
	Long: `Migrate secrets from SOURCE_AZURE_KEY_VAULT_URL to DESTINATION_AZURE_KEY_VAULT_URL
For example:
azure_key_vault migrateSecrets -s <secret_key>,<secret_key>`,

	Run: migrateKeyVaultSecrets,
}

func init() {
	rootCmd.AddCommand(migrateSecretsCmd)
	migrateSecretsCmd.Flags().StringSliceP("secretKeys", "s", []string{}, "A list of key vault secrets separated by comma")
}

func migrateKeyVaultSecrets(cmd *cobra.Command, args []string) {
	sourceVaultUrl := env.GetKeyVaultUrl("SOURCE_AZURE_KEY_VAULT_URL")
	sourceVaultClient, err := azure_key_vault.InitAzureClient(sourceVaultUrl)

	if err != nil {
		log.Printf("Unable to obtain source vault client for: %s", sourceVaultUrl)
	}

	destinationVaultUrl := env.GetKeyVaultUrl("DESTINATION_AZURE_KEY_VAULT_URL")
	destinationVaultClient, err := azure_key_vault.InitAzureClient(destinationVaultUrl)

	if err != nil {
		log.Printf("Unable to obtain source vault client for: %s", destinationVaultUrl)
	}

	secrets, _ := cmd.Flags().GetStringSlice("secretKeys")
	if len(secrets) == 0 {
		log.Fatal("No secrets key provided")
	}

	version := ""
	for _, secret := range secrets {
		resp, err := sourceVaultClient.ReadSecret(strings.Trim(secret, " "), version)

		if err != nil {
			log.Printf("Unable to get secret: %v", err)
		}
		fmt.Printf("Setting up secret key: %s to the destination kv: %s \n", secret, destinationVaultUrl)

		destinationVaultClient.SetSecret(secret, *resp.Value)
	}

}
