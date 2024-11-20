package main

import (
	"fmt"
	"github.com/balub/simple-http/api"
	"sync"
)

func main() {
	    var wg sync.WaitGroup


	pokemonList := [3]string{"charmander", "pikachu","ditto"}

	for i := 0; i < len(pokemonList); i++ {

wg.Add(1)

		go func(name string){
			
		pokemon, err := api.GetPokemonDetails(name)
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println(pokemon)
		 defer wg.Done()
		}(pokemonList[i])	
	}

	 wg.Wait()
}
