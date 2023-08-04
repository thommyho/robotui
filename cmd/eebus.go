package cmd

import (
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	"github.com/thommyho/robotui/charger/eebus"
)

// teslaCmd represents the vehicle command
var eebusCmd = &cobra.Command{
	Use:   "eebus-cert",
	Short: "Generate EEBUS certificate for using EEBUS compatible chargers",
	Run:   runEEBUSCert,
}

func init() {
	rootCmd.AddCommand(eebusCmd)
}

const tmpl = `
Add the following to the robotui config file:

eebus:
  certificate:
    public: |
{{ .public | indent 6 }}
    private: |
{{ .private | indent 6 }}
`

func generateEEBUSCert() {
	cert, err := eebus.CreateCertificate()
	if err != nil {
		log.FATAL.Fatal("could not create certificate", err)
	}

	pubKey, privKey, err := eebus.GetX509KeyPair(cert)
	if err != nil {
		log.FATAL.Fatal("could not process generated certificate", err)
	}

	t := template.Must(template.New("out").Funcs(sprig.TxtFuncMap()).Parse(tmpl))
	if err := t.Execute(os.Stdout, map[string]interface{}{
		"public":  pubKey,
		"private": privKey,
	}); err != nil {
		log.FATAL.Fatal("rendering failed", err)
	}
}

func runEEBUSCert(cmd *cobra.Command, args []string) {
	generateEEBUSCert()
}
