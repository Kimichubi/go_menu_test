package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)



func main () {
 menu()
}

func menu() {
 var option int
 fmt.Println("Welcome to GLCalculator the most inteligent calculator in the world!")
 for option != 6 {
   fmt.Println("Please Select,\n 1: Calculate\n 2:Convert Celsius to Fahrenheit\n 3:Count Words\n 4:Guess Number \n 5:Timer \n 6:Exit")  
   fmt.Scan(&option)
   switch option {
   case 1:
   calculate()
   case 2:
    convert()
   case 3:
    countWords()
   case 4:
    guessNumb()
   case 5:
    timer()
    case 6:
    fmt.Println("Goodbye...")
   default:
    fmt.Println("Invalid option, please try again.")
   }
   
 }
}
func calculate()  {
   var firstNumber, secondNumber int
   fmt.Println("Please enter the first number: ")
   fmt.Scan(&firstNumber)
   fmt.Println("Please enter the second number: ")
   fmt.Scan(&secondNumber)
   sum := firstNumber + secondNumber
   fmt.Println("The sum is: ", sum)
}
func convert() {
var option, value int
for option != 3 {
fmt.Println("Please select the type of conversion: ")
fmt.Println("1:Celsius to Fahrenheit\n2:Fahrenheit to Celsius")
fmt.Scan(&option)
switch option {
case 1:
  fmt.Println("Press the number to convert: ")
  fmt.Scan(&value)
  fmt.Println("Celsius to Fahrenheit: ", (value * 9 / 5) + 32)
  option = 3 
case 2: 
fmt.Println("Press the number to convert: ")
fmt.Scan(&value)
fmt.Println("Fahrenheit to Celsius: ", (value - 32) * 5 / 9 )
  option = 3
default:
  fmt.Println("Invalid option...")
}
}
}
func countWords() {
 var word string
 fmt.Println("Please type the word: ")
 fmt.Scan(&word)
 fmt.Println("The word len is: ", len(word))  
}

func guessNumb() {
  var numb int
  randomNumb := rand.IntN(10)
  fmt.Println("Guess the number between 0 and 9")
  fmt.Scan(&numb)  
  if numb != randomNumb{
     for numb != randomNumb{
     fmt.Println("Try again!")
     fmt.Scan(&numb) 
     if numb == randomNumb{
      fmt.Println("Congratulations!!!!")
     }
    }

  } else {
    fmt.Println("Congratulations!!!!")
  }


}

func timer(){
  var timeToSpend int
  fmt.Println("Please type the time to spend: ")
  fmt.Scan(&timeToSpend)
  for timepassed := range timeToSpend{
   time.Sleep(1 * time.Second)
   fmt.Println(timepassed + 1," second have passed!")
  }
}