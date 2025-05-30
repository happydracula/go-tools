package main
import ("fmt" 
 "time")
func hello(name string, c chan string){
	for i:=0;i<5;i++{
		c <- fmt.Sprintf("%s %d", name, i)
		time.Sleep(1 * time.Second)
	}
}
func temp() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	go hello("ani", c1)
	go hello("delly", c2)
	go func(){
		for i:=0;i<10;i++{
			select{
				case s := <- c1: c3 <- s
				case s := <- c2: c3 <- s
				}
		}
		
	}()
	for i:=0;i<10;i++{
		fmt.Println(<- c3)
	}
}
