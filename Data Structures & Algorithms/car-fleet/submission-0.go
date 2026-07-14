func carFleet(target int, position []int, speed []int) int {
    cars := make([]car, 0)
    for i := 0 ; i < len(position); i++ {
        cars = append(cars, car{
            pos: position[i],
            spd: speed[i],
        })
    }

    sort.Slice(cars, func (i, j int) bool {
        return cars[i].pos > cars[j].pos
    })

    mp := make(map[int]bool)
    for i := 0 ; i < len(cars); i++ {
        val := (target - cars[i].pos) % cars[i].spd
        steps := 0
        if val == 0 {
            steps = (target - cars[i].pos) / cars[i].spd
        } else {
            steps =  ((target - cars[i].pos) / cars[i].spd) + 1
        }
        _, ok := mp[steps]
        if !ok {
            mp[steps] = true
        }
    }

    return len(mp)
}

type car struct {
	pos int
    spd int
}
