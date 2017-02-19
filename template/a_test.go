package main
import ( "testing" )

type anyType struct{
	abc string
	cool string
	result bool
}

func TestSomeFunc(t *testing.T){

	

	fmt.Println("main func running.")
	if someFunc("abc") == "abc" { 
		fmt.Println("pass")
	}else{
		t.Fail()
	}
}