// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/console/clientset/versioned/typed/console/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConsoleV1 struct {
	*testing.Fake
}

func (c *FakeConsoleV1) ConsoleCLIDownloads() v1.ConsoleCLIDownloadInterface {
	return &FakeConsoleCLIDownloads{c}
}

func (c *FakeConsoleV1) ConsoleExternalLogLinks() v1.ConsoleExternalLogLinkInterface {
	return &FakeConsoleExternalLogLinks{c}
}

func (c *FakeConsoleV1) ConsoleLinks() v1.ConsoleLinkInterface {
	return &FakeConsoleLinks{c}
}

func (c *FakeConsoleV1) ConsoleNotifications() v1.ConsoleNotificationInterface {
	return &FakeConsoleNotifications{c}
}

func (c *FakeConsoleV1) ConsolePlugins() v1.ConsolePluginInterface {
	return &FakeConsolePlugins{c}
}

func (c *FakeConsoleV1) ConsoleQuickStarts() v1.ConsoleQuickStartInterface {
	return &FakeConsoleQuickStarts{c}
}

func (c *FakeConsoleV1) ConsoleYAMLSamples() v1.ConsoleYAMLSampleInterface {
	return &FakeConsoleYAMLSamples{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConsoleV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
