package main

import (
	"testing"
)

//test case struct
type tCard struct{
	expType string //expected card type
	cardNum string //input card number
	expValid bool //expected validity
}

func TestValidateCardNum(t *testing.T){
	sample := []tCard{
		{"VISA", "4111111111111111", true},
		{"VISA", "4111111111111" ,false},
		{"VISA", "4012888888881881", true},
		{"AMEX", "378282246310005", true},
		{"Discover", "6011111111111117", true},
		{"MasterCard", "5105105105105100", true},
		{"MasterCard", "5105105105105106", false},
		{"Unknown", "9111111111111111", false},
	}
	
	for _, m := range sample{
		passing := true
		ct, b:= ValidateCardNum(m.cardNum)
		if b != m.expValid { 
			t.Error(m.cardNum, " should be valid:", m.expValid)
			passing = false
		}
		if m.expType != ct {
			t.Error(m.cardNum, " should be ", m.expType, " got ", ct , " instead.")
			passing = false
		}
		if !passing { t.Fail() }
	}
}
