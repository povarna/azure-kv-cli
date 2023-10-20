Azure-KV-Cli
----------------------
A  CLI for interacting with Azure Key Vault

Prerequisites
- Install `GO` > `1.20`
- Install `azure-cli`

Build
--------------------
```
$ go build -o bin/azure-kv-cli
```

Use the cli directly from the `$GOPATH/bin` path:
```
$ go build
$ go install
``` 

Running the application
--------------------

1. Login to azure:
```
$ az login
```

2. Export a Azure Key Vault URL:
```
export AZURE_KEY_VAULT_URL=""
```

3. Run the CLI:
```
Usage:
Azure Key Vault cli application

Usage:
  azure-kv-cli [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  help           Help about any command
  listSecrets    List all Key Vault Secrets
  migrateSecrets Migrate Key Vault Secrets
  readSecret     Read Key Vault Secret
  readSecrets    Read Key Vault Secrets
  setSecret      Set Key Vault Secret

Flags:
  -h, --help     help for azure-kv-cli
  -t, --toggle   Help message for toggle

Use "azure-kv-cli [command] --help" for more information about a command.
```

Go Lang SDK Source:
https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets#pkg-examples