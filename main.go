package main

func main() {
	sur := leerCartas("Sur")
	sur.print()

	contrato := getContrato()
	contrato.print()
	// (2)Info adicional

	// Salida Oponente (Quien y carta)

	norte := leerCartas("Norte")
	norte.print()

	competencia := otrasCartas(sur,norte)
	competencia.print()

	//Loop hasta ultima vuelta
	// carta propuesta (Next a quien y Carta)
	// carta oponente
	//Fin Loop
}