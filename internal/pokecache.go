package main
import "fmt"
import "bufio"
import "io"
import "os"
import "strings"
import "net/http"
import "encoding/json"
import "time"
type Cache struct {
	
}
type cacheEntry struct {
	created time.Time
	val []byte
}
func NewCache(interval time.Duration) {

}