package verify

import (
	"fmt"
	"os"
	"strings"

	"github.com/containrrr/shoutrrr/internal/util"
	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/router"

	"github.com/spf13/cobra"
)

// Cmd verifies the validity of a service url
var Cmd = &cobra.Command{
	Use:    "verify",
	Short:  "Verify the validity of a notification service URL",
	PreRun: util.MoveEnvVarToFlag,
	Run:    Run,
	Args:   cobra.MaximumNArgs(1),
}

var sr router.ServiceRouter

func init() {
	Cmd.Flags().StringP("url", "u", "", "The notification url")
}

// Run the verify command
func Run(cmd *cobra.Command, args []string) {
	URL, _ := cmd.Flags().GetString("url")
	sr = router.ServiceRouter{}

	service, err := sr.Locate(URL)

	if err != nil {
		fmt.Printf("error verifying URL: %s", err)
		os.Exit(1)
	}

	configMap, maxKeyLen := format.GetConfigMap(service)
	for key, value := range configMap {
		pad := strings.Repeat(" ", maxKeyLen-len(key))
		fmt.Printf("%s%s: %s\n", pad, key, value)
	}
}
