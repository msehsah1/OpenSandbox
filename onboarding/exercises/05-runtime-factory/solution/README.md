# Solution steps

1. Normalize with `.lower()`.
2. Use an explicit registry dict — that is how the real factory is extended.
3. The kubernetes branch returns composite, not `"kubernetes"`. That is the
   whole lesson: FastSandbox is not `runtime.type=fsb`.
