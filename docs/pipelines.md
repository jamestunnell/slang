# Pipelines

A pipeline is used to create and transform data. After the pipeline is started, each stage of the pipeline is an expression that can transform stage inputs into the stage outputs. Outputs are passed on to the next stage as inputs.

```
use "stdout"

func main() {
    var x {5,6,7}
    var y1 x.map(mul(7), add(10))   // y = (7 * x) + 10
    var y2 x.map(sqr, mul(0.5), add(0.2))   // y = (0.5 * x^2) + 0.2
    
    stdout.println(y1) // {45,52,59}
    stdout.println(y2) // {12.7,18.2,24.7}
}
