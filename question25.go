//pattern printing

package main

import "fmt"

func main() {

	for i := 1; i <= 12; i++ {
		for k := 1; k <= 1; k++ {
			if (i == 2) || (i == 3) || (i == 4) || (i == 6) || (i == 7) || (i == 8) || (i == 9) || (i == 10) || (i == 11) {
				fmt.Print("*")
			}else if (i==1){
				fmt.Print("**")
			}else if (i==5){
				fmt.Print("****")
			}else if (i==12){
				fmt.Print("******")
			}

		}
		fmt.Println()
	}

}
