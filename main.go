package main

import (
	"fmt"
)

func main() {
	// 	Créez et initialisez une map m avec en clée des strings et en valeur des ints
	m := make(map[string]int)
	// Créez une variable de type rune nommée letter et initialiser sa valeur à 'a'
	var letter rune = 'a'
	// Créez une boucle for avec i comme itérateur sur 26 éléments
	for i := 0; i < 26; i++ {
		// Insérer dans m la valeur de letter après l’avoir casté en string comme clée
		// et comme valeur insérer l’incrément i de la boucle for
		m[string(letter)] = i
		// Auto-incrémentez la valeur de letter
		letter++
	}
	// Après la boucle for afficher la valeur de la clé "w" dans la map m
	// via la fonction fmt.Println( ).
	fmt.Println(m["w"])
}
