package ejercicios

import "fmt"

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	actividadesElegidas := make([]Actividad, 0)
	actividadesElegidas = append(actividadesElegidas, actividades[0])
	aux(actividades[1:], &actividadesElegidas)
	return actividadesElegidas
}

func aux(actividad []Actividad, actSeleccionada *[]Actividad) {

	if len(actividad) == 0 {
		return
	}
	ultima := (*actSeleccionada)[len(*actSeleccionada)-1]
	if actividad[0].Inicio >= ultima.Fin {
		fmt.Println("entro")
		*actSeleccionada = append(*actSeleccionada, actividad[0])
	}
	aux(actividad[1:], actSeleccionada)
}
