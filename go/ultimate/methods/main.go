// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.

type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p player) average() float64 {
	return float64(p.hits) / float64(p.atBats)
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	ps := []player{
		{name: "Jacob", atBats: 20, hits: 16},
		{name: "Jenn", atBats: 24, hits: 18},
		{name: "Kell", atBats: 14, hits: 6},
		{name: "Carter", atBats: 12, hits: 6},
	}

	// Display the batting average for each player in the slice.
	for _, p := range ps {
		fmt.Printf("%s \thas an average of %f\n", p.name, p.average())
	}
}
