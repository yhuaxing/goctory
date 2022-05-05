# goctory
A struct factory implementation based on go generics

## Installation

```
$ go get -u github.com/yhuaxing/goctory
```
## Dependencies

go version >= 1.18

## Usage

It is very simple to use

```golang
type UserData struct {
    Id int
    Name string
    Age int
}

func main() {
    user := goctory.NewStruct[UserData](
		goctory.WithAttr[UserData, int]("Id", 1),
		goctory.WithAttr[UserData, string]("Name", "arthur"),
		goctory.WithAttr[UserData, int]("Age", 22),
	)
}
```
