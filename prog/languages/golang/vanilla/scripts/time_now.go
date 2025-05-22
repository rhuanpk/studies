package main

import "time"

func main() {

	println("retorna a data atual no formato yyyy-mm-dd")
	println(time.Now().Format("2006-01-02"))
	println("retorna a hora atual no formato am/pm")
	println(time.Now().Format("03:04:05"))

	println("-------------------------------------------")

	println("retorna a data atual no formato dd/mm/yyyy")
	println(time.Now().Format("02/01/2006"))
	println("retorna a hora atual no formato 24h")
	println(time.Now().Format("15:04:05"))

}
