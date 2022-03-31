//4
package sport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Display struct {
	char_1 string
	char_2 int
	char_3 string
	char_4 []string
	char_5 float64
	char_6 int
}

func Sport() {

	viper.BindEnv("FORMAT")
	assign := viper.Get("FORMAT")
	fmt.Println("FORMAT =", assign)
	fmt.Println("Converting into yaml.....")
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
	}

	var Disp Display

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		dlmt_open := strings.Split(scanner.Text(), "[")
		dlmt_open1 := dlmt_open[0]
		dlmt_open2 := dlmt_open[1]
		com_split := strings.Split(dlmt_open1, ",")
		Disp.char_1 = com_split[0]

		Disp.char_3 = com_split[2]

		conv_char_2, err := strconv.Atoi(com_split[1])
		if err != nil {

			fmt.Println(err)

		}

		Disp.char_2 = conv_char_2
		delmt_close := strings.Split(dlmt_open2, "]")
		com_split_a := strings.Split(delmt_close[0], ",")
		Disp.char_4 = com_split_a
		com_split_b := strings.Split(delmt_close[1], ",")
		conv_char_6, _ := strconv.Atoi(com_split_b[2])

		Disp.char_6 = conv_char_6

		conv_char_5, _ := strconv.ParseFloat(com_split_b[1], 64)

		Disp.char_5 = float64(conv_char_5)
		//

		if assign == "yaml" {
			//fmt.Println(Disp.char_1, Disp.char_2, Disp.char_3, Disp.char_4, Disp.char_5, Disp.char_6)

			yamlData, _ := yaml.Marshal(Disp.char_1)
			fmt.Println(string(yamlData))

			file, _ := os.Create("o.yaml")

			op, _ := file.WriteString(string(yamlData))
			fmt.Println(op)

			yamlData1, _ := yaml.Marshal(Disp.char_2)
			fmt.Println(string(yamlData1))

			op1, _ := file.WriteString(string(yamlData1))
			fmt.Println(op1)

			yamlData2, _ := yaml.Marshal(Disp.char_3)
			fmt.Println(string(yamlData2))

			op2, _ := file.WriteString(string(yamlData2))
			fmt.Println(op2)

			yamlData3, _ := yaml.Marshal(Disp.char_4)
			fmt.Println(string(yamlData3))

			op3, _ := file.WriteString(string(yamlData3))
			fmt.Println(op3)

			yamlData4, _ := yaml.Marshal(Disp.char_5)
			fmt.Println(string(yamlData4))

			op4, _ := file.WriteString(string(yamlData4))
			fmt.Println(op4)

			yamlData5, _ := yaml.Marshal(Disp.char_6)
			fmt.Println(string(yamlData5))

			op5, _ := file.WriteString(string(yamlData5))
			fmt.Println(op5)
		} else {
			fmt.Println("Something went wrong")
		}

	}
}
