# Apache Kafka

### O que é o Apache Kafka?

* É uma plataforma distribuída de streaming de eventos open-source que é utilizada por milhares de empresas para ula alta performance em _pipeline de dados_, _stream de analytics_ integração de de dados e aplicações de missão crítica

* Site: [clique aqui](https://kafka.apache.org/)

### O mundo dos eventos

* Cada dia mais precisamos processar mais e mais eventos em diversos tipos de plataforma. Desde sistema que precisam se comunicar, _devices para IOT_, monitoramento de aplicações, sistemas de alarmes, etc

### Kafka e seus "super poderes"

* Altíssimo _throughput_

* Latência extremamente baixa (até 2ms)

* Escalável

* Armazenamento

* Alta disponibilidade

* Se conecta com quase todas as linguagens

* Bibliotecas prontas para as mais diversas tecnologias

* Ferramentas open-source

### Empresas usando Kafka

* Linkedin (criador do Kafka)

* Netflix

* Uber

* Twitter

* Dropbox

* Spotify

* Paypal
 
* Bancos... (**principalmente**)

### Conceitos e dinâmica básica de funcionamento

![](./assets/representacao-kafka.png)

* **Producer**

  * **Definição**: responsável pelo envio de mensagens(eventos) para o Kafka

  * **Representação**

    ![](./assets/representacao-consumer.png)

    > Consumer lê os dados contidos dentro das partições

  * **Consumer Groups**

    ![](./assets/representacao-consumer-groups.png)

    > Aumenta a vazão dos dados (mensagens)

    > **IMPORTANTE**: 1 partição para 1 consumer (Kafka não permite ter 1 partição para 2 ou + consumer)

* **Broker**: nome dado ao _cluster_

* **Consumer** (aplicação): responsável pela leitura das mensagens(eventos) armazenadas no Kafka

> **OBS**: o Kafka **não** envia mensagem para o consumer

* **Tópicos** 
  
  * **Definição**: é o canal de comunicação responsável por receber e disponibilizar os dados enviados para o Kafka

  * **Representação**

    ![](./assets/representacao-topicos.png)

  * **Tópicos são como "Logs"**

    ![](./assets/exemplo-topicos.png)

    > **IMPORTANTE**: os tópicos são **armazenados em disco**, e não em memória

* **Registro**

  * **Definição**: composto por alguns metadados (headers, key, value, timestamp)

  * **Representação**

    ![](./assets/representacao-registro.png)

  * **Partições**

    * **Definição**: cada tópico pode ter uma ou mais partições para conseguir garantir a distribuição e resiliência de seus dados

    * **Representação**

      ![](./assets/representacao-particao.png)

      > **OBS**: quanto mais partições menores são os riscos

    * **Partições distribuídas**

      ![](./assets/representacao-particoes-distribuidas.png)

      > Cópias de cada partição distribuídas em vários _brokers_ -> evita perda de dados (mensagens)