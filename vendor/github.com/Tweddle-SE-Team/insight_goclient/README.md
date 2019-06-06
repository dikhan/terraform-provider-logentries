# Insight Go Client

Full fledged Go client for [Insight](https://insight.rapid7.com). The library supports CRUD operations for the following
resources provided via the [Insight REST Api](https://insightops.help.rapid7.com/docs/rest-api-overview).

- [LogSets](https://insightops.help.rapid7.com/docs/logsets)
- [Logs](https://insightops.help.rapid7.com/docs/logs)
- [Tags](https://insightops.help.rapid7.com/docs/api-tags)
- [Labels](https://insightops.help.rapid7.com/docs/labels)

The above resources are available in the client via its seamless easy-to-use interface and in a matter of few lines you
can have a working client ready to be used with Insight.

# How to use the client?

Insight Go Client is really easy to use. The client exposes multiple resources available in Insight and
each of them offer create, read, update and delete (CRUD) operations.

Here is an example on how you can create a insight client and query all the logsets under the account tight
to the API key which the client was configured with:

```

import (
	"github.com/dikhan/insight_goclient"
)

func main() error {
	c := insight_goclient.NewInsightClient("INSIGHT_API_KEY", "eu")
	logsets, err := c.GetLogsets()
	if err != nil {
	    return err
	}
	fmt.println(logsets)
}
```

## Contributing

- Fork it!
- Create your feature branch: git checkout -b my-new-feature
- Commit your changes: git commit -am 'Add some feature'
- Push to the branch: git push origin my-new-feature
- Submit a pull request :D

## Authors

Daniel I. Khan Ramiro

See also the list of [contributors](https://github.com/dikhan/insight_goclient/graphs/contributors) who participated in this project.
