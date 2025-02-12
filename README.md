## Desafio #08 - Stress Test

O objetivo deste desafio é criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

### Requisitos:

A aplicação utilizará os seguintes parâmetros de entrada:

**--url:** URL do serviço a ser testado.
**--requests:** Número total de requests.
**--concurrency:** Número de chamadas simultâneas.

> Após cada execução, o sistema deverá gerar um relatório com informações específicas após a execução dos testes.

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                  
    │    └── root.go         # Configuração do Cobra CLI
    ├── pkg                  # Pacotes reutilizáveis utilizados na aplicação
    │    └── fcst            # Stress Test
    │           ├── http          ### Implementações de requisições HTTP
    │           └── fc_stress.go  ### Implementação do Stress Test
    ├── Dockerfile           # Arquivo de configuração do Docker da aplicação
    └── README.md

#### 🚀 Execução:
Para executar a aplicação em ambiente local, basta utilizar o docker-compose disponível na raiz do projeto. Para isso, execute o comando abaixo:
```bash
docker run tiagoncardoso/fc-stress-test --url=<URL_A_SER_TESTADA> --requests=<NUM_REQUISICOES> --concurrency=<NUM_DE_AGENTES>

# ou
docker run tiagoncardoso/fc-stress-test -u=<URL_A_SER_TESTADA> -r=<NUM_REQUISICOES> -c=<NUM_DE_AGENTES>
```

#### 📝 Relatório:

O relatório final contém as seguintes informações:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

#### ✍️ Exemplo de utilização do Rate Limiter

```shell
docker run tiagoncardoso/fc-stress-test --url=https://google.com --requests=100 --concurrency=10
```