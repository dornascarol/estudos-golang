package main

import "fmt"

type FilaModeracao struct {
    anuncios []string
}

func (f *FilaModeracao) enfileirar(anuncio string) {
    f.anuncios = append(f.anuncios, anuncio)
}

func (f *FilaModeracao) desenfileirar() string {
    if f.estaVazia() {
        return "fila vazia"
    }

    primeiro := f.anuncios[0]
    f.anuncios = f.anuncios[1:]  
    return primeiro
}

func (f *FilaModeracao) estaVazia() bool {
    return len(f.anuncios) == 0
}

func main() {
    fila := FilaModeracao{}

    fila.enfileirar("iPhone 13")
    fila.enfileirar("Notebook Dell")
    fila.enfileirar("Galaxy S22")

    fmt.Println(fila.desenfileirar()) 
    fmt.Println(fila.desenfileirar()) 
    fmt.Println(fila.estaVazia())     
}
