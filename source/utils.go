package main

import (
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