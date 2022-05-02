package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	Name  string
	Model string
}

func (r *Rocket) Launch() {

	fmt.Println(r.Name + " " + r.Model + " launched successfully")
}

func main() {

	rocket := &Rocket{"Qiruvchi", "RTX432"}

	time.AfterFunc(10*time.Second, rocket.Launch)

}
