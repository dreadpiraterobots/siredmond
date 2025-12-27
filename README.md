# SirEdmond
A fast and efficient way to interact with Microsoft security advisory data.

## Data sources
### CVRF
Microsoft has published a significant amount of security advisory data via the MSRC CVRF since around 2017. This is a welcome replacement for earlier MS<yy>-<nnnn> advisories.

### Other data sources
Currently, SirEdmond does not know about data sources other than the MSRC CVRF. In future, we could add support for other sources including:
- CSAF: potentially a more fully-featured replacement for CVRF, but the MSRC implementation currently lacks historic info vs. CVRF.
- KB pages & release notes: Microsoft has an unfortunate habit of publishing a variety of essential information outside of the CVRF/CSAF. Examples include remediation versions for specific Microsoft 365 update channels, and .NET Framework file version remediation indices.

## Project standards
### Branch strategy, CI policy, and testing philosophy
`main` is protected. This is a small enough project that I don't anticipate needing a `release` branch, so anything merged to `main` is in production.

We will enforce standard Go CI checks: `go build`, `go vet`, and unit testing.
