# Error Values

`error` is a built-in type used to convey failure. It will either have nil value, if unassigned, or it will contain an error category and message.

```
use "errors"

func verifyPositive(x int) (err error) {
    if x <= 0 {
        err = errors.constraintViolation("{{x}} is not positive")
    }
}
```