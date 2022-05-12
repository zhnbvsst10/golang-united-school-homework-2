package square

//square

import (
	//"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
//#yourTypeNameHere#
const SidesCircle = 0
const SidesSquare = 4
const SidesTriangle = 3

func CalcSquare(sideLen float64, sidesNum int8) float64 {

	var ans float64
	if sidesNum == 3 {
		ans = (sideLen * sideLen * math.Sqrt(3)) / 4
		//return ans
	} else if sidesNum == 4 {
		ans = sideLen * sideLen
		//return ans
	} else if sidesNum == 0 {
		ans = math.Pi * sideLen * sideLen / 4
		//return ans
	}
	return ans
}

//func main() {
//	fmt.Println(CalcSquare(10.0, SidesCircle))
//}
