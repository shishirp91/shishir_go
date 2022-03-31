package sport

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Display struct {
	char_1 string
	char_2 int
	char_3 string
	char_4 []string
	char_5 float64
	char_6 int
}

var (
	//WarningLogger *log.Logger
	//InfoLogger    *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)

	}
	//InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("logging done, please view logs.txt file for info")
}
func Sport() {
	file, err := os.Open("NX_file.txt")
	if err != nil {
		ErrorLogger.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		dlmt_open := strings.Split(scanner.Text(), "[")
		dlmt_open1 := dlmt_open[0]
		dlmt_open2 := dlmt_open[1]
		com_split := strings.Split(dlmt_open1, ",")
		var Disp Display
		Disp.char_1 = com_split[0]
		fmt.Printf(Disp.char_3)
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

		fmt.Println(Disp.char_1, Disp.char_2, Disp.char_3, Disp.char_4, Disp.char_5, Disp.char_6)
		//log.Println("Hello world!")
		//file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		//	if err != nil {
		//	log.Fatal(err)
		//}
		//log.Println("logging done")
		//log.SetOutput(file)

		//log.Println("logging done")
		// InfoLogger.Println("Starting the application...")
		// InfoLogger.Println("Something noteworthy happened")
		//WarningLogger.Println("There is something you should know about")
		//ErrorLogger.Println("Something went wrong")

	}

}
