package main

import (
	"fmt"
	"syscall/js"

	"github.com/ghodss/yaml"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/renderutil"
	"k8s.io/helm/pkg/timeconv"
)

func render(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")
	metadata := `apiVersion: v1
description: A Helm chart for Kubernetes
name: Helm Parser
version: 0.1.0`
	templates := document.Call("getElementById", "templates").Get("value").String()
	values := document.Call("getElementById", "values").Get("value").String()

	m := &chart.Metadata{}
	err := yaml.Unmarshal([]byte(metadata), m)
	if err != nil {
		fmt.Printf("cannot unmarshal chart metadata: %v", err)
	}

	t := []*chart.Template{}
	t = append(t, &chart.Template{Name: "service.yaml", Data: []byte(templates)})
	c := &chart.Chart{
		Metadata:  m,
		Templates: t,
	}
	config := &chart.Config{Raw: values, Values: map[string]*chart.Value{}}

	renderOpts := renderutil.Options{
		ReleaseOptions: chartutil.ReleaseOptions{
			Name:      "wasm",
			IsInstall: false,
			IsUpgrade: false,
			Time:      timeconv.Now(),
			Namespace: "wasm",
		},
		KubeVersion: "1.20",
	}

	renderedTemplates, err := renderutil.Render(c, config, renderOpts)

	var renderError string
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		renderError = err.Error()
	}
	fmt.Println(renderedTemplates)
	var result string
	for _, v := range renderedTemplates {
		result += v
	}

	if renderError != "" {
		js.Global().Get("document").Call("getElementById", "result").Set("textContent", renderError)
	} else {
		js.Global().Get("document").Call("getElementById", "result").Set("textContent", result)
	}
	return nil
}

func registerCallbacks() {
	js.Global().Set("render", js.FuncOf(render))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
