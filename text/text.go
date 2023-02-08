package gText

import (
	"fmt"
	"log"
	"strconv"
)

func Int642Str(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Str2Int64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}
func Int2Str(i int) string {
	return strconv.Itoa(i)
}
func Str2Int(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}
func Float642Str(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)

}
func Str2Float64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return i, err
}
func Float642Float64(f float64, w string) float64 {
	f, _ = strconv.ParseFloat(fmt.Sprintf("%."+w+"f", f), 64)
	return f
}
func Test() {
	s, _ := Str2Int64("320")
	log.Printf(`Str2Int64("320") %s`, s)
	log.Printf(`Int642Str(320) %s`, Int642Str(320))
	f, _ := Str2Float64("320.08")
	log.Printf(`Str2Float64("320.08") %s`, f)
	log.Printf(`Float642Str(320.08) %s`, Float642Str(f))
	i, _ := Str2Int("320")
	log.Printf(`Str2Int("320") %s`, i)
	log.Printf(`Int2Str(320) %s`, Int2Str(i))
	log.Printf(`Float642Float64(1725577.59 - 1381341.21 ,"2") %s`, Float642Float64(1725577.59-1381341.21, "2"))
}
