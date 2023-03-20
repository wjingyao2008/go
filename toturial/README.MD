# create module

```
  <home>/
   |-- greetings/
   |-- hello
```
### 1 First create greetings module
```
 cd hello
 go mod init example.com/greetings
```
### 2. Then create Hello module
```
cd hello
go mod init example.com/hello
```
### 3. Direct the greetings module location
tell "hello" module where the greetings module is
```
go mod edit -replace example.com/greetings=../greetings
```
then synchronize module:
```
go mod tidy
```
### 4. Import the module in code
```
import (
	"example.com/greetings"
	"fmt"
)
```