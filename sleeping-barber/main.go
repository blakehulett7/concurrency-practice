package main

func main() {
    ColorPrint(Cyan, "Dominus Iesus Christus")
    ColorPrint(Cyan, "----------------------")

    

	const waiting_room_size = 10
	waiting_customers := []Customer{}
	barber := Barber{IsAsleep: false}

	if len(waiting_customers) == 0 {
		barber.Sleep()
	}
}

type BarberShop struct {

}

type Customer struct {
}

type Barber struct {
	IsAsleep bool
}

func (b *Barber) Sleep() {
	b.IsAsleep = true
}

func CustomerArrivesSystem(barber_is_asleep chan bool) {
	for {
        is_asleep := <- barber_is_asleep
        if is_asleep {

        }

        if len()
	}
}
