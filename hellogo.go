package main

import (
	"fmt"
)

var pl = fmt.Println

// func sayHello() {
// 	pl("Hello!")
// }

// func getSum(x int, y int) int {
// 	return x + y
// }

// func getTwo(x int) (int, int) {
// 	return x + 1, x + 2
// }

// func getQuotient(x float64, y float64) (ans float64, err error) {
// 	if y == 0 {
// 		return 0, fmt.Errorf("You can't divide by zero")
// 	} else {
// 		return x / y, nil
// 	}
// }

// func getSum2(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum
// }

// func getArraySum(arr []int) int {
// 	sum := 0
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	return sum
// }

// func changeVal(f3 int) int {
// 	f3 += 1
// 	return f3
// }

// func changeVal2(myPtr *int) {
// 	*myPtr = 12
// }

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var NumSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / NumSize)
}

func main() {
	// pl("What is your name!")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	pl("Hello", name)
	// } else {
	// 	log.Fatal(err)
	// }

	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, "", '\u0000'

	// cV1 := 1.5
	// cV2 := int(cV1)
	// pl(cV2)

	// cV3 := "5000000"
	// cV4, err := strconv.Atoi(cV3)
	// pl(cV4, err, reflect.TypeOf(cV4))

	// cV5 := 5000000
	// cV6 := strconv.Itoa(cV5)
	// pl(cV6)

	// cV7 := "3.1415926"
	// if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
	// 	pl(cV8)
	// }
	// cV9 := fmt.Sprintf("%f", 3.1415926)
	// pl(cV9)

	// iAge := 8
	// if (iAge >= 1) && (iAge <= 18) {
	// 	pl("Important Birthday")
	// } else if (iAge == 21) || (iAge == 50) {
	// 	pl("Important Birthday")
	// } else if iAge >= 65 {
	// 	pl("Important Birthday")
	// } else {
	// 	pl("Not an important birthday")
	// }

	// pl("!true =", !true)

	// sV1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// sV2 := replacer.Replace(sV1)
	// pl(sV2)
	// pl("Length of string :", len(sV2))
	// pl("Contains Another :", strings.Contains(sV2, "Another"))
	// pl("o index :", strings.Index(sV2, "o"))
	// pl("Replace :", strings.Replace(sV2, "o", "0", -1))
	// sV3 := "\nSome Words\n" // \t \" \\
	// sV3 = strings.TrimSpace(sV3)
	// pl("Split :", strings.Split("a-b-c-d", "-"))
	// pl("Lower :", strings.ToLower(sV2))
	// pl("Upper :", strings.ToUpper(sV2))
	// pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	// pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// rStr := "abcdefg"
	// pl("Rune Count :", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr {
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	// pl("5 + 4 =", 5+4)
	// pl("5 - 4 =", 5-4)
	// pl("5 * 4 =", 5*4)
	// pl("5 / 4 =", 5/4)
	// pl("5 % 4 =", 5%4)

	// mInt := 1
	// mInt++
	// pl("mInt++ =", mInt)

	//pl("Float Precision =", 3.14159265358979323846264338327950288419716939937510582097494459)

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs) // rand.seed is depricated must use newSeed(Seed(seedSecs)) instead
	// randNum := rand.Intn(50) + 1
	// pl("Random :", randNum)

	// There are many math functions
	// pl("Abs(-10) =", math.Abs(-10))
	// pl("Pow(4, 2) =", math.Pow(4, 2))
	// pl("Sqrt(16) =", math.Sqrt(16))
	// pl("Cbrt(8) =", math.Cbrt(8))
	// pl("Ceil(9.25) =", math.Ceil(9.25))
	// pl("Floor(9.25) =", math.Floor(9.25))
	// pl("Round(9.25) =", math.Round(9.25))
	// pl("Log2(8) =", math.Log2(8))
	// pl("Log10(100) =", math.Log10(100))
	// pl("Mod(5, 3) =", math.Mod(5, 3))
	// Get the log of e to the power of 2
	// pl("Log(7.389) =", math.Log(math.Exp(2)))
	// pl("Max(5,4) =", math.Max(5, 4))
	// pl("Min(5,4) =", math.Min(5, 4))
	// Convert 90 degrees to radians
	// r90 := 90 * math.Pi / 180
	// d90 := r90 * 180 / math.Pi
	// fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	// pl("Sin(r90) =", math.Sin(r90))
	// There also functions for Cos, Tan, Acos, Asin, Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Tanh
	// Htpot
	// for x := 1; x <= 5; x++ {
	// 	pl(x)
	// }
	// fX := 0
	// for fX < 5 {
	// 	pl(fX)
	// 	fX++
	// }

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1
	// for true {
	// 	fmt.Print("Guess a number between 0 and 50: ")
	// 	pl("Random Number is :", randNum)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		pl("Pick a Lower Value")
	// 	} else if iGuess < randNum {
	// 		pl("Pick a Higher Value")
	// 	} else {
	// 		pl("You Guessed Correctly")
	// 		break
	// 	}
	// }

	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }

	// var arr1 [5]int
	// arr1[0] = 1
	// arr2 := [5]int{1, 2, 3, 4, 5}
	// pl("Index 0 :", len(arr2))
	// pl("Arr Length :", len(arr2))
	// for i := 0; i < len(arr2); i++ {
	// 	pl(arr2[i])
	// }
	// for i, v := range arr2 {
	// 	fmt.Printf("%d : %d\n", i, v)
	// }
	// arr3 := [2][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		pl(arr3[i][j])
	// 	}
	// }

	// aStr1 := "abcde"
	// rArr := []rune(aStr1)
	// for _, v := range rArr {
	// 	fmt.Printf("Rune Array : %d\n", v)
	// }
	// byteArr := []byte{'a', 'b', 'c'}
	// bStr := string(byteArr[:])
	// pl("I'm a string :", bStr)

	// sl1 := make([]string, 6)
	// sl1[0] = "Society"
	// sl1[1] = "of"
	// sl1[2] = "the"
	// sl1[3] = "Simulated"
	// sl1[4] = "Universe"
	// pl("Slice Size :", len(sl1))
	// for i := 0; i < len(sl1); i++ {
	// 	pl(sl1[i])
	// }
	// for _, x := range sl1 {
	// 	pl(x)
	// }
	// sArr := [5]int{1, 2, 3, 4, 5}
	// sl3 := sArr[0:2]
	// pl("1st 3 :", sArr[:3])
	// pl("Last 3 :", sArr[2:])
	// sArr[0] = 10
	// pl("s13 :", sl3)
	// sl3[0] = 1
	// pl("sArr :", sArr)

	// sl3 = append(sl3, 12)
	// pl("sl3 :", sl3)
	// pl("sArr :", sArr)

	// sl4 := make([]string, 6)
	// pl("sl4 :", sl4)
	// pl("sl4[0] :", sl4[0])

	// func funcName(parameters) returnType {BODY}
	// sayHello()
	// pl(getSum(5, 4))
	// pl(getTwo(5))
	// pl(getQuotient(5, 4))
	// pl(getSum2(1, 2, 3, 4, 5))
	// vArr := []int{1, 2, 3, 4}
	// pl("Array Sum :", getArraySum(vArr))
	// f3 := 5
	// pl("f3 before func :", f3)
	// changeVal(f3)
	// pl("f3 after func :", f3)
	// f4 := 10
	// pl("f4 before func :", f4)
	// changeVal2(&f4)
	// pl("f4 after func", f4)

	// var f4Ptr *int = &f4
	// pl("f4 Address :", f4Ptr)
	// pl("f4 Vallue :", *f4Ptr)
	// *f4Ptr = 11
	// pl("f4 Vallue :", *f4Ptr)
	// pl("f4 before func :", f4)
	// changeVal2(&f4)
	// pl("f4 after func", f4)

	// pArr := [4]int{1, 2, 3, 4}
	// dblArrVals(&pArr)
	// pl(pArr)

	// iSlice := []float64{11, 13, 17}
	// fmt.Printf("Average : %.3f\n", getAverage(iSlice...))

	// f, err := os.Create("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// iPrimeArr := []int{2, 3, 5, 7, 11}
	// var sPrimeArr []string
	// for _, i := range iPrimeArr {
	// 	sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	// }
	// for _, num := range sPrimeArr {
	// 	_, err := f.WriteString(num + "\n")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// f, err = os.Open("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// scan1 := bufio.NewScanner(f)
	// for scan1.Scan() {
	// 	pl("Prime :", scan1.Text())
	// }
	// if err := scan1.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	/*
		Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY : open the file read-only.
		O_WRONLY : open the file write-only.
		O_RDWR : open the file read-write.

		These can be or'ed with the following:

		O_APPEND : append data to the file when writing.
		O_CREATE : create a new file if none exists.
		O_EXCL : used with O_CREATE, file must not exist.
		O_SYNC : open for synchronous I/O.

	*/

	// _, err := os.Stat("data.txt")
	// if errors.Is(err, os.ErrNotExist) {
	// 	pl("File does not exist")
	// } else {
	// 	pl("File exists")
	// 	f, err := os.OpenFile("data.txt",
	// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()
	// 	if _, err := f.WriteString("13\n"); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

}
