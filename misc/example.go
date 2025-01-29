package misc

import (
	"fmt"
	"os"

	markdown "github.com/Klaus-Tockloth/go-term-markdown"
)

func main() {
	path := "Readme.md"
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	result := markdown.Render(string(source), 80, 6)

	fmt.Println(result)
}
