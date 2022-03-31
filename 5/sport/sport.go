package sport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Display struct {
	Name      string
	Age       int
	Gender    string
	Sportname []string
	Height    float64
	Weight    int
}

func Sport() {
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
	}
	x := make([]Display, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		dlmt_open := strings.Split(scan.Text(), "[")
		dlmt_open1 := dlmt_open[0]
		dlmt_open2 := dlmt_open[1]
		com_split := strings.Split(dlmt_open1, ",")
		var Disp Display
		Disp.Name = com_split[0]
		fmt.Printf(Disp.Gender)
		Disp.Gender = com_split[2]

		conv_char_2, err := strconv.Atoi(com_split[1])
		if err != nil {

			fmt.Println(err)

		}

		Disp.Age = conv_char_2
		delmt_close := strings.Split(dlmt_open2, "]")
		com_split_a := strings.Split(delmt_close[0], ",")
		Disp.Sportname = com_split_a
		com_split_b := strings.Split(delmt_close[1], ",")
		conv_char_6, _ := strconv.Atoi(com_split_b[2])

		Disp.Weight = conv_char_6

		conv_char_5, _ := strconv.ParseFloat(com_split_b[1], 64)

		Disp.Height = float64(conv_char_5)
		x = append(x, Disp)
	}
	//fmt.Println(Disp.char_1, Disp.char_2, Disp.char_3, Disp.char_4, Disp.char_5, Disp.char_6)

	//yamlData, err := yaml.Marshal(Disp.char_1)
	//fmt.Println(string(yamlData))

	//op, err := file.WriteString(string(yamlData))
	//fmt.Println(op)

	//yamlData1, err := yaml.Marshal(Disp.char_2)
	//fmt.Println(string(yamlData1))

	//op1, err := file.WriteString(string(yamlData1))
	//fmt.Println(op1)

	//yamlData2, err := yaml.Marshal(Disp.char_3)
	//fmt.Println(string(yamlData2))

	//op2, err := file.WriteString(string(yamlData2))
	//fmt.Println(op2)

	//yamlData3, err := yaml.Marshal(Disp.char_4)
	//fmt.Println(string(yamlData3))

	//op3, err := file.WriteString(string(yamlData3))
	//fmt.Println(op3)

	//yamlData4, err := yaml.Marshal(Disp.char_5)
	//fmt.Println(string(yamlData4))

	//op4, err := file.WriteString(string(yamlData4))
	//fmt.Println(op4)

	//yamlData5, err := yaml.Marshal(Disp.char_6)
	//fmt.Println(string(yamlData5))

	//op5, err := file.WriteString(string(yamlData5))
	//fmt.Println(op5)
	fmt.Println(x[0])

	file, err = os.Create("o.yaml")

	fmt.Println(x)
	y, err := yaml.Marshal(&x)
	fmt.Println(string(y))
	op, err := file.WriteString(string(y))
	fmt.Println(op)
}
