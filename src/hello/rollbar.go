package main
import (
  "github.com/rollbar/rollbar-go" 
  "time" 
  "fmt"
) 
func catchError() {  
    if r := recover(); r!= nil {
		rollbar.Critical(r)
        fmt.Println("recovered from ", r)
		rollbar.Wait()	
    }
}

func main() {
	defer catchError()
	rollbar.SetToken("69609226b4014ef39f62d4b2008c56ed")
	rollbar.SetEnvironment("production") 
	var timer *time.Timer = nil
	timer.Reset(10) // crash*/
	
}

