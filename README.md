
# Microservices GO

![Fluxograma](/.github/fluxograma-trabalho-pos.png "Fluxograma das aplicações").
Obs: No fluxograma contém apenas 3 microserviços pelo fato de ser o desafio solicitado para a conclusão do trabalho da pós graduação, os outros foram criados pelo próprio professor na aula. 

#### Comandos para executar o projeto
```
cp ./ticket-create/.env.dist ./ticket-create/.env
cp ./notification/.env.dist ./notification/.env

docker-compose up -d --build
```