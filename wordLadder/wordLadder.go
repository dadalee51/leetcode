package main

import (
        "fmt"
)

type choiceTree struct {
        s    string
        next []choiceTree
}

func main() {
        fmt.Println("hello")
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
        fmt.Println(beginWord, endWord, wordList)
        foundFirst := false
        for i, v := range wordList {
                if v == beginWord {
                        foundFirst = true
                }
        }
        //if begin or end not in the list then return empty list.
        if !foundFirst {
                return [][]string{}
        }

        var root choiceTree
        root.s = beginWord

        for len(wordList) > 0 {
                //perform bfs by constructing a tree.
                if HasOneDiff(wordList[i], root.s) {
                        append(root.next, wordList[i])
                        if i < len(wordList)-1 {
                                wordList = append(wordList[:i], wordList[i+1:])
                        } else {
                                wordList = wordList[:i-1]
                        }
                }

        }
        return [][]string{{""}}
}

//utility method
func HasOneDiff(a string, b string) bool {
        ai := 0
        for i, v := range a {
                if b[i] == byte(v) {
                        //do nothing
                } else {
                        ai++
                }
        }
        if ai == len(a)-1 {
                return true
        }

        return false
}