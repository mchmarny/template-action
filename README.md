# action

WIP: Template for GitHub Actions 

## inputs

* `project` - (required) GCP Project ID
* `digest` - (required) Image digest
* `file` - (required) Path to the vulnerability file
* `format` - (required) Format of the vulnerability file


## usage

> Make sure to use the latest tag release (e.g. `v0.0.1`)

```yaml
uses: mchmarny/action@v0.2.14
with:
  project: ${{ env.PROJECT_ID }}
  digest: ${{ steps.build.outputs.digest }}
  file: ${{ steps.scan.outputs.output }}
  format: ${{ steps.scan.outputs.format }}
```

> Fully working example can be found in [.github/workflows/validate.yaml](../../.github/workflows/validate.yaml).


## Disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.
