# ArgoCD Application Sync Action

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Find%20and%20Replace-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/OpsVerseIO/argocd-sync)

A GitHub action that syncs ArgoCD applications.

## Usage

### Example workflow

This example syncs an ArgoCD application.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Sync ArgoCD Application
        uses: OpsVerseIO/argocd-sync@0.2.0
        with:
          address: ${{ secrets.ARGOCD_SERVER }}
          token: ${{ secrets.ARGOCD_TOKEN }}
          action: sync
          appName: "my-example-app"
          disableTlsVerification: "false" # Default is false. Only enable this if ArgoCD doesn't have TLS / has self signed certificate / you see any sort of x509 errors
```

### Inputs

| Input                    | Description                            |
|--------------------------|----------------------------------------|
| `address`                | ArgoCD server address.                 |
| `token`                  | ArgoCD Token.                          |
| `action`                 | ArgoCD Action i.e. sync.               |
| `appName`                | Application name to execute action on. |
| `disableTlsVerification` | Skip TLS validation.                   |

## Examples

### Sync Application

You can sync ArgoCD application after building an image etc.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Sync ArgoCD Application
        uses: OpsVerseIO/argocd-sync@0.2.0
        with:
          address: ${{ secrets.ARGOCD_SERVER }}
          token: ${{ secrets.ARGOCD_TOKEN }}
          action: sync
          appName: "my-example-app"
          disableTlsVerification: "false" # Default is false. Only enable this if ArgoCD doesn't have TLS / has self signed certificate / you see any sort of x509 errors
```
