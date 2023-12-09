package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	var themenum int = 1
	var messnum int = 1

	var themesst [1001]int
	var themesnm [1001]string
	var mess [1001]int

	var maxans int
	

	for i := 0; i != n; i++ {
		sc.Scan()
		tms, _ := strconv.Atoi(sc.Text())

		if tms == 0 {
			var name string

			sc.Scan()
			name = sc.Text()
			sc.Scan()

			themesnm[themenum] = name
			mess[messnum] = themenum
			themesst[themenum] = 1

			maxans = int(math.Max(float64(maxans), float64(themesst[themenum])))

			themenum++
			messnum++
		} else {
			linkms := tms
			sc.Scan()

			mess[messnum] = mess[linkms]
			themesst[mess[messnum]]++
			maxans = int(math.Max(float64(maxans), float64(themesst[mess[messnum]])))

			messnum++
		}
	}

	for i := 1; i != n + 1; i++ {
		if themesst[i] == maxans {
			fmt.Println(themesnm[i])
			break
		}
	}
}