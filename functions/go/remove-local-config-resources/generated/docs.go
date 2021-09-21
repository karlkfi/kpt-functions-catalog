

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var RemoveLocalConfigResourcesShort = `Removes resources with the annotation ` + "`" + `config.kubernetes.io/local-config: true` + "`" + ` from the resource list.`
var RemoveLocalConfigResourcesLong = `
## Usage

The function will execute as follows:

1. Searched for defined resources in a package
2. Deletes the resources with the following annotation:
   ` + "`" + `config.kubernetes.io/local-config: true` + "`" + `

` + "`" + `remove-local-config-resources` + "`" + ` function can be executed imperatively as follows:

  $ kpt fn eval -i gcr.io/kpt-fn/remove-local-config-resources:unstable

To execute ` + "`" + `remove-local-config-resources` + "`" + ` declaratively include the function in kpt package pipeline as follows:
  ...
  pipeline:
    mutators:
      - image: gcr.io/kpt-fn/remove-local-config-resources:unstable
  ...
`
