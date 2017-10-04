Logentries Terraform Provider
=============================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

###### Powered by: https://www.terraform.io

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/dikhan/terraform-provider-logentries`

```sh
$ go get github.com/dikhan/terraform-provider-logentries
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/dikhan/terraform-provider-logentries
$ go build
```

Using the provider
------------------

Refer to the READMEs inside the [examples](https://github.com/dikhan/terraform-provider-logentries/examples) folder to 
see how to configure each resource provided by this terraform provider. 

Developing the Provider
-----------------------

To compile the provider, run `go build`. This will build the provider and put the provider binary in the current
`$GOPATH/src/github.com/dikhan/terraform-provider-logentries` directory.

```sh
$ go build
...
$ ls -la terraform-provider-logentries
...
```

In order to test the provider, you can simply run:

```sh
$ TF_ACC=1 LOGENTRIES_API_KEY="API_KEY" go test $(go list ./...) -timeout 120m -v
```
Expected output:

```
$ TF_ACC=1 LOGENTRIES_API_KEY="<API_KEY>" go test $(go list ./...) -timeout 120m -v
?       github.com/dikhan/terraform-provider-logentries [no test files]
=== RUN   TestLogentriesProvider
--- PASS: TestLogentriesProvider (0.00s)
=== RUN   TestAccLogentriesLog_Create
--- PASS: TestAccLogentriesLog_Create (8.81s)
=== RUN   TestAccLogentriesLog_Update
--- PASS: TestAccLogentriesLog_Update (11.08s)
=== RUN   TestAccLogentriesLogSets_Create
--- PASS: TestAccLogentriesLogSets_Create (0.98s)
=== RUN   TestAccLogentriesLogSets_Update
--- PASS: TestAccLogentriesLogSets_Update (1.60s)
=== RUN   TestAccLogentriesTags_Create
--- PASS: TestAccLogentriesTags_Create (13.36s)
=== RUN   TestAccLogentriesTags_Update
--- PASS: TestAccLogentriesTags_Update (19.71s)
PASS
ok      github.com/dikhan/terraform-provider-logentries/logentries      55.636s

```

Or specific tests can also be executed as follows:

```sh
$ TF_ACC=1 LOGENTRIES_API_KEY="<API_KEY>" go test github.com/dikhan/terraform-provider-logentries/logentries -run  ^TestAccLogentriesTags_Create$ -timeout 120m -v
```

The acceptance tests require a LOGENTRIES_API_KEY to be set. This env variable value will be used within the tests to 
successfully interact with the log entries api.

*Note: Acceptance tests create real resources and perform clean up tasks afterwards.*

Contributing
------------

- Fork it!
- Create your feature branch: git checkout -b my-new-feature
- Commit your changes: git commit -am 'Add some feature'
- Push to the branch: git push origin my-new-feature
- Submit a pull request :D

Authors
-------

Daniel I. Khan Ramiro

See also the list of [contributors](https://github.com/dikhan/terraform-provider-logentries/graphs/contributors) who 
participated in this project.