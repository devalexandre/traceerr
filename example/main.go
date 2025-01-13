package main

import (
	"github.com/devalexandre/traceerr"
)

func funcA() error {
	traceerr.Print("Starting funcA")
	traceerr.Printf("Processing step %d in funcA", 1)
	return traceerr.Errorf("Something went wrong in funcA")
}

func funcB() error {
	traceerr.Print("Starting funcB")
	if err := funcA(); err != nil {
		return traceerr.Errorf("funcB encountered an error: %v", err)
	}
	return nil
}

func main() {
	err := funcB()
	if err != nil {
		traceerr.Errorln("Program completed with error:", err)
	} else {
		traceerr.Print("Program completed successfully")
	}
}
