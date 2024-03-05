package cmd

import (
	"github.com/spf13/cobra"
	"go_learn/database/seeders"
	"go_learn/pkg/console"
	"go_learn/pkg/seed"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to database",
	Run:   runSeeders,
	Args:  cobra.RangeArgs(0, 1),
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 有传参情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + seeder.Name)
		}
	} else {
		// 默认全部运行
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
