package main

type Order struct {
	🍔 int
	🍟 int
	🍦 int
	🍺 int
}

func 🍺() string {
	return "beer"
}

func main() {
	o := Order{
		"🍔": 1,
		"🍟": 2,
		"🍦": 0,
		"🍺": 2,
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s\n", b)
}
