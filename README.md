# template-action

Template repo for creating new GitHub Actions using Go. Bootstraps fully functional action so you can focus on writing the code. Includes: 

* PR Qualification (on-push to the repo)
  * Go and YAML linting 
  * Static code analysis and vulnerability scanning
  * Unit test
* Release (on git tag in the repo)
  * All PR qualification test
  * Image build (scratch + single binary)
  * Image SBOM generation and attestation signing
  * SLSA provenance generation and validation 
  * Integration test in GitHub Actions action
* Repo hygiene 
  * Go, GitHub Actions, and Docker dependabot update configuration 
  * Scorecards analysis with in-repo Sarif report
  * Code test coverage report 

## template usage

To create a new repo, click the green `Use this Template` button and follow the wizard. When done, clone your new repo locally, and navigate into it:

```shell
git clone git@github.com:$GIT_HUB_USERNAME/$REPO_NAME.git
cd $REPO_NAME
```

Initialize your new repo. This will update all the references to your newly clone GitHub repository.

```shell
tools/init
```

When completed, commit and push the updates to your repository: 

```shell
git add --all
git commit -m 'repo init'
git push --all
```

> The above push will trigger the `on-push` flow. You can navigate to the `/actions` in your repo to see the status of that pipeline. 

#### trigger release pipeline

The canonical version of the entire repo is stored in [.version](.version) file. Feel free to edit it (by default: `v0.0.1`). When done, trigger the release pipeline:

> If you did edit the version, make sure to commit and push that change to the repo first. You can also use `make tag` to automate the entire process.

```shell
export VERSION=$(cat .version)
git tag -s -m "initial release" $VERSION
git push origin $VERSION
```

#### monitor the pipeline 

Navigate to `/actions` in your repo to see the status of that release pipeline. Wait until all steps (aka jobs) have completed (green). 

> If any steps fail, click on them to see the cause. Fix it, commit/push changes to the repo, and tag a new release to re-trigger the pipeline again.

#### review produced image

When successfully completed, that pipeline will create an image. Navigate to the your repo packages to review.

https://github.com/YOUR-USERNAME-OR-ORG/template-action/pkgs/container/action

The image is the line item tagged with version (e.g. `v0.0.1`). The other two OCI artifacts named with the image digest in the registry are signature (`.sig`) and attestation (`.att`).

You can now take the image digest and query sigstore transparency service (Rekor). Easiest way to do that is to use the Chainguard's [rekor-search-ui](https://github.com/chainguard-dev/rekor-search-ui). Here is the entry for [v0.0.1](https://rekor.tlog.dev/?hash=sha256:54c4d185322c87d05835f2f9ac72526ee5ada36a6145993adf87bd9c271334f5).

#### provenance verification  

Whenever you tag a release in the repo and an image is push to the registry, that image has an "attached" attestation in a form of [SLSA provenance (v0.2)](https://slsa.dev/provenance/v0.2). This allows you to trace that image all the way to its source in the repo (including the GitHub Actions that were used to generate it). That ability for verifiable traceability is called provenance. 

For example on how to verify the provenance of an image that was generated by the `on-tag` pipeline using cosign:

```shell
COSIGN_EXPERIMENTAL=1 cosign verify-attestation \
   --type slsaprovenance \
   --certificate-identity-regexp "^https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@refs/tags/v[0-9]+.[0-9]+.[0-9]+$" \
   --certificate-oidc-issuer https://token.actions.githubusercontent.com \
   $IMAGE_DIGEST
```

> The `COSIGN_EXPERIMENTAL` environment variable is necessary to verify the image with the transparency log until cosign v2 lands.

The terminal output will include the checks that were executed as part of the validation, as well as information about the subject (URI of the tag ref that triggered that workflow), with its SHA, name, and Ref.

```shell
Verification for ghcr.io/mchmarny/template-action@sha256:54c4d185322c87d05835f2f9ac72526ee5ada36a6145993adf87bd9c271334f5 --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - Existence of the claims in the transparency log was verified offline
  - The code-signing certificate was verified using trusted certificate authority certificates
Certificate subject:  https://github.com/slsa-framework/slsa-github-generator/.github/workflows/generator_container_slsa3.yml@refs/tags/v1.5.0
Certificate issuer URL:  https://token.actions.githubusercontent.com
GitHub Workflow Trigger: push
GitHub Workflow SHA: 45e3420e89478ccd0ffd97b8ee209eb5a5c59c69
GitHub Workflow Name: on_tag
GitHub Workflow Trigger mchmarny/template-action
GitHub Workflow Ref: refs/tags/v0.0.1
```

The output will also include JSON, which looks something like this (`payload` abbreviated): 

```json
{
   "payloadType": "application/vnd.in-toto+json",
   "payload": "eyJfdHl...V19fQ==",
   "signatures": [
      {
         "keyid": "",
         "sig": "MEUCIQCl+9dSv9f9wqHTF9D6L1bizNJbrZwYz0oDtjQ1wiqmLwIgE1T1LpwVd5+lOnalkYzNftTup//6H9i6wKDoCNNhpeo="
      }
   ]
}
```

The `payload` field (abbreviated) is the base64 encoded [in-toto statement](https://in-toto.io/) containing the predicate containing the GitHub Actions provenance:

```json
{
    "_type": "https://in-toto.io/Statement/v0.1",
    "predicateType": "https://slsa.dev/provenance/v0.2",
    "subject": [
        {
            "name": "ghcr.io/mchmarny/template-action",
            "digest": {
                "sha256": "54c4d185322c87d05835f2f9ac72526ee5ada36a6145993adf87bd9c271334f5"
            }
        }
    ],
    "predicate": {...}
}
```

## action usage 

Fow example on how to use the resulting actions in GitHub Actions see [.github/workflows/valid.yaml](.github/workflows/valid.yaml).

## disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.