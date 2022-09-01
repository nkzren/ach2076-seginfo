# Criptoanálise

## Modelo da Criptografia Simétrica


## Ataque Força-Bruta

Um ataque força-bruta consiste em testar todas as chaves possíveis. Em um
universo de `K` chaves, no pior caso temos que testar `|K|` e o valor esperado de
tentativas é `|K| / 2`

### Exemplo

* Cifra de deslocamento
```
K = Z(26)

|K| = 26
```

* Cifra de Substituição
```
K = Perm(Z(26))

|K| = |Z(26)| = 26! ~ 4 * 10^26 ~ 2 ^ 88
```

## Ataques de Frequência

Em uma cifra de substituição, podemos decifrar o texto analisando a frequencia
de ocorrência dos símbolos sem precisarfazer um ataque força-bruta.

## Teste de Kasiski

```
m = C R Y P T O I S S H O R T F O R C R Y P T O G R A P H Y

k = A B C D A B C D A B C D A B C D A B C D A B C D A B C D 

c = C S A S T P K V S I Q U T G Q U C S A S T P I U A Q J B

    C       T       S       T       C       T       A
      S       P       I       G       S       P       Q
        A       K       Q       Q       A       I       J
          S       V       U       U       S       U       B
```

## Sigilo Perfeito

* (E) = pertence a

* (!E) = não pertence a

**Def.:** Um sistema de criptografia `pi = < Gen, E, D >` garante o *sigilo
perfeito* se para toda distribuição de probabilidades sobre M temos que:
`Pr[M=m|C=c] = Pr[M=m]` supondo que `Pr[C=c]=0` (revisar essa última fórmula)

De maneira equivalente, para todo `m0, m1 (E) M` e  c (E) C temos que `Pr[C=c |
M=m0] = Pr[C=c|M=m1]`

Um sistema `pi` garante sigilo perfeito se para qualquer adversário A:

```
Pr[Priv(K) = 1] = 1 / 2
```

*tldr: a probabilidade de uma mensagem e sua cifra ocorrerem são independentes*
*- a cifra não dá informações sobre a mensagem*

### Exemplo

Seja `pi = < Gen, E, D >` o sistema de criptografia de substituição. Sejam c =
ANA, m0 = OVO e m1 = EVA

```
Pr[C=c|M=m0] = 1 / (26 * 25)

Pr[C=c|M=m1] = 0
```

## *One Time Pad (OTP)*

* Gen = algoritmo que gera uma chave
* E = algoritmo que criptografa a chave
* D = algoritmo que descriptografa a chave
* (+) = XOR

```
pi = < Gen, E, D >

M = C = K = {0,1} ^ n // mensagem, cifra, chave são sequencias de n bits

Gen := k <- {0,1} ^ n

E(k,m) = [m0 + k0 mod2] ... [mN + kN mod2] = m (+) k

D(k,c) = c (+) k
```

```
           (c)
            |
E(k,m) = m (+) k

D(k,m (+) k) = (m (+) k) (+) k = m (+) (k (+) k) = m
                                           |
                                           0
```

### Exemplo

```
m =      1 0 1 0 1 0
                      (+)
k =      0 1 0 0 0 1

E(k,m) = 1 1 1 0 1 1 = c

D(k,m) = 1 0 1 0 1 0
```

-----------

**Teorema:** O *One Time Pad* garante sigilo perfeito. 

P.: Dada uma cifra `c` e uma mensagem qualquer `m (E) M` existe uma única chave
`k` tal que `E(k,m) = c`. A saber, `k = m (+) c`

Portanto, a probabilidade `Pr[C=c|M=m] = 1 / (2 ^ n)` para qualquer `m (E) M`.
Ou seja, para quaisquer `m0, m1 (E) M`, `Pr[C=c|M=m0] = Pr[C=c|M=m1] = 1 / (2 ^ n)` 

**Teorema (Shannon):** Seja `pi` um sistema que garante o sigilo perfeito então
`|K| >= |M|`

P.: Seja `M(c)` o conjunto de todas as mensagens que podem produzir `c`. Se
existissem `m0 =/= m1` tais que `E(k, m0) = E(k, m1) = c`, `pi` não teria como
ser correto.

Portanto `|M(c)| <= |K|`. Agora suponha que `|K| < |M|`. Neste caso existiria `m
(E) M` tal que `m (!E) M(c)`. Teríamos então `Pr[C=c|M=m] = 0`. Mas `Pr[M=m] =/=
0`, contradizendo a definição de sigilo perfeito.

Concluímos que `|M| <= |K|`.

