package main

import "fmt"
import "errors"
import "math"
func main() {
  ret, err  := calculateSqrt(-1)    

    if err != nil {
        fmt.Println(err)
    }else{
           fmt.Println("The square root is ", ret)
    }

}


func calculateSqrt(f float64) (float64, error) {
 //return an error as second parameter if invalid input
   if (f < 0) {
     return float64(math.NaN()),  errors.New("Not able to take the square root of a negative number")
   }
//otherwise use default square root function
  return math.Sqrt(f),nil
}
