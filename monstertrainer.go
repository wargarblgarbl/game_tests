

package main
import (
	"fmt"
//	"math/rand"
//	"time"
)

var entities []*entity

type entity struct  {
	Name string
	Title string

	HP int
	MP int
	STR int
	INT int

	HPcounter int
	MPcounter int
	STRCount int
	IntCount int
	SPDCount int
	LZYCount int

	Speed int
	Gender string
	Type string

	PosX int
	PosY int
}

//So with the *entity in the func and &entity in the return this will be a pointer to the struct
//Without the star, this will be a copy. Potentially useful for creating heros.
func createHero (name string) *entity  {
	return &entity {
		Name : name,
		HP : 10,
		MP : 10,
		STR : 3,
		Speed : 3,
		INT : 2,
	}
}

//FOR EACH TYPE CREATE ENTITY.
func populous() {
	critters := []string{"human", "goblin", "cobold" }
	for _, i := range critters {
		hero := createHero(i)
		fmt.Println(hero)
		entities = append(entities, hero)
	}
}

//var hero = createHero()



func main () {
	//init population
	fmt.Println("initializing population"	)
	populous()
	for _, i := range entities {
		fmt.Println(i.Name)
	}
	
}
/*
	// main game loop
		counter := 0

	for {
		action := rand.Intn(10)
		time.Sleep(100 * time.Millisecond)

		switch {
		case action == 0:
			hero.LZYCount++
		case action == 1:
			hero.HPcounter++
		case action == 2:
			hero.MPcounter++
		case action == 3:
			hero.STRCount++
		case action == 4:
			hero.SPDCount++
		case action == 5:
			hero.STRCount++
		case action == 6:
			hero.MPcounter++
		case action == 7:
			hero.HPcounter++
		case action == 8:
			hero.STRCount++
		case action == 9:
			hero.HPcounter++
		case action == 10:
			hero.LZYCount++
		}
		switch {
		case hero.HPcounter == 10 + (hero.HP/10):
			hero.HP++
			hero.HPcounter = 0
		case hero.MPcounter == 10 + (hero.MP/10):
			hero.MP++
			hero.MPcounter = 0
		case hero.STRCount == 10 + (hero.STR/10):
			hero.STR++
			hero.STRCount = 0
		case hero.SPDCount == 10 - (hero.LZYCount/10):
			hero.Speed++
			hero.SPDCount = 0
		case hero.LZYCount == 20:
			hero.Speed--
			hero.LZYCount = 0
		}
		fmt.Println("HP: ", hero.HP)
		fmt.Println("MP: ", hero.MP)
		fmt.Println("STR", hero.STR)
		fmt.Println("Speed", hero.Speed)
		fmt.Println(hero.LZYCount)
		fmt.Println(counter)
		fmt.Println("--------")
				counter++


	}
}
*/
