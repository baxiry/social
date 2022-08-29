package helps

import "fmt"

func Check(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
