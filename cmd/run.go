package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "run",
	Short: "run new app",
	Long:  `run new app`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) != 1 {
	// 		return errors.New("requires a color argument")
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(11111)
		// fileName := "posts/" + args[0]
		// _, err := os.Stat("posts")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// if os.IsNotExist(err) {
		// 	err := os.Mkdir("posts", 644)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		// _, err = os.Stat(fileName)
		// if os.IsNotExist(err) {
		// 	file, err := os.Create(fileName)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	log.Printf("create file %s", fileName)
		// 	defer file.Close()
		// } else {
		// 	log.Printf("file exists %s", fileName)
		// }
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
