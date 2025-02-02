package cmd

import (
	"fmt"
	"log"

	"github.com/siesto1elemento/Aws-cli-manager/internal/aws"

	"github.com/siesto1elemento/Aws-cli-manager/internal/state"

	"github.com/spf13/cobra"
)

var stateFile string

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision resources from a state file",
	Run: func(cmd *cobra.Command, args []string) {
		stateData, err := state.LoadState(stateFile)
		if err != nil {
			log.Fatal("Failed to load state:", err)
		}

		for id, resource := range stateData {
			fmt.Println(stateData)
			if resource.Status == "provisioned" {
				fmt.Printf("Resource %s already provisioned.\n", id)
				continue
			}

			switch resource.Type {
			case "ec2":
				fmt.Println(resource)
				err = aws.ProvisionEC2(&resource)
			default:
				fmt.Printf("Unknown resource type: %s\n", resource.Type)
			}

			if err != nil {
				fmt.Printf("Failed to provision %s: %v\n", id, err)
				continue
			}
			resource.Status = "provisioned"
			stateData[id] = resource
		}

		state.SaveState(stateFile, stateData)
	},
}

func init() {
	provisionCmd.Flags().StringVarP(&stateFile, "state", "s", "state.json", "Path to the state file")
	rootCmd.AddCommand(provisionCmd)
}
