# sortingIP

Скрипт для сортировки-упорядочивания списков ip с поддержкой разделителей диапазонов как в входном так и в выходном файле

## Install sortingIP
```bash
go install github.com/ihariv/sortingIP
```
## Using sortingIP
```bash
#help page
sortingIP --help

#custom input file and separator
sortingIP -in=input.txt -sep=- > output.txt

#custom input file, output file and separator 
sortingIP -in=input.txt -sep=- -out=output.txt
```
## Example input.txt

see in repository

## Exampe output
```bash
$ sortingIP
8.3.29.125-8.3.30.12
```