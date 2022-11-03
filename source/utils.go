package main

import (
	"fmt"
	"time"
	"strconv"
)
func SnowflakeTimestamp(ID string) (t time.Time, err error) {
	i, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return
	}
	timestamp := (i >> 22) + 1420070400000
	t = time.Unix(0, timestamp*1000000)
	return
}

func How_many_texts() int {
	var C int
	for i := 1; i <= len(Texts); i++ {
	    if Texts[i] != "" {
	    	C++
	    } else { return C+1 }
	}
	return 0
}

func ID_Text() {
	
}

func First_n(s string, n int) string {
    i := 0
    for j := range s {
        if i == n {
            return s[:j]
        }
        i++
    }
    return s
}

func String_to_binary(s string) (binString string) {
    for _, c := range s {
        binString = fmt.Sprintf("%s%b",binString, c)
    }
    return 
}