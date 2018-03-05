package main
import (
  "github.com/rollbar/rollbar-go" 
  "errors"
  "fmt"
) 
func catchError() {  
    if r := recover(); r!= nil {
		rollbar.Critical(r)
        fmt.Println(r)
		rollbar.Wait()	
    }
}

func main() {
	defer catchError()
	rollbar.SetToken("69609226b4014ef39f62d4b2008c56ed")
	rollbar.SetEnvironment("production") 
	/*"time" 
	var timer *time.Timer = nil
	timer.Reset(10) // crash*/
	
	r := generateError("Test Exception")
    if r != nil {
        fmt.Println("recovered from ", r)
		rollbar.Critical(r)
        fmt.Println(r)
		rollbar.Wait()
        // Add Rollbar later here to report the error.
    }
	
}

func generateError(err string) error {

    return errors.New(err)
}



