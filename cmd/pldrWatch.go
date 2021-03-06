package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"path"

	"github.com/plunder-app/pldrctl/pkg/ux"
	"github.com/plunder-app/plunder/pkg/apiserver"
	"github.com/plunder-app/plunder/pkg/plunderlogging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	// pldrctlGet.PersistentFlags().StringVar(&urlFlag, "url", os.Getenv("pURL"), "The Url of a plunder server")

	// getLogs.Flags().IntVarP(&watch, "watch", "w", 0, "Setting a watch timeout until \"completion\" or \"fail\"")
	// getUnLeased.Flags().BoolVarP(&colour, "colour", "c", false, "Use Colourful output")
	pldrctlWatch.AddCommand(pldrctlHTTP)

}

var pldrctlWatch = &cobra.Command{
	Use:   "watch",
	Short: "Watch events generated by the Plunder server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var pldrctlHTTP = &cobra.Command{
	Use:   "http",
	Short: "Watch events generated by the Plunder web server",
	Run: func(cmd *cobra.Command, args []string) {
		// Parse through the flags and attempt to build a correct URL
		log.SetLevel(log.Level(logLevel))

		u, c, err := apiserver.BuildEnvironmentFromConfig(pathFlag, urlFlag)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		c.Timeout = 0
		u.Path = path.Join(u.Path, "/logs/http/11-22-33-44-55")
		resp, err := c.Get(u.String())
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		var logs plunderlogging.JSONLogEntry
		logDecoder := json.NewDecoder(resp.Body)

		for {
			// Read one JSON object and store it in a map.
			if err := logDecoder.Decode(&logs); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			ux.LogsStreamFormat(logs)
		}
		fmt.Printf("?")
	},
}

// func watcher() {

// }
