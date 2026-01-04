### Summary

PoC for a panic in code signature parsing that could be triggered by malformed Mach‑O files. If a CodeSignature blob (e.g., the entitlements blob) declared a length smaller than its header size, the parser computed a negative payload length and make panicked.

### Use
Test with [`panic.go`](panic.go).

#### Patched with [PR #80](https://github.com/blacktop/go-macho/pull/80)

### Credit
Jacob Prezant