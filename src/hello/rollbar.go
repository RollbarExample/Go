package main
import (
  "github.com/rollbar/rollbar-go" 
  "fmt"
  "time" 
) 
func recoverError() {  
	if r := recover(); r!= nil {
		rollbar.Critical(r)
		fmt.Println(r)
		rollbar.Wait()	
	}
}

func main() {
	defer recoverError()
	rollbar.SetToken("69609226b4014ef39f62d4b2008c56ed")
	rollbar.SetEnvironment("production") 
	var timer *time.Timer = nil
	timer.Reset(10) // crash
}



