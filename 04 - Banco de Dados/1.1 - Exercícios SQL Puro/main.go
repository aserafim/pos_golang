// 1 - Inserir múltiplos produtos de uma vez:
// Crie uma função que receba um slice de Product e
// insira todos os produtos de uma só vez no banco
// de dados usando uma transação.

/*
Paginação de resultados:

Modifique a função selectAllProducts para retornar um número limitado de produtos por vez (ex.: 10 produtos), implementando paginação.
Filtragem de produtos:

Adicione um filtro por preço e nome. Crie uma função que selecione produtos com base em um intervalo de preços e/ou uma parte do nome.
Atualização em massa:

Crie uma função que permita atualizar o preço de todos os produtos que estejam dentro de um determinado intervalo de preço (ex.: aumentar o preço de todos os produtos que custam entre R$ 100 e R$ 500 em 10%).
Logs de operações:

Modifique suas funções para que, a cada operação de inserção, atualização ou exclusão de produto, seja criado um registro em uma tabela de logs. Esse registro pode conter o ID do produto, a operação realizada e a data/hora.
Exclusão em lote:

Implemente uma função que exclua todos os produtos cujo preço esteja acima de um valor especificado.
Validação de dados antes da inserção:

Adicione uma camada de validação para garantir que o nome do produto não seja vazio e o preço seja maior que zero antes de realizar a inserção ou atualização.
*/
