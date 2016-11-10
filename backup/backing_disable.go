// +build !qcow

package backup

import (
	"fmt"

	"github.com/openebs/longhorn/replica"
)

func openBackingFile(file string) (*replica.BackingFile, error) {
	if file == "" {
		return nil, nil
	}
	return nil, fmt.Errorf("Backing file is not supported")
}
