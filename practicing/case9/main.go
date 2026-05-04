package main

import "fmt"

type HistoricoVisualizacao struct {
	vistos map[string]bool
}

func novoHistorico() HistoricoVisualizacao {
	return HistoricoVisualizacao{
		vistos: make(map[string]bool),
	}
}

func (h *HistoricoVisualizacao) registrarVisualizacao(idAnuncio string) {
	h.vistos[idAnuncio] = true
}


func (h *HistoricoVisualizacao) jaVisualizou(idAnuncio string) bool {
	return h.vistos[idAnuncio]
}


func (h *HistoricoVisualizacao) totalVisualizados() int {
	return len(h.vistos)
}

func main() {
	historico := novoHistorico()

	historico.registrarVisualizacao("anuncio-001") 
	historico.registrarVisualizacao("anuncio-002") 
	historico.registrarVisualizacao("anuncio-001") 

	fmt.Println(historico.jaVisualizou("anuncio-001")) 
	fmt.Println(historico.jaVisualizou("anuncio-999")) 
	fmt.Println(historico.totalVisualizados())         
}
