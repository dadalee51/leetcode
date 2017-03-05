package main

import (
 "testing"
 "fmt"
 "reflect"
)

type tTestCase struct{
	beg string
	end string
	list [] string
	expList [][] string
}

func TestFindLadders(t *testing.T){
	fmt.Println("test find ladder")
	a:=[]tTestCase{
		{"got", "het", []string{"got", "get", "hot", "hat", "het"}, [][]string{{"got", "get", "het"}}},
		{"aaa", "bbb", []string{"bbb", "aaa", "acb", "cab", "ccb", "aab", "abb" }, [][]string{{"aaa","aab","abb","bbb"}}},
	}
	fmt.Println(a)
	for _,v:=range a{
		if reflect.DeepEqual(findLadders(v.beg, v.end, v.list) ,v.expList){
			fmt.Println("test passed")
		}else{
			fmt.Println("test failed at", v)
			t.Fail()
		}
	}
}

func TestHasOneDiff(t *testing.T){
	fmt.Println("test has one diff")
	if !HasOneDiff("abc", "bbc"){ t.Fail() }
	
}