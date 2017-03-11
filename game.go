package main
import (
	"fmt"
	"math/rand"
	"time"
	"github.com/wawandco/fako"
)


type entity struct  {
	Name string
	Title string
	HP int
	MP int

	Strength int
	Speed int64
	Intelligence int

	Level int
	CurrentXP int
	DropXP int
}





func randomize(min, max int, seed int64) int {
	rand.Seed(seed + seed ^ time.Now().Unix())
	inbetween := max - min
	if inbetween <= 0 {
		inbetween = 1337
	}
	return rand.Intn(inbetween) + min
}


func main () {
	var hero entity
	var monster entity
	fako.Fuzz(&hero)
	fako.Fuzz(&monster)
	
	for {
		var input string
		fmt.Println("Attack monster?")
		fmt.Println("Your currnet health: ", hero.HP)
		fmt.Println("Monster health: ", monster.HP)
		fmt.Scanln(&input)
		switch {
		case input == "y":
			hit := randomize(hero.Strength, hero.Intelligence, hero.Speed)
			mhit := randomize(monster.Strength, monster.Intelligence, monster.Speed)
			fmt.Println("You have hit the monster for ", hit ,"damage")
			monster.HP = monster.HP - hit

			if monster.HP <= 0 {
				fmt.Println("You have killed the monster, yay, new monster!")
				fako.Fuzz(&monster)
			} else { 

			fmt.Println("The monster hits you for", mhit, "damage")
			hero.HP = hero.HP - mhit
			if hero.HP <= 0 {
				fmt.Println("You have died, blarg")
				fmt.Println("New hero time")
				fako.Fuzz(&hero)
				
			}
			}
		
	}
	
}
}
