Você está no time de backend.

O time de produto quer lançar uma funcionalidade de buscas recentes.
Igual ao que você vê em apps de marketplace quando clica na barra de busca e aparecem seus últimos termos pesquisados.

As regras são:

- Guardar no máximo 5 buscas recentes por usuário
- Se o usuário buscar um termo que já buscou antes, não duplicar — só mover para o topo
- A busca mais recente aparece primeiro
- Se já tiver 5 buscas e chegar uma nova, a mais antiga é removida

Seu tech lead pediu duas operações:

adicionar_busca(termo) — registra uma nova busca
buscar_recentes() — retorna a lista de buscas na ordem mais recente primeiro

Exemplo de uso esperado:
historico = HistoricoBuscas()

historico.adicionar_busca("iPhone 13")
historico.adicionar_busca("Notebook Dell")
historico.adicionar_busca("iPhone 13") # já existe — move para o topo

historico.buscar_recentes()
→ ["iPhone 13", "Notebook Dell"] # iPhone 13 aparece primeiro

historico.adicionar_busca("Galaxy S22")
historico.adicionar_busca("Bike Trek")
historico.adicionar_busca("Van Fiat") # agora tem 5
historico.adicionar_busca("PS5") # sexta busca — remove a mais antiga

historico.buscar_recentes()
→ ["PS5", "Van Fiat", "Bike Trek", "Galaxy S22", "iPhone 13"]
