package output

import (
	"fmt"
	"io"

	"github.com/spacescale/cli/internal/config"
)

type Printer struct{}

func NewPrinter() Printer {
	return Printer{}
}

func (Printer) PrintRootStub(w io.Writer, cfg config.Config) error {
	_, err := fmt.Fprintf(w, "spacescale cli stub (env: %s)\n", cfg.Environment)
	return err
}
