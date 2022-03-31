package sport

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

type Display_txtfile struct {
	char_1 string
	char_2 string
}

func Sport() {
	sli := make([]Display_txtfile, 0, 1)

	open_file, err := os.Open("s.txt")
	if err != nil {
		fmt.Println(err)
	}

	scan := bufio.NewScanner(open_file)
	for scan.Scan() {

		s := strings.Split(scan.Text(), " ")
		var x Display_txtfile
		x.char_1, x.char_2 = s[0], s[1]
		sli = append(sli, x)

	}

	for _, i := range sli {
		fmt.Println(i.char_1, i.char_2)
	}

}
