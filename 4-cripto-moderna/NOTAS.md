# Criptografia moderna

Tópicos principais dessa aula:

* Definições formais de segurança
* Explicitar as suposições
* Demonstrações matemáticas de que dadas certas suposições, o sistema garante
  certa noção de segurança

## Definição formal de segurança

Toda definição de segurança possui dois componentes:

1. *Modelo de ameaça (threat model):* quais informações supomos que nosso
   adversário possui
2. *Garantias de segurança:* O que pode ser considerado um ataque bem sucedido

No caso do sigilo perfeito: 
* Modelo de ameaça: Um atacante *A* só tem acesso ao texto cifrado *c*
* Garantia de segurança: a probabilidade de qualquer adversário *A* acertar qual
  das mensagens foi criptografada é 1/2

---

Durante o curso veremos noções de segurança com diferentes modelos de ameaça:

* *Cyphertext only*
* *Chosen plaintext*: além de conhecer a cifra, supomos que *A* sabe como certas
  mensagens *m* (escolhidas por ele) são criptografadas
* *Chosen cyphertext*: além de conhecer a cifra e saber como certas mensagens
  escolhidas são cifradas, *A* sabe como certas cifras (escolhidas por ele) são
  cifradas

## Abordagem assintótica

Digamos que uma função é desprezível se ela cresce mais devagar do que `1 /
P(n)` para qualquer polinômio *P*

1. O sistema deve ser seguro contra qualquer *A* polinomial
2. *A* pode vencer o jogo com probabilidade `1/2 + E(n)` em que `E` é
   desprezível

**Definição:** Seja `l` um polinômio e `G` um algoritmo determinístico que
recebe `s (E) {0,1} ^ n` e devolve `G(s) (E) {0,1} ^ l(n)` e seja `r <- {0,1} ^
l(n)`, então `G` é um *gerador de números pseudo-aleatórios* se:

1. `l(n) > n` para todo n
2. Para todo algoritmo polinomial D (distinguidor): ` | Pr[D(r) = 1] -
   Pr[D(G(s)) = 1] | <= E(n) ` para algum `E` desprezível

## Cifra de Fluxo

* `(+)` -- XOR

Seja `G` um gerador de números pseudo-aleatórios. Uma cifra de fluxo é o seguinte
sistema:

```
pi = <Gen, E, D>

Gen(n) = k <- {0,1} ^ n 

E(k, m) = G(k) (+) m
D(k, c) = G(k) (+) c
```

**Teorema:** Seja `G` um gerador de números pseudo-aleatórios com fator de
expansão `l`, o sistema `pi` definido anteriormente é seguro contra ataques
*cyphertext-only* para mensagens `m` de tamanho fixo `l(n)`
