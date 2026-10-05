# Solution steps

1. `HTTPException(status_code=..., detail={"code": ..., "message": ...})`.
2. Custom handler: if `detail` is a dict, `JSONResponse(content=detail)`.
3. Keep codes stable strings (`INVALID_PARAMETER`). The real server uses
   `SandboxErrorCodes`.

SDKs map these codes into typed exceptions — changing a code is a contract
change.
