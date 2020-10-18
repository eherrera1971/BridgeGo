package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)
type carta struct {
	suit string
	valor int
	letra string
}

type cartas []carta

func (m cartas) print(){
	for _, card:= range m {
		fmt.Print(card.suit+" "+card.letra+" ")}
	fmt.Println("")
}

func otrasCartas(s cartas,n cartas) cartas {
	o := cartas{}
	letras := map[int]string {
		14:"A", 13:"K",12:"Q",11:"J", 10:"10",
		9:"9",8:"8",7:"7",6:"6",5:"5",4:"4",3:"3",2:"2",}
	suits := map[int]string {
		0: "♠️", 1:"♥️",2:"♦️",3:"♣️" }
	for i := 14; i > 1; i-- {
		for j:=0; j < 4; j++ {
			c:= carta{suits[j],i,letras[i]}
			if (s.noContiene(c) && n.noContiene(c)) {
				o = append(o,c)
			}
		}
	}
	return o
}

func (d cartas) noContiene(c carta) bool {
    for _, n := range d {
        if c == n {
            return false
        }
    }
    return true
}

func (d cartas) toString() string {
	return d[0].letra+d[0].suit+","+d[1].letra
}

func leerCartas(jugador string) cartas {
	mano := cartas{}
	for len(mano)!=13 {
	mano = cartas{}
	mano.leerSuit(jugador,"♠️")
	mano.leerSuit(jugador,"♥️")
	mano.leerSuit(jugador,"♦️")
	mano.leerSuit(jugador,"♣️")
	fmt.Println("Cantidad de cartas: ",len(mano))
	}
	return mano
}

func (p *cartas) leerSuit(jugador string, suit string) {
	valores := map[string]int {
		"A":14, "K":13,"Q":12,"J":11,
		"T":10, "t":10, "10":10,
		"9":9,"8":8,"7":7,"6":6,"5":5,"4":4,"3":3,"2":2,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("("+jugador+") "+suit+">")
	text,_ := reader.ReadString('\n')
	if (text=="\n") {
		fmt.Println(suit + " Vacio")
	} else {
	i := 0
	for i < len(text)-1 {
		letra := strings.ToUpper(string([]rune(text)[i]))
		if (letra=="1") {
			letra = string([]rune(text)[i:i+1])
			i++
		}
		c := carta{suit,valores[letra],letra}
		*p = append((*p),c)
		i++
	}
}
}
