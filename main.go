package main

import (
	"fmt"
	"os"
)

func main() {

	favorietekleuren := map[string]string{
		"Blauw":  "Blauw zoals de lucht.",
		"Rood":   "Rood met passie.",
		"Geel":   "Geel als de stralen van de zon.",
		"Groen":  "Groen van de natuur.",
		"Oranje": "Oranje zoals een sinaasappel.",
		"Bruin":  "Bruin zoals het ongelukje in mijn broek.",
		"Paars":  "Paars als een zwerende vinger.",
	}

	var kleuren []string

	for {
		var kleur string
		fmt.Print("Wat is je favoriete kleur? ")
		fmt.Scanln(&kleur)

		if kleur == "stop" {
			break
		}

		text, ok := favorietekleuren[kleur]
		if !ok {
			fmt.Println("Fout: deze kleur wordt niet ondersteund")
			continue
		}

		kleuren = append(kleuren, text)
	}

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Fout bij het maken van output bestand")
		os.Exit(1)
	}

	for _, text := range kleuren {
		_, err = file.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Fout bij het schrijven naar output bestand")
			os.Exit(1)
		}
	}

	file.Close()
	fmt.Println("Tekst succesvol geschreven naar output bestand")
}
