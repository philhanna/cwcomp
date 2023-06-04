package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteCopyright writes the <COPYRIGHT> entry
func WriteCopyright(pal *al.AcrossLite) string {
	const TAG = `<COPYRIGHT>`
	return fmt.Sprintf("%s\n\t%s", TAG, pal.GetCopyright())
}
