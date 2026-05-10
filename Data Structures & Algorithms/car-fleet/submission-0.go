
type Car struct {
	pos int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = Car {
			pos: position[i],
			speed: speed[i],
		}
	}

	sort.Slice(cars, func (i, j int) bool{
		return cars[i].pos > cars[j].pos
	})

	fleet := 0
	var maxTime float64
	for _, car := range cars {
		currTime := float64(target - car.pos) / float64(car.speed)
		if currTime > maxTime || fleet == 0 {
			maxTime = currTime
			fleet++
		}
	} 

	return fleet
}
