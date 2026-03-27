package output

import (
	"fmt"
	"io"

	"github.com/spacescale/cli/internal/auth"
	"github.com/spacescale/cli/internal/config"
	"github.com/spacescale/cli/internal/version"
)

type Printer struct{}

func NewPrinter() Printer {
	return Printer{}
}

func (Printer) PrintRootStub(w io.Writer, cfg config.Config) error {
	_, err := fmt.Fprintf(w, "spacescale cli stub (env: %s)\n", cfg.Environment)
	return err
}

func (Printer) PrintVersion(w io.Writer) error {
	_, err := fmt.Fprintf(w, "spacescale version %s (commit: %s, date: %s)\n", version.Version, version.Commit, version.Date)
	return err
}

func (Printer) PrintLoginStub(w io.Writer, credentials auth.Credentials) error {
	if credentials.Token == "" {
		_, err := fmt.Fprintln(w, "login is not implemented yet. For now, set SPACESCALE_TOKEN in your environment.")
		return err
	}

	_, err := fmt.Fprintln(w, "login is not implemented yet. SPACESCALE_TOKEN is set and will be used by future commands.")
	return err
}
