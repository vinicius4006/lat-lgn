Claro! Aqui está um exemplo de `README.md` explicando como usar o package `latlgn` com base no trecho de código que você enviou:

---

# 📍 latlgn - Obtenha Latitude e Longitude a partir de um CEP

Este projeto utiliza o package `latlgn` para converter informações de endereço em coordenadas geográficas (latitude e longitude).

## ✅ Requisitos

- Go 1.18 ou superior
- Internet (para consumir o serviço de geolocalização)

## 📦 Instalação

Adicione o package ao seu projeto:

```bash
go get github.com/seu-usuario/latlgn
```

> Substitua pelo caminho real do módulo se for diferente.

## 🧩 Como usar

```go
package main

import (
	"fmt"
	"log"

	"github.com/seu-usuario/latlgn"
)

func main() {
	cep := latlgn.CEP{
		Logradouro: "Rua das Flores, 123",
		Localidade: "Curitiba",
		UF:         "Paraná",
	}

	lat, lng, err := latlgn.LatitudeLongitudeByCEP(cep)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lng)
}
```

### 📌 Struct CEP

A struct `CEP` é composta por:

- `Logradouro`: nome da rua e número
- `Localidade`: cidade
- `UF`: estado (por extenso ou sigla)

## 📤 Resultado esperado

```text
Latitude: -25.4284, Longitude: -49.2733
```