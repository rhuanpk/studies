package main

import "time"

var stopLoadingBar chan bool = make(chan bool)

func startLoadingBar1() {
	go func() {
		for {
			select {
			case <-stopLoadingBar:
				break
			default:
				// print("\r\033[KExecuting")
				print("\rExecuting")
				for index := 0; index < 3; index++ {
					time.Sleep(time.Millisecond * 100)
					print(".")
				}
				print(" ")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func startLoadingBar2() {
	go func() {
		items := []string{"|", "/", "-", "\\", "|", "/", "-", "\\"}
		for {
			select {
			case <-stopLoadingBar:
				break
			default:
				for _, item := range items {
					print("\rExecuting ", item, " ")
					time.Sleep(time.Millisecond * 75)
				}
			}
		}
	}()
}

func finishLoadingBar() {
	stopLoadingBar <- true
}

func main() {
	// startLoadingBar1()
	startLoadingBar2()
	time.Sleep(time.Second * 3)
	finishLoadingBar()
}
