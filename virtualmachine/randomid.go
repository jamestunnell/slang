package virtualmachine

import (
	"log"
	"strings"

	"github.com/blueimp/passphrase"
)

func RandomID(numWords int) string {
	id, err := passphrase.String(numWords)
	if err != nil {
		log.Printf("failed to make random words: %v\n", err)
	}

	return strings.Replace(id, " ", "-", numWords-1)
}
