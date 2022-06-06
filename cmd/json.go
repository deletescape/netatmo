package cmd

import (
	"fmt"
	"encoding/json"

	netatmo2 "github.com/mariusbreivik/netatmo/api/netatmo"
	"github.com/mariusbreivik/netatmo/internal/netatmo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// jsonCmd is the command for retrieving json data
var jsonCmd = &cobra.Command{
	Use:     "json",
	Short:   "read data from netatmo station as json",
	Long:    `read data from netatmo station as json`,
	Example: "netatmo json -i",
	RunE: func(cmd *cobra.Command, args []string) error {
		netatmoClient, err := netatmo.NewClient(netatmo.Config{
			ClientID:     viper.GetString("netatmo.clientID"),
			ClientSecret: viper.GetString("netatmo.clientSecret"),
			Username:     viper.GetString("netatmo.username"),
			Password:     viper.GetString("netatmo.password"),
		})

		if err != nil {
			return err
		}

		if indoor {
			printIndoorJson(netatmoClient.GetStationData())
		} else if outdoor {
			printOutdoorJson(netatmoClient.GetStationData())
		} else {
			fmt.Println(cmd.UsageString())
		}

		return nil
	},
}

func printIndoorJson(stationData netatmo2.StationData) {
	data, _ := json.Marshal(stationData.Body.Devices[0].DashboardData)
	fmt.Println(string(data))
}

func printOutdoorJson(stationData netatmo2.StationData) {
	data, _ := json.Marshal(stationData.Body.Devices[0].Modules[0].DashboardData)
	fmt.Println(string(data))
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().BoolVarP(&indoor, "indoor", "i", false, "netatmo json -i|--indoor")
	jsonCmd.Flags().BoolVarP(&outdoor, "outdoor", "o", false, "netatmo json -o|--outdoor")
}