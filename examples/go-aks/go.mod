module goaks

go 1.14

require (
	github.com/pulumi/pulumi-azuread/sdk/v2 v2.4.0
	github.com/pulumi/pulumi-azurerm/sdk v0.1.0
	github.com/pulumi/pulumi-random/sdk/v2 v2.2.0
	github.com/pulumi/pulumi-tls/sdk/v2 v2.2.0
	github.com/pulumi/pulumi/sdk/v2 v2.7.2-0.20200803183728-996fd85e0095
)

replace github.com/pulumi/pulumi-azurerm/sdk => ../../sdk