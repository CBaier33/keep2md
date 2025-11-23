# keep2md

convert google keep notes from json format to markdown (for obsidian migration)

## usage

```
keep2md -f file.md
```

For all files: 
```bash
find ~/Downloads/Takeout/Keep -type f -name "*.json" -print0 | xargs -0 -I{} keep2md -f "{}"
mv *.md ~/Your/Obsidian/Vault

```

## installation

```
git clone https://github.com/cbaier22/keep2md
cd keep2md
go build && go install
```


## configuration

Adjust the `defaultTemplate` to configure how the resulting file will be created.

There are other attributes in the json data that can be parsed but this was good enough for me.
