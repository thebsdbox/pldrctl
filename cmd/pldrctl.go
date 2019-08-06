package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Release - this struct contains the release information populated when building katbox
var Release struct {
	Version string
	Build   string
}

// PlunderServer - Contains all the details needed to interact with a server instance
var PlunderServer struct {
	URL *url.URL
}

var logLevel int
var urlString, username, password string

// TODO - thebsdbox(enable username/pass)
var disableauth bool

func init() {
	GetPlunderCmd.PersistentFlags().StringVar(&urlString, "url", os.Getenv("pURL"), "The Url of a plunder server")
	GetPlunderCmd.PersistentFlags().StringVar(&username, "user", os.Getenv("pUser"), "The Username for a plunder server")
	GetPlunderCmd.PersistentFlags().StringVar(&password, "pass", os.Getenv("pPass"), "The Password password a plunder server")

	pldrcltCmd.PersistentFlags().IntVar(&logLevel, "logLevel", int(log.InfoLevel), "Set the logging level [0=panic, 3=warning, 5=debug]")

	pldrcltCmd.AddCommand(GetPlunderCmd)
	pldrcltCmd.AddCommand(pldrcltVersion)
}


// Execute - starts the command parsing process
func Execute() {
	if os.Getenv("VCLOG") != "" {
		i, err := strconv.ParseInt(os.Getenv("VCLOG"), 10, 8)
		if err != nil {
			log.Fatalf("Error parsing environment variable [VCLOG")
		}
		// We've only parsed to an 8bit integer, however i is still a int64 so needs casting
		logLevel = int(i)
	}

	if err := pldrcltCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var pldrcltCmd = &cobra.Command{
	Use:   "pldrctl",
	Short: "VMware vCenter Text User Interface",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		cmd.Help()
		return
	},
}

//GetPlunderCmd - is used for it's subcommands for pulling data from a plunder server
var GetPlunderCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve data from a Plunder server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var pldrcltVersion = &cobra.Command{
	Use:   "version",
	Short: "Version and Release information about the plunder tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Plunder Release Information\n")
		fmt.Printf("Version:  %s\n", Release.Version)
		fmt.Printf("Build:    %s\n", Release.Build)
	},
}
