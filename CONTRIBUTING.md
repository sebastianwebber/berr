# Updating docs

to update the docs, run:

```shell
nix-shell -p gomarkdoc --run "gomarkdoc -v" 
```

to update the gif file, run:

```shell
nix-shell -p vhs --run "vhs examples.tape" 
```