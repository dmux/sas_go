# sas_go

## Visão Geral

Este projeto tem como objetivo desenvolver um interpretador open-source para a linguagem de programação SAS (Statistical Analysis System) como um esforço acadêmico. O interpretador facilitará a execução de código SAS básico, permitindo que os usuários realizem tarefas de manipulação de dados, análises estatísticas e outras funcionalidades comumente associadas ao SAS.

## Objetivos

- **Compreender a Sintaxe do SAS**: O objetivo principal é analisar e executar construções fundamentais do código SAS, como etapas DATA e passos PROC, permitindo que os usuários adquiram conhecimentos sobre programação em SAS.


- **Propósito Educacional**: Este projeto serve como uma experiência prática no design e implementação de linguagens de programação, com foco específico na análise e interpretação de uma linguagem bem estabelecida como o SAS.


- **Contribuição Open Source**: Ao tornar este projeto open-source, incentivamos a colaboração e contribuições da comunidade para melhorar e expandir as capacidades do interpretador.

## Por que Go?

- **Desempenho**: Go é conhecido por sua execução eficiente e baixa latência, tornando-o adequado para desenvolver um interpretador que processa código SAS rapidamente.


- **Simplicidade e Legibilidade**: A sintaxe limpa do Go promove legibilidade e manutenibilidade, o que é crucial para o desenvolvimento colaborativo e futuras melhorias.


- **Concorrência**: O suporte integrado do Go para programação concorrente permite a implementação potencial de processamento multithread de código SAS, melhorando o desempenho para conjuntos de dados maiores.

## Por que ANTLR?

- **Framework de Análise Poderoso**: ANTLR (Another Tool for Language Recognition) é uma ferramenta robusta para gerar analisadores, fornecendo extensos recursos que simplificam a criação de gramáticas de linguagem e interpretadores.


- **Definição de Gramática**: ANTLR nos permite definir a sintaxe da linguagem SAS de forma clara e concisa, permitindo atualizações e modificações fáceis à medida que expandimos as capacidades do interpretador.


- **Suporte a Múltiplas Linguagens**: Embora nossa implementação se concentre no Go, o ANTLR pode gerar analisadores para várias linguagens, oferecendo flexibilidade na escolha das linguagens de implementação no futuro.- 

- **Suporte da Comunidade**: O ANTLR possui uma grande base de usuários e extensa documentação, facilitando a solução de problemas e o aprendizado.

