package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lychee",
	Short: "LYCHEE MICROSERVICES MANAGER\n\n- Starts each micro-xxxxx except micro-federation\n- Then starts micro-federation which will federate all microservices",
}

// // start the gateway executable
// func main() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }
