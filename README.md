# action

WIP: Template for GitHub Actions 

## inputs

* `file` - (required) File path
* `required` - (optional) Whether required or ot


## usage

> Make sure to use the latest tag release (e.g. `v0.0.1`)

```yaml
- uses: mchmarny/action@v0.0.19
  with:
    project: ${{ inputs.target_project }}
    digest: ${{ inputs.image_digest }}
    file: ${{ inputs.report_path }}
    format: trivy
```

> Fully working example can be found in [.github/workflows/valid.yaml](../../.github/workflows/valid.yaml).


## Disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.
