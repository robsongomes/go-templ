# Live reloading nativo

No live reloading nativo a aplicação irá reiniciar sempre que um arquivo templ for modificado

```bash
$ templ generate -watch -proxy="http://localhost:3000/" -cmd="go run ."
```

# Live reloading com air

Com o [air](https://github.com/cosmtrek/air), a aplicação irá reiniciar sempre que os arquivos de template ou outros arquivos .go forem modificados.

## air.toml

```conf
cmd = "go build -o ./tmp/main . && templ generate -notify-proxy"
exclude_regex = ["_test.go",".*_templ.go"]
include_ext = ["go", "tpl", "tmpl", "html", "templ"]
```

## Execução

Para usar essa arbodagem, 2 terminais deverão ser usados, um para o air e outro para o proxy do templ

```bash
$ air
```


```bash
$ templ generate -watch -proxy="http://localhost:3000/"
```