package learn

import (
	"encoding/json"
	"fmt"
	farm "github.com/dgryski/go-farm"
	"github.com/spaolacci/murmur3"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Discretizer interface {
	Discretize(i interface{}) string
}

type Uniq struct{}

func (Uniq) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.Itoa(int(v))
	case float64:
		return strconv.Itoa(int(v))
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

type Log struct {
	Base float64
}

func (self Log) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case int32:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case int64:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case float32:
		return strconv.Itoa(int(self._log(float64(v), self.Base)))
	case float64:
		return strconv.Itoa(int(self._log(v, self.Base)))
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

func (Log) _log(f float64, base float64) float64 {
	if f < 1 {
		f = 1
	}
	return math.Log(f) / math.Log(base)
}

type Bin struct {
	Splits []float64
}

func (self Bin) Discretize(i interface{}) string {
	switch v := i.(type) {
	case int:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case int32:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case int64:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case float32:
		return strconv.Itoa(self.BinDiscretize(float64(v), self.Splits))
	case float64:
		return strconv.Itoa(self.BinDiscretize(v, self.Splits))
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

func (Bin) BinDiscretize(f float64, splits []float64) int {
	index := 0
	for _, split := range splits {
		if f < split {
			break
		}
		index++
	}
	return index
}

type Hour struct{}

func (Hour) Discretize(i interface{}) string {
	switch v := i.(type) {
	case time.Time:
		return strconv.FormatInt(int64(v.Hour()), 10)
	case int64:
		return strconv.FormatInt(int64(time.Unix(v, 0).Hour()), 10)
	default:
		return ""
	}
}

type Transformer interface {
	Hash(string, int) uint64
}

type Murmur struct{}

func (Murmur) Hash(str string, featureId int) uint64 {
	ss := strconv.FormatInt(int64(featureId), 10) + ":" + str
	hash := murmur3.New64()
	hash.Write([]byte(ss))
	z := hash.Sum64()
	hash.Reset()
	return z
}

type FarmHash struct{}

func (FarmHash) Hash(str string, featureId int) uint64 {
	return farm.Hash64WithSeed([]byte(str), uint64(featureId))
}

type Location struct {
	Province string
	City     string
}

type User1 struct {
	Name    string
	Age     int
	Gender  byte
	Address *Location
}

type Product struct {
	Id          int
	Name        string
	Sales       int
	Feedback    float32
	Seller      *User1
	OnShelfTime time.Time
	Tags        []string
}

type FeatureConfig struct {
	Id             int         `json:"id"`
	Path           string      `json:"path"`
	Discretize     string      `json:"discretize"`
	Hash           string      `json:"hash"`
	DiscretizeFunc Discretizer `json:"-"`
	HashFunc       Transformer `json:"-"`
}

type FeatureConfigList []*FeatureConfig

var (
	blankReg *regexp.Regexp
)

func init() {
	var err error
	blankReg, err = regexp.Compile("\\s+")
	if err != nil {
		panic(err)
	}
}

type FeatureExtractor []*FeatureConfig

func (m *FeatureExtractor) Init(confFile string) error {
	if fin, err := os.Open(confFile); err != nil {
		return err
	} else {
		defer fin.Close()
		const MAX_CONFIG_SIZE = 1 << 20
		bs := make([]byte, MAX_CONFIG_SIZE)
		if n, err := fin.Read(bs); err != nil {
			return err
		} else {
			if n >= MAX_CONFIG_SIZE {
				return fmt.Errorf("config file size more than %dB", MAX_CONFIG_SIZE)
			}
			var confList FeatureConfigList
			if err = json.Unmarshal(bs[:n], &confList); err != nil {
				return err
			} else {
				productType := reflect.TypeOf(Product{})
				for _, conf := range confList {
					paths := splitString(conf.Path, '.')
					if field, ok := productType.FieldByName(paths[0]); !ok {
						return fmt.Errorf("field %s not fount,full path %s", paths[0], conf.Path)
					} else {
						if !field.IsExported() {
							return fmt.Errorf("field %s is not exported", paths[0])
						}
						for i := 1; i < len(paths); i++ {
							fieldType := field.Type
							if fieldType.Kind() == reflect.Ptr {
								fieldType = fieldType.Elem()
							}
							if field, ok = fieldType.FieldByName(paths[i]); !ok {
								return fmt.Errorf("field %s not found,full path %s", paths[i], conf.Path)
							} else {
								if !field.IsExported() {
									return fmt.Errorf("field %s is not exported", paths[i])
								}
							}
						}
					}
					conf.DiscretizeFunc = parseDiscretFunc(conf.Discretize)
					conf.HashFunc = parseHashFunc(conf.Hash)
					if conf.DiscretizeFunc == nil {
						return fmt.Errorf("id %d Discretize %s is INVALID", conf.Id, conf.Discretize)
					}
					if conf.HashFunc == nil {
						return fmt.Errorf("id %d Hash %s is INVALID", conf.Id, conf.Hash)
					}
					*m = append(*m, conf)
				}
			}
		}
	}
	return nil
}

func parseDiscretFunc(str string) Discretizer {
	parts := blankReg.Split(strings.ToLower(str), 2)
	switch parts[0] {
	case "uniq":
		return Uniq{}
	case "log":
		if len(parts) == 1 {
			return nil
		}
		base, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("log base is not float")
			return nil
		}
		return Log{Base: base}
	case "bin":
		if len(parts) == 1 {
			return nil
		}
		arr := splitString(parts[1], ',')
		splits := make([]float64, 0, len(arr))
		for _, ele := range arr {
			if f, err := strconv.ParseFloat(ele, 64); err != nil {
				fmt.Println("ele is not float")
				return nil
			} else {
				splits = append(splits, f)
			}
		}
		if len(splits) == 0 {
			return nil
		}
		return Bin{Splits: splits}
	case "hour":
		return Hour{}
	default:
		return nil
	}
}

func parseHashFunc(str string) Transformer {
	switch strings.ToLower(str) {
	case "murmur":
		return Murmur{}
	case "farm":
		return FarmHash{}
	default:
		return FarmHash{}
	}
}

func (m FeatureExtractor) extractOneField(conf *FeatureConfig, field reflect.Value, result *[]uint64) {
	switch field.Kind() {
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			m.extractOneField(conf, field.Index(i), result)
		}
	case reflect.Array, reflect.Map, reflect.Chan:
		return
	case reflect.Struct:
		fieldType := field.Type()
		if fieldType.String() != "time.Time" && fieldType.String() != "*time.Time" {
			return
		}
		fallthrough
	default:
		disc := conf.DiscretizeFunc.Discretize(field.Interface())
		h := conf.HashFunc.Hash(disc, conf.Id)
		*result = append(*result, h)
	}
}

func (m FeatureExtractor) Extract(product *Product) []uint64 {
	rect := make([]uint64, 0, 100)
	productValue := reflect.ValueOf(product).Elem()
	for _, conf := range m {
		paths := splitString(conf.Path, '.')
		field := productValue.FieldByName(paths[0])
		for i := 1; i < len(paths); i++ {
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
			field = field.FieldByName(paths[i])
		}
		result := make([]uint64, 0, 10)
		m.extractOneField(conf, field, &result)
		fmt.Println(conf.Path, result)
		rect = append(rect, result...)
	}
	return rect
}

func splitString(str string, delimiter byte) []string {
	rect := make([]string, 0, 10)
	begin := 0
	d := rune(delimiter)
	for i, r := range str {
		if r == d {
			rect = append(rect, str[begin:i])
			begin = i + 1
		}
	}
	if begin < len(str) {
		rect = append(rect, str[begin:])
	}
	return rect
}
