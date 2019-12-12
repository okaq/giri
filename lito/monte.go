// monte hall problem
package main

import (
	"fmt"
	"math/rand"
)

type Counter struct {
	Run int
	SwitchWin int
	SwitchLoss int
	Win int
	Loss int
}

func NewCounter() *Counter {
	c0 := new(Counter)
	c0.Run = 0
	c0.SwitchWin = 0
	c0.SwitchLoss = 0
	c0.Win = 0
	c0.Loss = 0
	return c0
}

func (c *Counter) Tick(r *Round) {
	if r.Switch == true && r.Win == false {
		c.SwitchWin = c.SwitchWin + 1
	}
	if r.Switch == true && r.Win == true {
		c.SwitchLoss = c.SwitchLoss + 1
	}
	if r.Switch == false && r.Win == false {
		c.Loss = c.Loss + 1
	}
	if r.Switch == true && r.Win == false {
		c.Win = c.Win + 1
	}
}

func (c *Counter) Tally() {
	fmt.Printf("Number of SwitchWins = %d\n", c.SwitchWin)
	fmt.Printf("Number of SwitchLoss = %d\n", c.SwitchLoss)
	fmt.Printf("Number of Wins = %d\n", c.Win)
	fmt.Printf("Number of Loss = %d\n", c.Loss)
}

type Round struct {
	Doors [3]bool
	Pick int
	Reveal int
	Switch bool
	Win bool
}

func NewRound() *Round {
	r0 := new(Round)
	f0 := rand.Intn(3)
	r0.Doors[f0] = true
	return r0
}

func (r *Round) Picker() {
	r.Pick = rand.Intn(3)
	r.Reveal = r.Revealer()
	f0 := rand.Intn(2)
	if f0 == 1 {
		r.Switch = true
		r.Pick = r.SwitchPick()
	}
}

func (r *Round) Revealer() {
	p0 := []int
}

func (r *Round) SwitchPick() int {
/*
	i0 := 0
	for i := 0; i < 3; i++ {
		if i != r.Pick && r.Doors[i] == false {
			i0 = i
			break
		}
	}
	return i0
	*/
	if r.Doors[r.Pick] == true {
		
	}
}

func (r *Round) Test() {
	if r.Doors[r.Pick] == true {
		r.Win = true
	}
}

func main() {
	// set number of runs
	N := 10000
	fmt.Printf("starting monte hall problem simulator with %d runs\n", N)
	
	// init rand
	rand.Seed(0x1337)
	
	// start counter
	C := NewCounter()
	
	// run simulation
	for i := 0; i < N; i++ {
		C.Run = C.Run + 1
		R := NewRound()
		R.Picker()
		R.Test()
		C.Tick(R)
		if (i % 1000) == 0 {
			fmt.Printf("Percent complete: %f%%. Switches: %d. Wins: %d.\n", 100.0*float32(i)/float32(N), C.SwitchWin+C.SwitchLoss, C.SwitchWin)
		}
	}
	
	fmt.Println("ending runs and counting tallies")
	C.Tally()
}

// monte hall problem
package main

import (
	"fmt"
	"math/rand"
)

type Counter struct {
	Run int
	SwitchWin int
	SwitchLoss int
	Win int
	Loss int
}

func NewCounter() *Counter {
	c0 := new(Counter)
	c0.Run = 0
	c0.SwitchWin = 0
	c0.SwitchLoss = 0
	c0.Win = 0
	c0.Loss = 0
	return c0
}

func (c *Counter) Tick(r *Round) {
	if r.Switch == true && r.Win == true {
		c.SwitchWin = c.SwitchWin + 1
	}
	if r.Switch == true && r.Win == false {
		c.SwitchLoss = c.SwitchLoss + 1
	}
	if r.Switch == false && r.Win == false {
		c.Loss = c.Loss + 1
	}
	if r.Switch == false && r.Win == true {
		c.Win = c.Win + 1
	}
}

func (c *Counter) Tally() {
	fmt.Printf("Number of SwitchWins = %d\n", c.SwitchWin)
	fmt.Printf("Number of SwitchLoss = %d\n", c.SwitchLoss)
	fmt.Printf("Number of Wins = %d\n", c.Win)
	fmt.Printf("Number of Loss = %d\n", c.Loss)
	// switch win probabilities
	p0 := float(c.SwitchWin)/(float(c.SwitchWin)+float(c.SwitchLoss))
	p1 := float(c.Switch)/(float(c.Win)+float(c.Loss))
	fmt.Printf("Win probability with switching: %f%%\n", 100.0*p0)
	fmt.Printf("Win probability with remaining: %f%%\n", 100.0*p1)
}

type Round struct {
	Doors [3]bool
	Pick int
	Reveal int
	Choice int
	Switch bool
	Win bool
}

func NewRound() *Round {
	r0 := new(Round)
	f0 := rand.Intn(3)
	r0.Doors[f0] = true
	r0.Pick = 0
	r0.Reveal = 0
	r0.Choice = 0
	r0.Switch = false
	r0.Win = false
	return r0
}

func (r *Round) Picker() {
	r.Pick = rand.Intn(3)
	r.Reveal = r.Revealer()
	r.Choice = r.Chooser()
	f0 := rand.Intn(2)
	if f0 == 1 {
		r.Switch = true
		r.Pick = r.Choice
	}
}

func (r *Round) Revealer() int {
	i0 := 0
	for i := 0; i < 3; i++ {
		if i != r.Pick && r.Doors[i] == false {
			i0 = i
			break
		}
	}
	return i0
}

func (r *Round) Chooser() int {
	i0 := 0
	// index left to choose
	for i := 0; i < 3; i++ {
		if i != r.Reveal && i != r.Pick {
			i0 = i
			break
		}
	}
	return i0
}

func (r *Round) Test() {
	if r.Doors[r.Pick] == true {
		r.Win = true
	}
}

func main() {
	// set number of runs
	N := 100000
	fmt.Printf("starting monte hall problem simulator with %d runs\n", N)
	
	// init rand
	rand.Seed(0x1337)
	
	// start counter
	C := NewCounter()
	
	// run simulation
	for i := 0; i < N; i++ {
		C.Run = C.Run + 1
		R := NewRound()
		R.Picker()
		R.Test()
		C.Tick(R)
		/*
		if (i % 1000) == 0 {
			fmt.Printf("Percent complete: %f%%. Switches: %d. Wins: %d.\n", 100.0*float32(i)/float32(N), C.SwitchWin+C.SwitchLoss, C.SwitchWin)
		}
		*/
	}
	
	fmt.Println("ending runs and counting tallies")
	C.Tally()
}

// golang playground
// https://play.golang.org/p/RGkJ9dToda9


