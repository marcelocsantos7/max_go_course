package maps

import (
	"fmt"
)

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon web services": "https://aws.com",
	} // almost lika struct
	fmt.Println(websites)
	fmt.Println(websites["google"])

	delete(websites, "google")

	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	fmt.Println(websites["linkedin"])
}

// ############################################################
