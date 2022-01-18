package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main(){
	if len(os.Args) == 1{
		fmt.Println("usage:")
		fmt.Println("\tDoge-Obf.exe string")
		fmt.Println("\tDoge-Obf.exe string hi")
		fmt.Println("\tDoge-Obf.exe string 8\t//0-8 Num")
		fmt.Println("\tDoge-Obf.exe str.txt hi")
		return
	}
	if os.Args[1] == "-h"{
		fmt.Println("usage:")
		fmt.Println("\tDoge-Obf.exe string")
		fmt.Println("\tDoge-Obf.exe string hi")
		fmt.Println("\tDoge-Obf.exe string 8\t//0-8 Num")
		fmt.Println("\tDoge-Obf.exe str.txt hi")
		return
	}

	var obfStr []string
	if strings.Contains(os.Args[1],".txt"){
		obfStr,_ = readLines(os.Args[1])
		if obfStr == nil{
			obfStr = append(obfStr,os.Args[1])
		}
	}else{
		obfStr = append(obfStr,os.Args[1])
	}

	for _,val := range obfStr{
		b := Str2byte(val)
		idx := GenerateRangeNum(0,8)
		if len(os.Args) == 3 {
			if os.Args[2] == "hi"{
				idx = GenerateRangeNum(4,8)
			}
			i0,err := strconv.Atoi(os.Args[2])
			if err == nil{
				idx = i0
			}
		}
		switch idx{
		case 0:
			b.tostring1()
		case 1:
			b.tostring2(false)
		case 2:
			b.tostring3(false)
		case 3:
			b.tostring4(false)
		case 4:
			b.tostring3(true)
		case 5:
			b.tostring4(true)
		case 6:
			b.tostring2(true)
		case 7:
			b.tostring5()
		case 8:
			b.tostring6()
		}
	}

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


type strByte struct{
	rawstr string
	str []string
	bytes []byte
}

func Str2byte(str0 string)*strByte{
	var s0 []string
	b0 := []byte(str0)
	for _,v := range b0{
		s0 = append(s0,string(v))
	}
	return &strByte{rawstr:str0, str:s0, bytes: b0}
}

func (sb *strByte)tostring1(){
	fmt.Println("//"+sb.rawstr)

	fmt.Printf("string([]byte{")
	for i := 0;i < len(sb.str)-1;i++{
		fmt.Printf("'%s',",sb.str[i])
	}
	fmt.Printf("'%s'",sb.str[len(sb.str)-1])
	fmt.Println("})\n")
}

func (sb *strByte)tostring2(offset bool){
	if offset == false{
		fmt.Println("//"+sb.rawstr)

		fmt.Printf("string([]byte{")
		for i := 0;i < len(sb.bytes)-1;i++{
			fmt.Printf("0x%x, ",sb.bytes[i])
		}
		fmt.Printf("0x%x",sb.bytes[len(sb.bytes)-1])
		fmt.Println("})\n")
	}else{
		fmt.Println("//"+sb.rawstr)

		fmt.Printf("string([]byte{")
		for i := 0;i < len(sb.bytes)-1;i++{
			rand0 := GenerateRangeNum(-16,16)
			if rand0 >0{
				fmt.Printf("0x%x-%d,",sb.bytes[i]+byte(rand0),rand0)
			}else if rand0 <0{
				fmt.Printf("0x%x+%d,",sb.bytes[i]+byte(rand0),-rand0)
			}else{
				fmt.Printf("0x%x,",sb.bytes[i],rand0)
			}

		}
		rand0 := GenerateRangeNum(-16,16)
		if rand0 >0{
			fmt.Printf("0x%x-%d",sb.bytes[len(sb.bytes)-1]+byte(rand0),rand0)
		}else if rand0 <0{
			fmt.Printf("0x%x+%d",sb.bytes[len(sb.bytes)-1]+byte(rand0),-rand0)
		}else{
			fmt.Printf("0x%x",sb.bytes[len(sb.bytes)-1],rand0)
		}
		fmt.Println("})\n")
	}
}

func (sb *strByte)tostring3(offset bool){
	if offset == false {

		fmt.Println("//" + sb.rawstr)

		fmt.Printf("string(append([]byte{}")
		for i := 0; i < len(sb.bytes); i++ {
			fmt.Printf(", byte(0x%x)", sb.bytes[i])
		}
		fmt.Println("))\n")
	}else {

		fmt.Println("//" + sb.rawstr)

		fmt.Printf("string(append([]byte{}")
		for i := 0; i < len(sb.bytes); i++ {
			rand0 := GenerateRangeNum(-16,16)
			if rand0 == 0{
				fmt.Printf(", byte(0x%x)", sb.bytes[i])
			}else if rand0 < 0{
				fmt.Printf(", byte(0x%x+%d)", sb.bytes[i]+byte(rand0),-rand0)
			}else{
				fmt.Printf(", byte(0x%x-%d)", sb.bytes[i]+byte(rand0),rand0)
			}
		}
		fmt.Println("))\n")
	}
}


func (sb *strByte)tostring4(offset bool){
	if offset == false{
		fmt.Println("//"+sb.rawstr)
		fmt.Printf("string(append([]byte{}")
		for i := 0;i < len(sb.bytes);i++{
			fmt.Printf(", []byte{0x%x}[0]",sb.bytes[i])
		}
		fmt.Println("))\n")
	} else{
		fmt.Println("//"+sb.rawstr)
		fmt.Printf("string(append([]byte{}")
		for i := 0;i < len(sb.bytes);i++{
			rand0 := GenerateRangeNum(-16,16)
			if rand0 == 0{
				fmt.Printf(", []byte{0x%x}[0]",sb.bytes[i])
			}else if rand0 < 0{
				fmt.Printf(", []byte{0x%x+%d}[0]",sb.bytes[i]+byte(rand0),-rand0)
			}else{
				fmt.Printf(", []byte{0x%x-%d}[0]",sb.bytes[i]+byte(rand0),rand0)
			}
		}
		fmt.Println("))\n")
	}
}

func (sb *strByte)tostring5(){
	randArr := RandStringRunes(6)
	fmt.Println("//"+sb.rawstr)
	fmt.Println("var "+ randArr + " []byte")
	flag := 0
	for i := 0;i < len(sb.bytes);i++{
		flag ++
		if flag == 2{
			rand0 := GenerateRangeNum(1,len(sb.bytes))
			rand1 := GenerateRangeNum(1,len(sb.bytes))
			tmp_byte := fmt.Sprintf("%s = append(%s, []byte{",randArr,randArr)
			for j := 0;j <rand0;j++{
				tmp_byte = tmp_byte + fmt.Sprintf("0x%x,",[]byte(RandStringRunes(1))[0])
			}

			tmp_byte = tmp_byte+fmt.Sprintf("0x%x,0x%x,",sb.bytes[i-1],sb.bytes[i])

			for j := 0;j <rand1;j++{
				tmp_byte = tmp_byte + fmt.Sprintf("0x%x,",[]byte(RandStringRunes(1))[0])
			}
			tmp_byte = tmp_byte + fmt.Sprintf("}[%d:%d]...)",rand0,rand0+2)
			fmt.Println(tmp_byte)
			flag = 0
		}
	}
	if flag == 1{
		rand0 := GenerateRangeNum(1,len(sb.bytes))
		rand1 := GenerateRangeNum(1,len(sb.bytes))
		tmp_byte := fmt.Sprintf("%s = append(%s, []byte{",randArr,randArr)
		for j := 0;j <rand0;j++{
			tmp_byte = tmp_byte + fmt.Sprintf("0x%x,",[]byte(RandStringRunes(1))[0])
		}
		tmp_byte = tmp_byte+fmt.Sprintf("0x%x,",sb.bytes[len(sb.bytes)-1])
		for j := 0;j <rand1;j++{
			tmp_byte = tmp_byte + fmt.Sprintf("0x%x,",[]byte(RandStringRunes(1))[0])
		}
		tmp_byte = tmp_byte + fmt.Sprintf("}[%d:%d]...)",rand0,rand0+1)
		fmt.Println(tmp_byte)
	}
	fmt.Println("string("+randArr+")\n")
}

func (sb *strByte)tostring6(){
	randArr := RandStringRunes(6)
	fmt.Println("//"+sb.rawstr)
	fmt.Println("var "+ randArr + " []byte")
	flag := 0
	for i := 0;i < len(sb.bytes);i++{
		flag ++
		if flag == 2{
			rand0 := GenerateRangeNum(1,len(sb.bytes))
			rand1 := GenerateRangeNum(1,len(sb.bytes))
			tmp_byte := fmt.Sprintf("%s = append(%s, []byte{",randArr,randArr)
			for j := 0;j <rand0;j++{
				tmp_byte = tmp_byte + fmt.Sprintf("'%s',",RandStringRunes(1))
			}

			tmp_byte = tmp_byte+fmt.Sprintf("'%s','%s',",sb.str[i-1],sb.str[i])

			for j := 0;j <rand1;j++{
				tmp_byte = tmp_byte + fmt.Sprintf("'%s',",RandStringRunes(1))
			}
			tmp_byte = tmp_byte + fmt.Sprintf("}[%d:%d]...)",rand0,rand0+2)
			fmt.Println(tmp_byte)
			flag = 0
		}
	}
	if flag == 1{
		rand0 := GenerateRangeNum(1,len(sb.bytes))
		rand1 := GenerateRangeNum(1,len(sb.bytes))
		tmp_byte := fmt.Sprintf("%s = append(%s, []byte{",randArr,randArr)
		for j := 0;j <rand0;j++{
			tmp_byte = tmp_byte + fmt.Sprintf("'%s',",RandStringRunes(1))
		}
		tmp_byte = tmp_byte+fmt.Sprintf("'%s',",sb.str[len(sb.str)-1])
		for j := 0;j <rand1;j++{
			tmp_byte = tmp_byte + fmt.Sprintf("'%s',",RandStringRunes(1))
		}
		tmp_byte = tmp_byte + fmt.Sprintf("}[%d:%d]...)",rand0,rand0+1)
		fmt.Println(tmp_byte)
	}
	fmt.Println("string("+randArr+")\n")
}

func GenerateRangeNum(min, max int) int {
	randNum := rand.Intn(max - min) + min
	if randNum == 0{
		randNum++
	}
	return randNum
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
