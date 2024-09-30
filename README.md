# Bad-Idea-as-a-Service

The Unprovider Capability Provider for wasmCloud.

Runs a pre-configured binary with arguments received from clients, returning the output to clients.

```
make
wash app deploy ./wadm.yaml

# call the test component
wash call unprovider-caller lxfontes:unprovider-example/caller.call

{
  "please": "indent",
  "this": "json",
  "forme": true
}

```
