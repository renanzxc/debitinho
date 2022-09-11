# 💰 Debitinho 📄 (🔨 WIP)
> CLI para visualizar e validar arquivos de remessa/retorno de Débito Automático (Débito em Conta) na versão 6 do Febraban.


## Executando exemplo:
```bash
go build
./debitinho exibir mock/files/remessa.txt 
```
Saída:
```json
{
    "codigo_do_registro": "A",
    "codigo_de_remessa": "1",
    "codigo_do_convenio": "00000000000000000000",
    "nome_da_empresa": "EMPRESATESTE        ",
    "codigo_do_banco": "111",
    "nome_do_banco": "BANCOTESTE          ",
    "data_de_geracao": "20220413",
    "nsa": "000001",
    "versao_do_layout": "06",
    "identificacao_do_servico": "DEBITO AUTOMATICO",
    "reservado_para_o_futuro": "                                                    "
}
{
    "codigo_do_registro": "Z",
    "total_de_registros_do_arquivo": "000002",
    "valor_total_dos_registros_do_arquivo": "00000000000000000",
    "reservado_para_o_futuro": "                                                                                                                              "
}

``` 
### To Do:
- [X] Adicionar parse do arquivo
- [ ] Adicionar parse das linhas
  - [X] A, Z
  - [ ] B, C, D, E, F, H, I, J, K, L, T, X
- [ ] Adicionar validações dos arquivos
- [ ] Adicionar validações das linhas
  - [ ] A, B, C, D, E, F, H, I, J, K, L, T, X, Z
- [ ] Adicionar modos de visualização dos arquivos
  - [ ] JSON
  - [ ] Terminal
  - [ ] HTML
- [ ] Gerar arquivo de retorno para arquivo de remessa
- [ ] Adicionar configuração para validação personalizada por banco
