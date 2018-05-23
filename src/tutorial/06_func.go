//ref: https://studygolang.com/articles/11892

package tutorial

import (
	"fmt"
)

func rectProp(width, height int) (int, int) {
	area := width * height
	perimeter := 2 * (width + height)
	return area, perimeter
}

func funcReturn() {
	area, perimeter := rectProp(3, 2)
	fmt.Printf("width=3,height=2, area is %v, perimeter is %v \n", area, perimeter)
}
