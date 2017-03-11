package main
import (
	"fmt"
	"math/rand"
	"time"
)


type entity struct  {
	Name string
	Title string
	HP int
	MP int

	Strength int
	Speed int
	Intelligence int

	HPcounter int
	MPcounter int
	StrengthCounter int
	IntCounter int
	SpeedCounter int
	LazyCounter int
}


//So with the *entity in the func and &entity in the return this will be a pointer to the struct
//Without the star, this will be a copy. Potentially useful for creating heros. 
func createHero () entity  {
	return entity {
		Name : "Cheesecake",
		Title : "TheBrave",
		HP : 10,
		MP : 10,
		Strength : 3,
		Speed : 3,
		Intelligence : 2,
		HPcounter : 0,
		MPcounter : 0,
		StrengthCounter : 0,
		SpeedCounter : 0,
		LazyCounter : 0,
	}
}


//FOR EACH TYPE CREATE ENTITY. 
func populous() {

}

var hero = createHero()



func main () {
	// main game loop
		counter := 0

	for {
		action := rand.Intn(10)
		time.Sleep(100 * time.Millisecond)

		switch {
		case action == 0:
			hero.LazyCounter++
		case action == 1:
			hero.HPcounter++
		case action == 2:
			hero.MPcounter++
		case action == 3:
			hero.StrengthCounter++
		case action == 4:
			hero.SpeedCounter++
		case action == 5:
			hero.StrengthCounter++
		case action == 6:
			hero.MPcounter++
		case action == 7:
			hero.HPcounter++
		case action == 8:
			hero.StrengthCounter++
		case action == 9:
			hero.HPcounter++
		case action == 10:
			hero.LazyCounter++
		}
		switch {
		case hero.HPcounter == 10 + (hero.HP/10):
			hero.HP++
			hero.HPcounter = 0
		case hero.MPcounter == 10 + (hero.MP/10):
			hero.MP++
			hero.MPcounter = 0 
		case hero.StrengthCounter == 10 + (hero.Strength/10):
			hero.Strength++
			hero.StrengthCounter = 0 
		case hero.SpeedCounter == 10 - (hero.LazyCounter/10):
			hero.Speed++
			hero.SpeedCounter = 0
		case hero.LazyCounter == 20:
			hero.Speed--
			hero.LazyCounter = 0 
		}
		fmt.Println("HP: ", hero.HP)
		fmt.Println("MP: ", hero.MP)
		fmt.Println("Strength", hero.Strength)
		fmt.Println("Speed", hero.Speed)
		fmt.Println(hero.LazyCounter)
		fmt.Println(counter)
		fmt.Println("--------")
				counter++

		
	}
}
