// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

var packages = []string{
	"axios@0.16.2",
	"babel-core@6.25.0",
	"babel-loader@7.1.1",
	"babel-preset-es2015@6.24.1",
	"babel-preset-react@6.24.1",
	"css-loader@0.28.4",
	"extract-text-webpack-plugin@3.0.0",
	"geolib@2.0.22",
	"import-glob-loader@1.1.0",
	"lodash@4.17.4",
	"node-sass@4.5.3",
	"prop-types@15.5.10",
	"query-string@5.0.0",
	"react@15.6.1",
	"react-dom@15.6.1",
	"react-google-maps@7.2.0",
	"react-redux@5.0.6",
	"react-router-dom@4.1.2",
	"redux@3.7.2",
	"redux-devtools@3.4.0",
	"redux-devtools-extension@2.13.2",
	"redux-thunk@2.2.0",
	"sass-loader@6.0.6",
	"style-loader@0.18.2",
	"webpack@3.3.0",
	"webpack-dev-server@2.5.1",
	"eslint@3.19.0",
	"eslint-config-airbnb@15.0.2",
	"eslint-plugin-import@2.7.0",
	"eslint-plugin-jsx-a11y@5.1.1",
	"eslint-plugin-react@7.1.0",
}

func Deps() {
	installPackages()
}

func DevServer() {
	sh.RunV("node_modules/.bin/webpack-dev-server")
}

func Webpack() {
	sh.RunV("yarn", "run", "start")
}

func installPackages() {
	for _, pkg := range packages {
		sh.RunV("yarn", "add", pkg)
	}
}

// ./node_modules/.bin/eslint --init
// yarn install
