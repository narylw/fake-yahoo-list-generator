package main
import (
    "fmt"
    "os"
	"math/rand"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
func main(){
	var str string;
	var num int;
		fmt.Println(`
 ____   _  _,             _,  _,  ____, ____,  _  _,  __,    _   _,
(-|__) (-\_/     ____,   (-|\ |  (-/_| (-|__) (-\_/  (-|    (-|  | 
 _|__)   _|,    (         _| \|, _/  |, _|  \,  _|,   _|__,  _|/\|,
(       (                (      (      (       (     (      (      
[ Instagram ] = @ Narylw
	  `)
	fmt.Println("Please enter the number of emails you want:")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		str = str
		 // + "fmailsabeli" + strconv.Itoa(i) + "@gmail.com\n";
		str += RandStringBytes(25) + "@yahoo.com\n";
	}
   file, errs := os.Create("Emails.txt")
   if errs != nil {
      fmt.Println("Failed to create file:", errs)
      return
   }
   defer file.Close()
   // Write the string "Hello, World!" to the file
   _, errs = file.WriteString(str)
   if errs != nil {
      fmt.Println("Failed to write to file:", errs) //print the failed message
      return
   }
   fmt.Println("Wrote to file 'Emails.txt'.") //print the success message
}