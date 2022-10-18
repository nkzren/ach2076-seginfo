# Lista 1
Os códigos para o exercício, quando houver, estão em `src/exN`, onde N é o
número do exercício

## Exercício 1
Deslocamento k = 3
R: sulydflgdghsxeolfdwudqsduhqfldsulydgd

Substituição com ZEBRASCDFGHIJKLMNOPQTUVWXY
R: MOFUZBFRZRAMTEIFBZQOZKMZOAKBFZMOFUZRZ

Vigenere com chave == 'senha'
R: HVVCAUMQHDWTHILAGNARSRCHRWRPPAHVVCAVE

## Exercício 2

Ao utilizar uma palavra, o universo das chaves numa cifra de Vígenere está
limitada à quantidade de palavras do dicionário.

Utilizando uma sequência de tamanho `l`, esse universo passa a ser `26 ^ l`,
ou `~ 2 ^ (4.7 * l)`

## Exercício 3

Dadas duas mensagens, `m1` e `m2`, pela definição de Sigilo Perfeito, temos:
```
Pr[C = c | M = m1] = Pr[C = c | M = m2]
```
Ou seja, a probabilidade de m1 produzir a cifra `c`, deve ser idêntica à
probabilidade de qualquer outra mensagem m2 produzir a mesma cifra.

No entanto, supondo uma cifra de deslocamento com chave `k = 3` e uma cifra `c =
aa` temos:
```
D(3, 1 1) = [-2 mod 26][-2 mod 26] = 24 24
```
Logo, existe uma única mensagem `m = yy` Onde `Pr[C = c | M = m] = 1`. Para
todas as outras mensagens `mN`, `Pr[C = c | M = mN] = 0`

## Exercício 4

Um sistema de criptografia utilizando cifras de fluxo parte de uma sequência
aleatória de bits, chamada *seed* para gerar uma sequência de bits aleatória
maior do que a anterior. Essa nova sequência é utilizada para encriptar
mensagens utilizando o operador XOR (da mesma forma que o OTP)

Para assumirmos que esse sistema seja seguro, devemos assumir que a função
pseudo aleatória utilizada gera sequências de bits aleatórias o suficiente para
que não seja capaz de distiguir uma cifra gerada a partir dessa função de uma
sequência de bits realmente aleatória.

Tendo isso como verdade, podemos assumir que um sistema de cifra de fluxo é
seguro contra ataques *ciphertext only*, onde o adversário tem acesso a uma
cifra de tamanho arbitrário

## Exercício 5

`1/2 + 1 / (2 ^ 64)`

## Exercício 7

No modo Ctr, apenas 1 bit da mensagem seria alterado. No modo CBC, todos os bits
tem uma probabilidade de 50% de serem alterados
