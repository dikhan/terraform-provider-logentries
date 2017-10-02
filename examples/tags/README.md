# Logentries - Resource Tags example

This example shows how to manage tags in logentries via terraform.

This example would create a new tag associated with the sources logs specified in the sources param and the appropriate labels
configured will be attached to the tag. Make sure that the values introduced for the log source as well as the labels exist
in your logentries account.

Logentries resource expect the var.api_key to be passed in as env variable. Also. in this specific example as we are 
creating a PagerDuty alert for the tag, an env variable PD_KEY is passed to terraform commands.

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
TF_VAR_api_key="YOUR_API_KEY" TF_VAR_pagerduty_key="YOUR_PD_API_KEY" terraform plan
```

The logging level can be configured by specifying TF_LOG="DEBUG" and pass it into the terraform commands. 
For more information about debugging in terraform refer to this [link](https://www.terraform.io/docs/internals/debugging.html).

### Apply Phase

For apply phase execute:

```
TF_VAR_api_key="YOUR_API_KEY" TF_VAR_pagerduty_key="YOUR_PD_API_KEY"  terraform apply
```

Upon successful apply completion, go ahead and check that the tag actually exist in logentries:

```
curl https://logentries.com/app/<YOUR_ACCOUNT_NUMBER>#/tags/edit/<NEW_TAG_ID>
```

### Destroy Phase

To remove the newly created tag (this can be found inside the state file - terraform.tfstate), execute:

```
TF_VAR_api_key="YOUR_API_KEY" TF_VAR_pagerduty_key="YOUR_PD_API_KEY"  terraform destroy
```

