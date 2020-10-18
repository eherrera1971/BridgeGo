func leerManoPropia(filename string) cartas {
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error:", err, "Hay un error en la lectura del archivo")
		os.Exit(1)
	}
	a := carta{"S",14,"B"}
	c := cartas{}
	c = append(c,a)
	return strings.Split(string(bs),",")
}
func deal(d cartas, size int) (cartas, cartas){
	return d[:size], d[size:]
}

func newDeck() cartas {
	cards := cartas{}
	suits := []string{"♠️","♥️","♦️","♣️"}
	values := []string{"A","K","Q","J","T","9","8","7","6","5","4","3","2"}
	for _,suit := range suits {
		for i,letra := range values {
			a := carta{suit,14-i,letra}
			cards = append(cards,a)
		}
	}
	return cards
}

func (d cartas) saveMano(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}