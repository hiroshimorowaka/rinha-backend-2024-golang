# API - Rinha de backend 2024

Eu sou um programador considerado iniciante, então foi um grande exercicio pra mim participar dessa rinha, além de muito divertido.  

Eu nunca tinha usado GO, foi a primeira vez e meu primeiro projeto, então com toda certeza não está otimizado ao máximo possivel e poderia ter sido bem melhor os resultados, mas gostei e serviu de muito aprendizado pra mim.   
Qualquer feedback será muito bem vindo

## Stack
  - Golang
  - Gin-Gonic (Web Server)
  - PostgreSQL (Banco de dados relacional)
  - pgxpool (Pool de conexões)

## Como rodar o projeto

  Esse projeto roda com network_mode HOST do docker, que só é suportado em sistemas operacionais Linux, ou seja, se você estiver usando Windows, você tem duas opções:
  - Usar o docker dentro do WSL2 (foi oq eu fiz)
  - Usar o Rancher Desktop, uma alternativa ao Docker Desktop 

O docker compose está configurado para BUILDAR a imagem do docker file na raiz do projeto, mas caso queira usar a imagem que eu criei, é só comentar a linha "build" e descomentar a linha "image" do arquivo  `compose.yaml` 

Caso queira rodar o projeto somente pela imagem, os arquivos necessário são:
  - compose.yaml
  - nginx.conf
  - init.sql

Agora é só clonar o projeto e executar:
  ```bash
  docker compose up [-d]
  ```
 Depois de subir os containers, a API já vai estar escutando na porta 9999  


  
