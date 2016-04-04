# pingdom-go-sdk

pingdom-go-sdk is a partial SDK for the [Pingdom][1] API.

This is a work in progress, mainly meant to support efforts to get 
Pingdom into Terraform. Many features may be missing, however the request
framework and data structuring is set up in a way to make it easily
exstensible, ensuring support for new features can be added in without fuss.

The Pingdom API reference can be found [here][2].

## Installing

If you are using Go 1.6 or higher, use the following to install the SDK.
Dependencies will be contained within the `vendor` directory, so no further
action is necessary.

```
go get -u github.com/paybyphone/pingdom-go-sdk
```

If you are using Go 1.5, you will need to set the `GO15VENDOREXPERIMENT=1`
vendoring flag to ensure vendored dependencies are used.

If you are using a version of Go below 1.5, you will need to tell go to get the
SDK and all of its dependencies.

```
go get -u github.com/paybyphone/pingdom-go-sdk/...
```

## Configuring Credentials

Credentials can be provided in one of two ways.

Note that in addition to your Pingdom account and password, an application key
is necessary. See the [authentication section][3] in the API for more info.

### Environment variables

Set the following variables:

```
export PINGDOM_EMAIL_ADDRESS=pingdom@example.com
export PINGDOM_PASSWORD=password
export PINGDOM_APP_KEY=pingdomappkey
```

If done this way, no config object needs to be passed to the `resource`
clients.

### Authentication via code

You can also configure the credentials through code. The below example
demonstrates how to get a [Checks client][4], passing in a custom config:

```
config := pingdom.Config{
  EmailAddress: "pingdom@example.com",
  Password:     "password",
  AppKey:       "pingdomappkey",
}

client := checks.New(config)
```

## Documentation

See [the GoDoc][5] for documentation.


## License

```
Copyright 2016 PayByPhone Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[1]: https://www.pingdom.com/
[2]: https://www.pingdom.com/resources/apia
[3]: https://www.pingdom.com/resources/api#authentication
[4]: https://godoc.org/github.com/paybyphone/pingdom-go-sdk/resource/checks
[5]: https://godoc.org/github.com/paybyphone/pingdom-go-sdk
