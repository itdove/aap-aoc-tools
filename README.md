# aap-aoc-tools

## Installation

1. Clone the project
2. Run `go install` will install the `aap-aoc-tools`

## Execute

`aap-aoc-tools -h`

## Quick Start

### Dependencies command

1. Install the requiements as describe at [Requirements](#requirements)
2. Run 
```
aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --output-file mygraph.jpeg --options "-Tjpeg"
```
For more explanation see [Dependencies command](#dependencies-command-1)

## Commands 
### Dependencies command

```
    aap-aoc-tools dep ....
```
1. Requirements
- gcloud CLI
if you want to directly retreive the manifest from your project, you will need to install `gcloud` CLI and set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.

- Graphviz binary

if you want to generate a vizualization of the graph, you will need to install Graphviz https://graphviz.org/download/

2. Examples:
- Generate graph from manifest-file

    - Retreive manifest file by using command: 
        ```
        gcloud deployment-manager manifests describe <manifest-name> --deployment <deployment-name> > mymanifest.yaml
        ```
    - Run
        ```
        aap-aoc-tools dep --manifest-file mymanifest.yaml 
        ```
        This command will display the Graphviz format of the dependencies

- Generate graph from google cloud account

    - Set the GOOGLE_APPLICATION_CREDENTIALS environment variable with the path of your gcloud crendentials file.
    - Run
        ```
        aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name>
        ```

3. Options

- Suppress some resources

    Maybe the generated graph is too large and so you would like to remove some resources. You can do that with the `--exclude` option.
    ```
    aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --exclude <resource_1_name> --exclude <resource_2_name>
    ```
- Focus on a given resource

    If you want to know the dependencies of a single resource, you can specify that resource with the option `--start-resource`.
    ```
    aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --start-resouce <resource_1_name>
    ```
- Reverse the graph

    By default the edges of the graph are representing the "dependsOn" but maybe you want to know which resource serves which other resources which requires to reverse the direction of the edges, this can be done with the option `--reverse`
    ```
    aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --reverse
    ```
- Output file

    The command above are displaying the Graphviz representation of the graph but you can save it in a file using the `--output-file` option.
    ```
    aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --output-file <output_file>
    ```
- Graphviz engine

    All the above commands are generating Graphviz format graph representation but you can instead directly generate the graph representation (jpeg, svg,...) and this is done by using the options `--engine` and `--options`. 
    The command will call behind the scene the Graphviz executable.

    The default engine is the Graphviz `dot` option. 

    The default option is `-Ttxt` which is not a valid Graphviz option but will generate the Graphviz graph representation as above. If the option is not `-Ttxt` then the option `--output-file` is mandatory. The output format, `-T` parameter can be found here https://graphviz.org/docs/outputs/.

    For exemple to generate a `dot` jpeg
    ```
    aap-aoc-tools dep --manifest <manifest_name> --deployment <deployment_name> --output-file <output_file> --options "-Tjpeg"
    ```
    Note the double-quote surrounding the `--options` value.

    If you want to use another engine than the default `dot` you can specify it in the `--engine` options. The available engines are described at https://graphviz.org/docs/layouts/