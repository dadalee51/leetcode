package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"regexp"
	"io"
)

func main() {
	
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Enter card number on each line (CRTL+C to QUIT) : ")
	for {
		raws, err := rd.ReadString('\n')
		if err == io.EOF { break }
		ValidateCardNum(raws)
	}
	fmt.Println("Program terminated.")
}

//Validate the card number
func ValidateCardNum(raws string) (cardType string, valid bool) {
	if raws == "" {return "Unknown", false}
	r, _ := regexp.Compile("[\r\n\t ]*")
	str := r.ReplaceAllString(raws, "")
	ct := GetCardType(str)
	if "Unknown" == ct {
		fmt.Printf("%-29s (invalid)\n", ct + ": " + str)
		return ct, false
	}
	sum := 0
	//there will be two cases, even or odd number of digits.
	if len(str) % 2 == 0 {
		//double every first digit
		for i, c := range str{
			x := int(c)-'0'
			if i % 2 == 0 {
				x *= 2
				if x >= 10 { x = x%10 + x/10 }
			}
			sum += x
		}
	}else{
		//double every second digit
		for i, c := range str{
			x := int(c)-'0'
			if i % 2 == 1 {
				x *= 2
				if x >= 10 { x = x%10 + x/10 }
			}
			sum += x
		}
	}
	if sum % 10 == 0 {
		fmt.Printf("%-29s (valid)\n", ct + ": " + str)
		return ct, true
	}else {
		fmt.Printf("%-29s (invalid)\n", ct + ": " + str)
		return ct, false
	}
}

/*
	get card type by prefix(p) and length(m).
*/
func GetCardType(s string) string{
	//check if all inputs are numbers
	for _, c:=range s{
		if int(c)-'0' > 10 || int(c) - '0' < 0{
			return "Unknown"
		}
	}
	m := len(s)
	if PrefixIn(s,[]string{"34", "37"}) && m == 15 { return "AMEX"
	} else if strings.HasPrefix(s,"6011") && m == 16 { return "Discover" 
	} else if PrefixIn(s,[]string{"51","52","53","54","55"}) && m == 16 { return "MasterCard" 
	} else if strings.HasPrefix(s,"4") && (m == 13 || m == 16) { return "VISA" 
	} else { return "Unknown" }

}

//utility method
func PrefixIn (s string, list [] string) bool {
	for _, a:= range list{
		if strings.HasPrefix(s, a){ return true }
	}
	return false
}