# Pipelines

A pipeline is used to create and transform data. After the pipeline is started, each stage of the pipeline is an expression that can transform stage inputs into the stage outputs. Outputs are passed on to the next stage as inputs.

A pipeline stage is specified using a function or expression.

When using a funciton, the signature must have a single input which is the pipeline value type at that stage.
When using an expression, the $ value represents a pipeline value at that stage.

```
use "stdout"

func main() {
    var x [5,6,7]
    
    // using expressions
    var y1 x.map | mul($ 7) | add($ 10)

    // using functions
    var y2 x.map | mul.bind(7) | add.bind(10)

    stdout.println(y1) // [45,52,59]
    stdout.println(y2) // [45,52,59]
}
