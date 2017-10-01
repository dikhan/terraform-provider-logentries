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

Refer to the [examples](https://github.com/dikhan/terraform-provider-logentries/) folder to see how to configure each 
resource provided by this terraform provider. Once the tf file is populated terraform can be executed as follows:

```
go build && terraform init && TF_VAR_api_key="API_KEY" terraform plan
```

The logging level can be configured by specifying TF_LOG="DEBUG" as part of the above command. For more information about
debugging in terraform refer to this [link](https://www.terraform.io/docs/internals/debugging.html)

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
$ TF_ACC=1 LOGENTRIES_API_KEY="API_KEY" SOURCE_ID="" LABEL_ID="" go test $(go list ./...) -timeout 120m
```

Or specific test can also be executed as follows:

```sh
$ TF_ACC=1 LOGENTRIES_API_KEY="API_KEY" SOURCE_ID="LOG_ID" LABEL_ID="LABEL_ID" go test ^TestAccLogentriesTags_Create$ -timeout 120m
```

The acceptance tests require various LOGENTRIES_API_KEY, SOURCE_ID and LABEL_ID env variables to be set. The env variables
values will be used within the tests to successfully interact with the log entries api. This values are specific to the
user and therefore should be passed in.

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