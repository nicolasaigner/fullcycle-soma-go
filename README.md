# Full Cycle - Desafio CI - Soma Go Lang

1) Você deverá criar uma simples aplicação que possua uma função chamada soma que receberá dois parâmetros e retornará a adição desses dois valores.

Essa função deverá ser chamada na função main do programa. Quando executada, deverá exibir da na tela o resultado da soma de 5 + 5.

2) Crie um teste unitário para essa função.

3) Ative um processo de CI que execute os seguintes passos:

Executar o teste unitário
Push da imagem gerada no processo de CI no Container Registry do GCP
Ative a App do Cloud Build no Github para que cada pull request execute automaticamente o processo de CI
