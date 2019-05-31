Insight Terraform Provider
=============================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

###### Powered by: https://www.terraform.io

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/dikhan/terraform-provider-insight`

```sh
$ go get github.com/dikhan/terraform-provider-insight
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/dikhan/terraform-provider-insight
$ go build
```

Using the provider
------------------

Refer to the READMEs inside the [examples](https://github.com/dikhan/terraform-provider-insight/examples) folder to
see how to configure each resource provided by this terraform provider.

Developing the Provider
-----------------------

To compile the provider, run `go build`. This will build the provider and put the provider binary in the current
`$GOPATH/src/github.com/dikhan/terraform-provider-insight` directory.

```sh
$ go build
...
$ ls -la terraform-provider-insight
...
```

In order to test the provider, you can simply run:

```sh
$ TF_ACC=1 INSIGHT_API_KEY="API_KEY" go test $(go list ./...) -timeout 120m -v
```
Expected output:

```
$ TF_ACC=1 INSIGHT_API_KEY="<API_KEY>" go test $(go list ./...) -timeout 120m -v
?       github.com/dikhan/terraform-provider-insight [no test files]
=== RUN   TestInsightProvider
--- PASS: TestInsightProvider (0.00s)
=== RUN   TestAccInsightLog_Create
--- PASS: TestAccInsightLog_Create (8.81s)
=== RUN   TestAccInsightLog_Update
--- PASS: TestAccInsightLog_Update (11.08s)
=== RUN   TestAccInsightLogSets_Create
--- PASS: TestAccInsightLogSets_Create (0.98s)
=== RUN   TestAccInsightLogSets_Update
--- PASS: TestAccInsightLogSets_Update (1.60s)
=== RUN   TestAccInsightTags_Create
--- PASS: TestAccInsightTags_Create (13.36s)
=== RUN   TestAccInsightTags_Update
--- PASS: TestAccInsightTags_Update (19.71s)
PASS
ok      github.com/dikhan/terraform-provider-insight/insight      55.636s

```

Or specific tests can also be executed as follows:

```sh
$ TF_ACC=1 INSIGHT_API_KEY="<API_KEY>" INSIGHT_REGION="<REGION>" go test github.com/dikhan/terraform-provider-insight/insight -run  ^TestAccInsightTags_Create$ -timeout 120m -v
```

The acceptance tests require a INSIGHT_API_KEY and INSIGHT_REGION to be set. These env variables value will be used within the tests to
successfully interact with the insight api.

*Note: Acceptance tests create real resources and perform clean up tasks afterwards.*

Contributing
------------
Please follow the guidelines from:

 - [Contributor Guidelines](.github/CONTRIBUTING.md)

Authors
-------

Daniel I. Khan Ramiro

See also the list of [contributors](https://github.com/dikhan/terraform-provider-insight/graphs/contributors) who
participated in this project.
