Você está no time de backend.

O time quer criar uma funcionalidade de busca inteligente que ignora acentos e maiúsculas/minúsculas. 

Para isso, antes de salvar um anúncio, o sistema precisa verificar se o título já existe de forma equivalente — mesmo que escrito de forma diferente.

Seu tech lead pediu uma função menor e mais específica primeiro, para validar o conceito: dado um título de anúncio, contar quantas vezes cada caractere aparece nesse título (ignorando espaços).

Entrada:
titulo = "iPhone Pro"

Saída esperada:
{ "i": 1, "P": 1, "h": 1, "o": 2, "n": 1, "e": 1, "r": 1 }
