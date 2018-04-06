# sortingIP

The script generates IP address lists with support for ranges in both input and output files.

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

## exampe output
```bash
$ sortingIP
8.3.29.125-8.3.30.12
```