package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//chaves devem ser do mesmo tipo [dentro dos colchetes]
	//valores devem ser do mesmo tipo []fora dos colchetes
	usuario := map[string]string{

		"nome":      "Lestat",
		"sobrenome": "De Lioncount",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Andrew Hozier",
			"ultimo":   "Hozier",
		},
		"curso": {
			"nome":   "antropologia",
			"campus": "USP",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"signo": "escorpião",
	}
	fmt.Println(usuario2)

}
