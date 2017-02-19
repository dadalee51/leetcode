package main
import ( "fmt" )
func main(){
	r:= ""
	r=reverseVowels("hello")
	fmt.Println("result is", r, "should be: holle")

}

func reverseVowels(s string) string {
    a:=[]byte(s)
	i:=0
	j:=len(a)-1
	for i < j{
		if IsVowel(a[i]) { 
			for j > i{
				if IsVowel(a[j]) {
					fmt.Println("before", string(a), i, j)
					a[i],a[j] = a[j],a[i]
					fmt.Println("after", string(a), i, j)
					j--
					break
				}else{
					j--
				}
			}
		}
		i++
	}
	return string(a)
}

//utility func
func IsVowel(bi byte) bool {
	for _,bo:=range []byte("aeiouAEIOU") {
		if bo==bi {return true}
	}
	return false

}