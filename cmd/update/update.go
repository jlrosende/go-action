package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	RunE:  update,
}

func init() {

}

func update(cmd *cobra.Command, args []string) error {
	fmt.Println("Update")
	return nil
}
