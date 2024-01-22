package main

import (
	"fmt"
	"sync"
	"time"
)

// Five philosophers dine together at the same table. Each philosopher has his own plate at the table.
// There is a fork between each plate. The dish served is a kind of spaghetti which has to be eaten with two forks.
// Each philosopher can only alternately think and eat. Moreover, a philosopher can only eat his spaghetti when he has both a left and right fork.
// Thus two forks will only be available when his two nearest neighbors are thinking, not eating. After an individual philosopher finishes eating,
// he will put down both forks. The problem is how to design a regimen (a concurrent algorithm) such that any philosopher will not starve; i.e.,
// each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or
// think (an issue of incomplete information).

//Struct store info about an individual philosopher
type Philosopher struct {
	name string
	rightFork int
	leftFork int
}

//List of philosophers
var philosophers = []Philosopher{
	{name: "Jean-Paul Sartre", leftFork: 4, rightFork: 0},
	{name: "Simone de Beauvoir", leftFork: 0, rightFork: 1},
	{name: "Albert Camus", leftFork: 1, rightFork: 2},
	{name: "Martin Heidegger", leftFork: 2, rightFork: 3},
	{name: "Søren Kierkegaard", leftFork: 3, rightFork: 4},
}


//Define variables
var hunger = 3 //how many time each philosopher eats until they are full
var eatTime = 1 * time.Second 
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second 

var orderMutex sync.Mutex 
var orderFinished []string

func main() {

	//Print welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The dining table is empty")

	time.Sleep(sleepTime)

	//Meal starts
	dine()
	
	//print end message
	fmt.Println("The dining table is empty")
}

func dine() {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// Map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// Meal starts
	for i := 0; i < len(philosophers); i++ {
		// goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// Seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()
	seated.Wait()

	// Eat 3 times
	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)	
		}
		

		fmt.Printf("\t%s has both forks and has started eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s has put down the forks.\n", philosopher.name)
	}
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")
	
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
