package jobs

import (
	"fmt"
	"time"

	"github.com/Brianllp/go_practice/controllers"
)

func GetContentfulEntries(name string) {
	for {
		time.Sleep(5 * time.Minute)
		fmt.Printf("%s processing...\n", name)
		controllers.GetContentfulEntries()
	}
}
