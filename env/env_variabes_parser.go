package env

import (
	"log"
	"os"
)

func GetKeyVaultUrl(keyVaultUrlEnvVariable string) string {
	vaultUrl := os.Getenv(keyVaultUrlEnvVariable)

	if len(vaultUrl) == 0 {
		log.Fatalf("%s env variable is empty!", keyVaultUrlEnvVariable)
	}

	return vaultUrl
}
