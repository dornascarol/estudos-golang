package main

import "fmt"

type HistoricoBuscas struct {
	historico []string
}

func novoHistorico() HistoricoBuscas {
	return HistoricoBuscas{
		historico: []string{},
	}
}

func (h *HistoricoBuscas) adicionarBusca(termo string) {

	
	for i, t := range h.historico {
		if t == termo {
			h.historico = append(h.historico[:i], h.historico[i+1:]...)
			break
		}
	}


	if len(h.historico) == 5 {
		h.historico = h.historico[:len(h.historico)-1]
	}

	
	h.historico = append([]string{termo}, h.historico...)
}

func (h *HistoricoBuscas) buscarRecentes() []string {
	return h.historico
}

func main() {
	historico := novoHistorico()

	historico.adicionarBusca("iPhone 13")
	historico.adicionarBusca("Notebook Dell")
	historico.adicionarBusca("iPhone 13") 

	fmt.Println(historico.buscarRecentes())

	historico.adicionarBusca("Galaxy S22")
	historico.adicionarBusca("Bike Trek")
	historico.adicionarBusca("Van Fiat")
	historico.adicionarBusca("PS5") 

	fmt.Println(historico.buscarRecentes())
	
}
