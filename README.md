# citibank-hackathon
projeto hackathon citibank 

<img width="742" alt="img1" src="https://github.com/giocarvalho07/citibank-hackathon/assets/52415453/814b6c85-e8e4-4886-bdbf-201ce05be6dc">

<br>
<br>

<p>
    Arquitetura do nosso projeto é baseada em upload, tratamento e geração de arquivos, onde temos como "base" o padrão CNAB 750
</p>


![](https://github.com/giocarvalho07/citibank-hackathon/issues/1)

# Tecnologias utilizadas 

## Skills

Badge | URL
------------ | -------------
<img src="https://img.shields.io/badge/HTML-239120?style=for-the-badge&logo=html5&logoColor=white" /> | `https://img.shields.io/badge/HTML-239120?style=for-the-badge&logo=html5&logoColor=white`
<img src="https://img.shields.io/badge/CSS-239120?style=for-the-badge&logo=css3&logoColor=white" /> | `https://img.shields.io/badge/CSS-239120?&style=for-the-badge&logo=css3&logoColor=white`
<img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black" /> | `https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black`
<img src="https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white" /> | `https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white`
<img src="https://img.shields.io/badge/Saas-CC6699?style=for-the-badge&logo=sass&logoColor=white" /> | `https://img.shields.io/badge/Sass-CC6699?style=for-the-badge&logo=sass&logoColor=white`
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> | `https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white`
<img src="https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white" /> | `https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white`
<img src="https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white" /> | `https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white`
<img src="https://img.shields.io/badge/Netlify-00C7B7?style=for-the-badge&logo=netlify&logoColor=white" /> | `https://img.shields.io/badge/Netlify-00C7B7?style=for-the-badge&logo=netlify&logoColor=white`


# como instalar GoLang


Acesse a página de downloads da Golang e baixe o instalador da Golang para Windows. No momento em que realizo a escrita desse artigo, a versão atual da linguagem 
é a 1.19. https://go.dev/dl/

Na tela seguinte será a tela de configuração do local de instalação da Golang, por padrão a linguagem é instalada na pasta C:\Program Files\Go\ ou C:\Program Files (x86)\Go\. Caso queira alterar o local de instalação basta clicar no botão “Change” e selecionar o novo local de instalação desejado, mas recomendo seguir com o padrão e simplesmente clicar no botão “Next”.

Para verificar se a instalação foi realizada com sucesso podemos abrir algum terminal do Windows (Prompt de Comando ou PowerShell) e digitar o seguinte comando:

### comando para verificar se o GoLang está na máquina

```bash
$ go version
```

## Rodando a aplicação local

- Para executar o projeto Go (Golang) direto no terminal
```bash
$ go run main.go
```

## Testes local

Apos o comando, seu terminal irá ficar semelhante a imagem, com o projeto em ambiente local:8080

```bash
$ go run main.go
```

![local](https://github.com/giocarvalho07/citibank-hackathon/assets/52415453/09e4c895-6d61-4212-9522-33c2de4d101c)



### Retorno Local

Após teste local, você pode testar esta Url no browser ou Postman

```bash
$ http://localhost:8080/cnab
```

<img width="448" alt="postmanXC" src="https://github.com/giocarvalho07/citibank-hackathon/assets/52415453/1f6e0b6e-d6bb-451e-b030-7bc71db08947">




### Json de retorno 

```bash
$ [
    {
        "TipoOperacaoRemessa": "001",
        "LiteralRemessa": "8566994",
        "CodigoServico": "02",
        "LiteralServico": "L",
        "ISPBParticipante": "123",
        "TipoPessoaRecebedor": "896",
        "CPForCNPJ": "0987656643",
        "Agencia": "0258",
        "Conta": "9633654",
        "TipoConta": "green",
        "ChavePix": "2966657",
        "DataGeracao": "09/05/2023",
        "CodigoConvenio": "741",
        "ExclusivoPSPRecebedor": "00007",
        "NomeRecebedor": "Carros SA",
        "Bancos": "Caixa",
        "NumeroSequenciaRemessa": "0025014",
        "VersaoArquivo": "089",
        "TipoRegistro": "001",
        "Identificador": "green",
        "Tipo": "FG",
        "TipoCobranca": "000585",
        "CodOcorrencia": "001",
        "TimesTampExpiracao": "98745550",
        "DataVencimento": "02/09/2023",
        "ValidadeAposVencimento": "985245",
        "ValorOriginal": "001",
        "TipoPessoaDevedor": "PL",
        "CPForCNPJDevedor": "824942636426",
        "NomeDevedor": "Banespa JH Bank",
        "CampoTextoAoPagador": "001",
        "ValorTotal": "895",
        "QTdRegistros": "750",
        "NumeroSequencial": "001",
        "StatusProjeto": "Processado"
    },
    {
        "TipoOperacaoRemessa": "002",
        "LiteralRemessa": "8566994",
        "CodigoServico": "02",
        "LiteralServico": "L",
        "ISPBParticipante": "123",
        "TipoPessoaRecebedor": "896",
        "CPForCNPJ": "8799222159658",
        "Agencia": "0258",
        "Conta": "9633654",
        "TipoConta": "red",
        "ChavePix": "2966657",
        "DataGeracao": "09/05/2023",
        "CodigoConvenio": "741",
        "ExclusivoPSPRecebedor": "00007",
        "NomeRecebedor": "carla",
        "Bancos": "BB",
        "NumeroSequenciaRemessa": "0025014",
        "VersaoArquivo": "089",
        "TipoRegistro": "002",
        "Identificador": "red",
        "Tipo": "FG",
        "TipoCobranca": "000585",
        "CodOcorrencia": "001",
        "TimesTampExpiracao": "98745550",
        "DataVencimento": "02/09/2023",
        "ValidadeAposVencimento": "985245",
        "ValorOriginal": "001",
        "TipoPessoaDevedor": "PL",
        "CPForCNPJDevedor": "824942636426",
        "NomeDevedor": "Mario",
        "CampoTextoAoPagador": "001",
        "ValorTotal": "895",
        "QTdRegistros": "750",
        "NumeroSequencial": "001",
        "StatusProjeto": "Processado"
    }
    ]
```


#### Escalabilidade do projeto

Realizamso testes para o projeto suportar até 2000 requisições por minuto, prevendo tolerância a falhas e atingindo o objeto de 100mil requisições por hora de contingênciamento

<img width="335" alt="request" src="https://github.com/giocarvalho07/citibank-hackathon/assets/52415453/26cc4f1a-7252-43ba-9d71-61ad4cc15a6d">


```bash
EXAMPLE:
    $ [GIN] 2023/06/18 - 12:15:05 | 200 |        86.5µs |       127.0.0.1 | GET      "/cnab"
```

