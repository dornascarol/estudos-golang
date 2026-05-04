Você está no time de backend.

O time de produto identificou que muitos usuários estão vendo os mesmos anúncios repetidamente no feed, o que piora a experiência.

A solução é simples: guardar quais anúncios cada usuário já visualizou e nunca exibir o mesmo anúncio duas vezes para a mesma pessoa.

Seu tech lead pediu para você implementar esse sistema com três operações:

registrar_visualizacao(id_anuncio) — marca que o usuário viu esse anúncio
ja_visualizou(id_anuncio) — retorna verdadeiro se o usuário já viu esse anúncio
total_visualizados() — retorna quantos anúncios distintos o usuário já viu

Exemplo de uso esperado:
historico = HistoricoVisualizacao()

historico.registrar_visualizacao("anuncio-001")
historico.registrar_visualizacao("anuncio-002")
historico.registrar_visualizacao("anuncio-001") # viu de novo — não conta duas vezes

historico.ja_visualizou("anuncio-001") → verdadeiro
historico.ja_visualizou("anuncio-999") → falso

historico.total_visualizados() → 2 # não 3, porque 001 não conta duas vezes
