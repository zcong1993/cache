package expire

import (
	"github.com/zcong1993/cache/expire/bytesexpire"
	"github.com/zcong1993/cache/expire/iexpire"
	"github.com/zcong1993/cache/expire/stringexpire"
	"time"
)

//go:generate mkdir -p bytesexpire
//go:generate go_generics -i internal/expire/expire.go -t T=[]byte -o bytesexpire/bytesexpire.go -p bytesexpire
//go:generate go_generics -i internal/expire/expire_test.go -t P=[]byte -o bytesexpire/bytesexpire_test.go -p bytesexpire

// NewBytesExpireMap return a bytes expire map instance
func NewBytesExpireMap(gcInterval time.Duration) *bytesexpire.ExipreMap {
	return bytesexpire.NewExpireMap(gcInterval)
}

//go:generate mkdir -p stringexpire
//go:generate go_generics -i internal/expire/expire.go -t T=string -o stringexpire/stringexpire.go -p stringexpire
//go:generate go_generics -i internal/expire/expire_test.go -t P=string -o stringexpire/stringexpire_test.go -p stringexpire

// NewStringExpireMap return a string expire map instance
func NewStringExpireMap(gcInterval time.Duration) *stringexpire.ExipreMap {
	return stringexpire.NewExpireMap(gcInterval)
}

// NewExpireMap return a expire map with interface{} value type
func NewExpireMap(gcInterval time.Duration) *iexpire.ExipreMap {
	return iexpire.NewExpireMap(gcInterval)
}
