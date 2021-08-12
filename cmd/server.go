package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var serverCmd = &cobra.Command{
	Use:   "server",
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

		// 我这里是要启动一个server

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
