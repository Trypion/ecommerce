# Distributed E-commerce System

Este projeto é um sistema distribuído para portfólio, desenvolvido em Go (Golang), utilizando API REST, gRPC e RabbitMQ. O objetivo é simular um ambiente profissional de mercado, com arquitetura moderna, escalável e desacoplada, composta por múltiplos serviços.

## Arquitetura

A arquitetura do sistema é baseada em microsserviços, com os seguintes componentes principais:

- **API Gateway**: Ponto de entrada para clientes externos. Expõe endpoints REST, faz roteamento de requisições para os serviços internos via gRPC e gerencia autenticação/autorização.
- **Order Service**: Responsável pelo gerenciamento de pedidos. Expõe APIs gRPC para o gateway e publica eventos no RabbitMQ para outros serviços.
- **Payment Service**: Processa pagamentos de pedidos. Consome eventos de pedidos do RabbitMQ, processa o pagamento e publica eventos de confirmação.
- **Notification Service**: Responsável por enviar notificações (e-mail, SMS, etc). Consome eventos de pagamento e pedido do RabbitMQ.
- **RabbitMQ**: Broker de mensagens para comunicação assíncrona entre os serviços.

### Fluxo de Comunicação

1. O cliente faz uma requisição REST para o API Gateway.
2. O gateway converte a requisição em uma chamada gRPC para o serviço apropriado (ex: Order Service).
3. O Order Service processa o pedido e publica um evento no RabbitMQ.
4. O Payment Service consome o evento, processa o pagamento e publica um novo evento.
5. O Notification Service consome eventos e envia notificações ao usuário.

### Tecnologias Utilizadas
- **Go (Golang)**: Linguagem principal dos serviços.
- **gRPC**: Comunicação eficiente entre microsserviços.
- **REST**: Interface externa para clientes.
- **RabbitMQ**: Mensageria para eventos assíncronos.
- **Docker**: Containerização dos serviços.
- **Makefile**: Automatização de tarefas.

## Estrutura de Pastas Inicial

```
/
├── api-gateway/
│   ├── cmd/
│   ├── internal/
│   └── Dockerfile
├── order-service/
│   ├── cmd/
│   ├── internal/
│   └── Dockerfile
├── payment-service/
│   ├── cmd/
│   ├── internal/
│   └── Dockerfile
├── notification-service/
│   ├── cmd/
│   ├── internal/
│   └── Dockerfile
├── proto/
│   └── (arquivos .proto para gRPC)
├── deployments/
│   └── (YAMLs do Docker Compose ou Kubernetes)
├── Makefile
├── README.md
└── .env.example
```

### Descrição das Pastas
- `api-gateway/`, `order-service/`, `payment-service/`, `notification-service/`: Cada serviço isolado, seguindo boas práticas de Go (pasta `cmd` para entrypoints, `internal` para lógica interna).
- `proto/`: Definições dos contratos gRPC.
- `deployments/`: Arquivos de orquestração (Docker Compose/K8s).
- `Makefile`: Comandos utilitários para build, test, run, etc.
- `.env.example`: Exemplo de variáveis de ambiente.

---

Este projeto serve como base para um portfólio profissional, demonstrando domínio de arquitetura distribuída, microsserviços, mensageria e boas práticas de desenvolvimento em Go.
