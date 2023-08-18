# rand-exp

This is a simple pet project that provides functions to generate random values for various purposes. 
It includes functions to generate random coordinates, text, sentences, and date/time within specified ranges.

## Getting Started

To use the random value generator functions in your Go project, follow these steps:

1. Import the required packages in your Go code:
   ```go
   import (

       "github.com/IraIvanishak/rand-exp/geo" // Import the geo package if needed
   )
2. Call the appropriate random value generator functions from your code. For example:
   ```go
   latitude, longitude := geo.RandomCoordinates()
   fmt.Printf("Generated Coordinates: Latitude: %f, Longitude: %f\n", latitude, longitude)
