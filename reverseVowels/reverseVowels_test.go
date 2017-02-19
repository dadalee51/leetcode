package main
import ( "testing" ; "fmt")
func TestReverseVowels(t *testing.T){
	a:=[][]string {
		{"hello", "holle"},
		{"bus", "bus"},
		{"apple", "eppla"},
		{"eeellectriciteee", "eeellictriceteee"},
		{"aA", "Aa"},
	}
	for _,v:=range a{
		if !comapreResults(v[0],v[1]){
			t.Error(v[0],"is not equal to ", v[1])
			t.Fail()
		}else{
			fmt.Println("pass!", v[0], v[1])
		}
	}
}

func comapreResults(input string, result string) bool{
	if reverseVowels(input)==result{return true}
	return false
}

