<div align="center">
  <h1>API REST em Go</h1>
  <img src="https://img.shields.io/static/v1?label=Go&message=language&color=blue&style=rounded&logo=GO"/>
  <img src="https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=rounded"/>
  <!-- github workflows -->
  <a href="https://github.com/walber-vaz/api-go/actions/workflows/main.yml">
    <img src="https://github.com/walber-vaz/api-go/actions/workflows/main.yml/badge.svg?branch=main"/>
  </a>
  <!-- version -->
  <img src="https://img.shields.io/github/v/release/walber-vaz/api-go?include_prereleases&style=rounded"/>
  <!-- code size -->
  <img src="https://img.shields.io/github/languages/code-size/walber-vaz/api-go?style=rounded"/>
  <!-- repo size -->
  <img src="https://img.shields.io/github/repo-size/walber-vaz/api-go?style=rounded"/>
  <!-- issues -->
  <img src="https://img.shields.io/github/issues/walber-vaz/api-go?style=rounded"/>
  <!-- pull requests -->
  <img src="https://img.shields.io/github/issues-pr/walber-vaz/api-go?style=rounded"/>
  <!-- last commit -->
  <img src="https://img.shields.io/github/last-commit/walber-vaz/api-go?style=rounded"/>
  <!-- contributors -->
  <img src="https://img.shields.io/github/contributors/walber-vaz/api-go?style=rounded"/>
   <!-- stars -->
  <img src="https://img.shields.io/github/stars/walber-vaz/api-go?style=social"/>
  <!-- forks -->
  <img src="https://img.shields.io/github/forks/walber-vaz/api-go?style=social"/>
  <!-- watchers -->
  <img src="https://img.shields.io/github/watchers/walber-vaz/api-go?style=social"/>
</div>

<a id="tabela-de-conteúdos"></a>

## :pushpin: Tabela de conteúdos

<!--ts-->
* [Sobre](#Sobre)
* [Status do Projeto](#status-do-projeto)
* [Features](#features)
* [Demonstração da Aplicação](#demonstração-da-aplicação)
* [Pré-requisitos](#pré-requisitos)
* [Como rodar a aplicação](#como-rodar-a-aplicação)
* [Tecnologias](#tecnologias)
* [Autor](#autor)
* [Como contribuir](#como-contribuir)
* [Licença](#licença)
<!--te-->

<a id="Sobre"></a>

## :bookmark: Sobre

<!-- volta para menu -->
:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

API REST em Go é uma aplicação para cadastro de jobs e freelancers. Esta em desenvolvimento e será utilizada para estudo de Go.

<a id="status-do-projeto"></a>

## :construction_worker: Status do Projeto

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

<h4 align="center">
  🚧  API REST em Go 🚀 Em construção...  🚧
</h4>

<a id="features"></a>

## :heavy_check_mark: Features

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

* [x] Setup inicial do projeto
* [x] Configuração do banco de dados
* [x] Configuração do servidor
* [x] Configuração do roteamento
* [ ] Cadastro de jobs
* [ ] Cadastro de freelancers
* [ ] Cadastro de usuários
* [ ] Autenticação de usuários
* [ ] Autorização de usuários
* [ ] Testes unitários
* [ ] Testes de integração

<a id="demonstração-da-aplicação"></a>

## :computer: Demonstração da Aplicação

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

A aplicação ainda esta em desenvolvimento.

<a id="pré-requisitos"></a>

## :warning: Pré-requisitos

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

Antes de começar, é necessário você ter instalado em sua máquina as seguintes ferramentas:

* [Git](https://git-scm.com)
* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [PostgreSQL](https://www.postgresql.org/)
* [pgAdmin](https://www.pgadmin.org/) (opcional)
* [Insomnia](https://insomnia.rest/) (opcional)
* [Make](https://www.gnu.org/software/make/)

<a id="como-rodar-a-aplicação"></a>

## :rocket: Como rodar a aplicação

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

```bash
# Clone este repositório
$ git clone https://github.com/walber-vaz/api-go.git

# Acesse a pasta do projeto no terminal/cmd
$ cd api-go

# Execute o comando para subir os containers em modo dev
$ make run

# O servidor inciará na porta:9000 - acesse http://localhost:9000
```

<a id="rodando-a-aplicação-web"></a>

## :hammer: Tecnologias

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

As seguintes ferramentas foram usadas na construção do projeto:

* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [PostgreSQL](https://www.postgresql.org/)
* [pgAdmin](https://www.pgadmin.org/)
* [Insomnia](https://insomnia.rest/)
* [GORM](https://gorm.io/)
* [JWT](https://jwt.io/)
* [Gin](https://gin-gonic.com/)

<a id="tecnologias"></a>

## :woman_technologist: Autor

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

Feito por Walber Vaz 👋🏽 Entre em contato!

[![Linkedin Badge](https://img.shields.io/badge/-Walber-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/walber-vaz/)](https://www.linkedin.com/in/walber-vaz/)
[![Github Badge](https://img.shields.io/badge/-Walber-000?style=flat-square&logo=Github&logoColor=white&link=walbervaz)](https://github.com/walber-vaz)

<a id="como-contribuir"></a>

## :recycle: Como contribuir para o projeto

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

1. Faça um **fork** do projeto.
2. Crie uma nova branch com as suas alterações: `git checkout -b my-feature`
3. Salve as alterações e crie uma mensagem de commit contando o que você fez: `git commit -m "feature: My new feature"`
4. Envie as suas alterações: `git push origin my-feature`

> Em caso de dúvidas acesse [esse guia de como contribuir no GitHub](./CONTRIBUTING.md)

<a id="licença"></a>

## :memo: Licença

:point_up_2: <a href="#tabela-de-conteúdos">Voltar ao menu</a>

Este projeto esta sobe a licença [MIT](./LICENSE).
