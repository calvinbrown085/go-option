# go-option
Optional type for Golang 



## Usage
```go
import (
	"context"

	gooption "github.com/calvinbrown085/go-option"
)

type Item struct {
	Id    string
	Value string
}

func GetItem(ctx context.Context, id string) (gooption.Option[Item], error) {
	return gooption.Some(Item{
		Id:    id,
		Value: "1234",
	}), nil
}

func main() {
	id := "1"

	maybeItem, err := GetItem(context.Background(), id)
	if err != nil {
		panic("")
	}

	_ = maybeItem.GetOrElse(Item{
		Id:    "2",
		Value: "4321",
	})
}
```