# API MERCA FACIL

API de cadastro de informção utilizando uma forma de multi-tenant. Aplicação configurada para utilizar-se de dois banco dinstintos, postgres e mysql.

# Execução

* Para facilitar o processo, é necessário que você tenha o docker e docker-compose funcionando corretamente em sua máquina.
* Passo 1: ```sudo docker-compose up postgresql -d```
* Passo 2: ```sudo docker-compose up mysql -d```
* Passo 3: Certifique-se que ambos os banco de dados estão em pleno funcionamento
* Passo 4: Execute: ```sudo docker-compose up api-merca-facil -d``` para que seja inicializado a api
* Passo 5: Para que seja possivel a utilização é necessário criar um usuário. Que será utilizado para autenticação. Essa autenticação é com base no token JWT
  * ```json:
    {
      "email": "varejao@varejao.com",
      "senha": "1234",
      "banco_dados": "POSTGRES_SQL",
      "is_customizavel": false,
      "cliente_name": "varejao"
    }
    ```
    * email: Será o email necessário para uso do login
    * senha: Será a senha necessário para uso do login
    * banco_dados: Será o banco de dados em que será salvo os dados para o usuario criado. Existe duas opção: POSTGRE_SQL ou MY_SQL
    * is_customizavel: É utilizado para que seja formatado o nome e número do telefome enviado
    * cliente_name: é o nome do cliente (usuário)
* Passo 6: Agora é hora do login
  * ```json:
      {
        "email": "varejao@varejao.com",
        "senha": "1234"
      }
    ```
  * O retorno será: 
    * ```json:
        {
          "id": "1",
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJiYW5jb0RhZG9zIjoiUE9TVEdSRVNfU1FMIiwiZXhwIjoxNjU3NjA4NzAyLCJpc0N1c3RtaXphdmVsIjpmYWxzZSwidXN1YXJpb0lkIjoxfQ.RBrBTWDUh0xXJtxStAYrLBDi4F9TeO3Po-WFM8rVDOQ"
        }
      ```
    * Utilizar o token no bearer para autenticação
* Passo 7: Para cadastra mais de um contato ao mesmo tempo, utilizar da rota: ```http://localhost:8000/api/v1/telefones/contatos```
  * Json de exemplo se encontra nos arquivos json contido neste projeto