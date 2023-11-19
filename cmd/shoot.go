package cmd

import (
	"fmt"
	"minaris/pkg/shooting"
	"minaris/pkg/utils/setting"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	shootingFlow string
	proxy        string
)

func init() {

	shootCmd.PersistentFlags().StringVarP(&shootingFlow, "shootingFlow", "f", "minaris.k4it0z11.yml", "Shooting flow file")
	shootCmd.MarkPersistentFlagRequired("shootingFlow")

	shootCmd.PersistentFlags().StringVarP(&proxy, "proxy", "p", "", "Shoot requests through a proxy")

	rootCmd.AddCommand(shootCmd)
}

var shootCmd = &cobra.Command{
	Use:   "shooter",
	Short: "Shooting multiple requests.",
	Long:  `Shooting multiple requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Printf("Shooting: %s\n", shootingFlow)
		if shootingFlow != "" {
			viper.SetConfigName(shootingFlow)
		}

		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			return err
		}

		var config shooting.Config

		viper.Unmarshal(&config)
		shooter := shooting.NewShooterWorkflow(&config)

		shootingConfig := &setting.ShootingConfig{}
		if proxy != "" {
			shootingConfig.Proxy = proxy
			fmt.Println("Proxy " + shootingConfig.Proxy)
		}

		shooter.UseConfig(shootingConfig)

		shooter.UseReader(new(shooting.RawReader))
		shooter.ShootWorkflow()
		return nil
	},
}
