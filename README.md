# [reconciler-runtime](https://github.com/vmware-labs/reconciler-runtime) testbed

```shell
❯ go test ./...
?       github.com/mamachanko/rr-test   [no test files]
?       github.com/mamachanko/rr-test/api/v1alpha1      [no test files]
--- FAIL: TestApplyMyHTTPProxy (0.01s)
    --- FAIL: TestApplyMyHTTPProxy/Updates_HTTPProxy_and_removes_TLS_Secret (0.00s)
        thing_reconciler_test.go:156: Unexpected update for config "default" (-expected, +actual):   &v1.HTTPProxy{
                ... // 1 ignored field
                ObjectMeta: {Name: "test-parent", Namespace: "test-ns", OwnerReferences: {{APIVersion: "things.mamachanko.com/v1alpha1", Kind: "Thing", Name: "test-parent", Controller: &true, ...}}},
                Spec: v1.HTTPProxySpec{
                        VirtualHost: &v1.VirtualHost{
                                Fqdn:            "test-fqdn.example.com",
            -                   TLS:             nil,
            +                   TLS:             &v1.TLS{SecretName: "test-tls"},
                                Authorization:   nil,
                                CORSPolicy:      nil,
                                RateLimitPolicy: nil,
                        },
                        Routes:   nil,
                        TCPProxy: nil,
                        ... // 2 identical fields
                },
                Status: {},
              }
FAIL
FAIL    github.com/mamachanko/rr-test/controllers       0.424s
FAIL

```

