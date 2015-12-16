// Go program to convert temperature from Fahrenhit to Celsius

package main
import "fmt"

func main(){
    fmt.Print("Enter the temperature: ")
    var tempF float64
    fmt.Scanf("%f", & tempF)
    fmt.Println("The entered temperature value is", tempF, "F")

    tempC := (tempF - 32) * 5/9
    fmt.Println("The temperature in Celcius is:", tempC, "C")
}
