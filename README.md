# redirgo

A very simple link redirector wirtten in Go.

# Build and run
```bash
go build -o bin/api ./src/
./bin/api
```

# Container
```bash
podman build -t redirgo .
podman run --rm -it -p 8080:8080 redirgo
```

Accessing `localhost:8080/source` will redirect you to `https://github.com/vkhashimoto/redirgo`.

# Config
An example file is available in `config/links.toml`. 
Each table `[<host>]` is a host that you can use, and each key/value pair is a match of a path and where to redirect to.

```toml
["vkha.sh"]
"git" = "https://github.com/vkhashimoto"
```

It's possible to have a fallback redirection. If the file doesn't have a **matching host**, it will try to find a link in the fallback `[*]` host.

```toml
["*"]
"github" = "https://github.com"
```

# Pages
- `public/404.html` is rendered when a link is not found.
- `public/index.html` is rendered when visiting the root page `localhost:8080/`
