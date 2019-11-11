# TODO List

1. Create TestMain in `postgres` package that creates DB pool that all tests use so that there can never be exhaustion of connections, which are possible now due to all postgres tests being run on transactions and thus being parellelizable(sp?)
2. Return TraceID (and/or SpanID) in response headers so that client's unique ID can be used to debug integration tests