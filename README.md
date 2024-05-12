# Temperday

## Introdução

Este projeto é uma aplicação que fornece um serviço web para consultar informações da temperatura com base no CEP. A aplicação utiliza APIs externas para obter informações de clima e dados de endereço com base no CEP.

## Funcionalidades

### 1. Buscar por CEP
   * Rota: localhost:8080/{cep}
   * Método HTTP:GET
   * Exemplo de uso:
   `curl http://localhost:8000/12345678`
   * Exemplo de resposta: 
   ```json
      {
         "temp_c": 19,
         "temp_f": 66.2,
         "temp_k": 292.15
      }
   ```

## Configuração do Projeto
 - `APIKEY`:  É necessário uma apikey do [weatherapi](https://www.weatherapi.com/) para que o projeto funcione corretamente. Essa key pode ser setada dentro do aquivo .env no diretório **cmd**

## Execução
1. Certifique-se de ter configurado corretamente as variáveis de ambiente necessárias.
2. Tenha o docker instalado.
3. Execute o seguinte comando para fazer a construção da imagem: `docker build -t goexpert/temperday .`
4. Execute a imagem construída  na etapa anterior: `docker run goexpert/temperday`

# Cloud Run

- Deploy: Deploy disponível em [temperday](https://temperday-n2m7djaeia-uc.a.run.app/70070550)
