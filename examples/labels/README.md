# Logentries - Resource Label example

This example shows how to manage labels in logentries via terraform.

This example would create a label.

Logentries resource expect the var.api_key to be passed in as env variable.

## Running the example

### Plan Phase

Fist and foremost, build and init terraform

```
$ cd $GOPATH/src/github.com/dikhan/terraform-provider-logentries
$ go install && terraform init
```

This will install the binary inside $GOPATH/bin so terraform is aware about the logentries plugin.

For planning phase execute:

```
TF_VAR_api_key="YOUR_API_KEY" terraform plan
```

The logging level can be configured by specifying TF_LOG="DEBUG" and pass it into the terraform commands.
For more information about debugging in terraform refer to this [link](https://www.terraform.io/docs/internals/debugging.html).

### Apply Phase

For apply phase execute:

```
TF_VAR_api_key="YOUR_API_KEY" terraform apply
```

Upon successful apply completion, go ahead and check that the log actually exist in logentries:

```
curl https://rest.logentries.com/management/logs/<LOG_ID> -H "x-api-key: <YOUR_API_KEY>" -vv
```

### Destroy Phase

To remove the newly created log (this can be found inside the state file - terraform.tfstate), execute:

```
TF_VAR_api_key="YOUR_API_KEY" terraform destroy
```

