# [reconciler-runtime](https://github.com/vmware-labs/reconciler-runtime) testbed

```shell
$ go test ./...
?       github.com/mamachanko/rr-test   [no test files]
?       github.com/mamachanko/rr-test/api/v1alpha1      [no test files]
--- FAIL: TestApplyMySecret (0.00s)
    --- FAIL: TestApplyMySecret/Does_not_update_my_secret (0.00s)
        thing_reconciler_test.go:98: Extra recorded event: {{Thing things.mamachanko.com/v1alpha1} test-ns/test-parent Warning CreationFailed Failed to create Secret "test-parent": secrets "test-parent" already exists}
        thing_reconciler_test.go:98: Extra create: testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"test-ns", Verb:"create", Resource:schema.GroupVersionResource{Group:"", Version:"v1", Resource:"Secret"}, Subresource:""}, Name:"", Object:(*v1.Secret)(0xc0000a6500)}
FAIL
FAIL    github.com/mamachanko/rr-test/controllers       0.389s
FAIL
```

