package cmd

import (
	"github.com/kainhuck/dod/core"
	"github.com/spf13/cobra"
	"log"
)

var view bool

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete docker containers or images",
	Long:  "delete docker containers or images",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			core.DeleteAll(view)
			return
		}
		obj := args[0]
		switch obj {
		case "image", "img":
			if len(args) > 1{
				for _, id := range args[1:]{
					core.DeleteImage(id, view)
				}
			}else{
				core.DeleteAllImage(view)
			}
		case "container", "con":
			if len(args) > 1{
				for _, id := range args[1:]{
					core.DeleteContainer(id, view)
				}
			}else{
				core.DeleteAllContainer(view)
			}
		default:
			log.Fatalf("unsupported object %v", obj)
		}
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&view, "view", "v", false, "show the process")
}