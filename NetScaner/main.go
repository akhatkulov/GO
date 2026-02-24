package main 

import (
	"fmt"
	"net"
	"strconv"
	"sort"
	"time"
)

func worker(target string,ports,results chan int){
	for p:=range ports{
		address := target + ":" + strconv.Itoa(p);
		connect,err := net.DialTimeout("tcp",address,1*time.Second);
		if err != nil{
			results <- 0;
			continue;
		}
		connect.Close();
		results <- p;
	}
	
}

func main(){
	fmt.Println("Hello NetScaner");
	var target string;
	fmt.Print("Nishonni kirit>>> ");
	fmt.Scanln(&target);

	
	ports := make(chan int,100);
	results := make(chan int);

	var open_ports []int;

	for i := 0; i<cap(ports); i++{
		go worker(target,ports,results)
	}
	start := 1
	end := 56535
	totalTasks := end-start;
	go func(){
		for i:=start; i<=end; i++{
			ports <- i;
		}
	}()

	for i:=0; i <= totalTasks; i++{
		port := <-results;
		if port != 0{
			open_ports = append(open_ports,port);
		}
	}

	close(ports);
	close(results);

	sort.Ints(open_ports);
	for _,port :=range open_ports{
		fmt.Printf("-[Y]- ---> %d <---\n",port)
	}
	fmt.Println("Task finished!")
}
