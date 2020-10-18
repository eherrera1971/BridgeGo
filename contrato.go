package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type contrato struct {
	suit string
	valor string
}

func getContrato() contrato {
	c := contrato{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Contrato>")
	text,_ := reader.ReadString('\n')
	c.suit = strings.ToUpper(string([]rune(text)[1]))
	c.valor = strings.ToUpper(string([]rune(text)[0]))
	return c
}

func (c contrato) print(){
	fmt.Println(c)
}


