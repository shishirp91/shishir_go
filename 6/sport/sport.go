package sport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type disp struct {
	Char_1 string
	Char_2 int
	Char_3 string
	Char_4 []string
	Char_5 float64
	Char_6 int
}

type DetailWriter interface {
	WriteToFile(disp) error
}

type DetailWriterJson struct {
	//
}
type DetailWriterYaml struct {
	//
}

func WriteToFile(f string) []disp {
	x := make([]disp, 0)
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		dlmt_open := strings.Split(scanner.Text(), "[")
		dlmt_open1 := dlmt_open[0]
		dlmt_open2 := dlmt_open[1]
		com_split := strings.Split(dlmt_open1, ",")
		var Disp disp
		Disp.Char_1 = com_split[0]
		fmt.Printf(Disp.Char_3)
		Disp.Char_3 = com_split[2]

		conv_char_2, err := strconv.Atoi(com_split[1])
		if err != nil {

			fmt.Println(err)

		}

		Disp.Char_2 = conv_char_2
		delmt_close := strings.Split(dlmt_open2, "]")
		com_split_a := strings.Split(delmt_close[0], ",")
		Disp.Char_4 = com_split_a
		com_split_b := strings.Split(delmt_close[1], ",")
		conv_char_6, _ := strconv.Atoi(com_split_b[2])

		Disp.Char_6 = conv_char_6

		conv_char_5, _ := strconv.ParseFloat(com_split_b[1], 64)

		Disp.Char_5 = float64(conv_char_5)

		x = append(x, Disp)
		file, err = os.Create("o.yaml")
		y, err := yaml.Marshal(&x)
		op, err := file.WriteString(string(y))
		fmt.Println(op)

	}
	fmt.Println(x)
	return nil
}

func Sport() {

	var f = "in.txt"
	WriteToFile(f)

}
