package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/acrosslite"
)

// WriteCopyright writes the <COPYRIGHT> entry
func WriteCopyright(pal *al.AcrossLite) string {
	const TAG = `<COPYRIGHT>`
	return fmt.Sprintf("%s\n    %s", TAG, pal.GetCopyright())
}
