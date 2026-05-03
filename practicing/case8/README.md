Você está no time de backend.

Todo anúncio criado na plataforma passa por um sistema de moderação antes de ser publicado.
Os anúncios entram em uma fila e devem ser revisados na ordem em que chegaram — o mais antigo primeiro.
Seu tech lead pediu para você implementar esse sistema de fila de moderação com três operações básicas:

enfileirar(anuncio) — adiciona um anúncio no final da fila
desenfileirar() — remove e retorna o anúncio mais antigo da fila
esta_vazia() — retorna verdadeiro se não houver anúncios aguardando

Exemplo de uso esperado:
fila = FilaModeracao()

fila.enfileirar("iPhone 13")
fila.enfileirar("Notebook Dell")
fila.enfileirar("Galaxy S22")

fila.desenfileirar() → retorna "iPhone 13"
fila.desenfileirar() → retorna "Notebook Dell"

fila.esta_vazia() → falso (ainda tem "Galaxy S22")
