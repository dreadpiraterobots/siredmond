# SirEdmond
A fast and efficient way to interact with Microsoft security advisory data.

If you want to know:
- How many critical CVEs did Microsoft publish on the most recent Patch Tuesday?
- Which vulns did Microsoft publish in 2023 and tag as exploited in the wild?
- How often does Microsoft assign a proprietary severity to a vulnerability which differs from the severity suggested by the CVSS vector?
- Who is most often credited with reporting critical vulnerabilities in the Windows kernel?

or if you want to:
- Extract information from the CVRF into CSV, JSON, HTML etc.
- Explore Microsoft security advisory data without reading XML files or relying on the MSRC web UI.

then SirEdmond might just be what you're looking for.

> [!IMPORTANT]
> Some of these aren't implemented yet. I'll remove this note once it's no longer needed.

## Data sources
### CVRF
Microsoft has published a significant amount of security advisory data via the MSRC CVRF since around 2017. This is a welcome replacement for the earlier `MS<yy>-<nnnn>` advisories.

The CVRF contains a lot of useful information, although full remediation information is sometimes missing for some products (e.g. Office/365, Edge).

### Other data sources
Currently, SirEdmond does not know about data sources other than the MSRC CVRF. In future, we could add support for other sources including:
- CSAF: potentially a more fully-featured replacement for CVRF, but the MSRC implementation currently lacks historic info vs. CVRF.
- KB pages & release notes: Microsoft has an unfortunate habit of publishing a variety of essential information outside of the CVRF/CSAF. Examples include remediation versions for specific Microsoft 365 update channels, and .NET Framework file version remediation indices.
- Other Microsoft data sources (e.g. the official-but-undocumented Edge version API, Office updaye channel version from the Office CDN).

## Project standards
### Branch strategy, CI policy, and testing philosophy
`main` is protected. This is a small enough project that I don't anticipate needing a `release` branch, so anything merged to `main` is in production.

We will enforce standard Go CI checks: `go build`, `go vet`, and unit testing.
