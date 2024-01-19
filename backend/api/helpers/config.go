package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lodashventure/nlp/infrastructure"
	"github.com/lodashventure/nlp/model"
)

func PrepareCredential() map[string]model.Credential {
	credentialPath := os.Getenv("CREDENTIAL_PATH")
	if credentialPath == "" {
		infrastructure.Log.Panicln(nil, "credential", "", "Path Credential Not Found in ENV", fmt.Errorf("%s", "Path Credential Not Found in ENV"))
	}

	/* Prepare configuration credential file */
	jsonByte, err := os.ReadFile(credentialPath)
	if err != nil {
		infrastructure.Log.Panicln(nil, "credential", "", "Can't read file, Type:ReadFile", err)
	}

	var credential map[string]model.Credential
	if err := json.Unmarshal(jsonByte, &credential); err != nil {
		infrastructure.Log.Panicln(nil, "credential", "", "Can't unmarshal credential, Type:Unmarshal", err)
	}

	return credential
}
