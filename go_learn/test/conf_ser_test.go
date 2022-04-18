package test

import (
	"encoding/json"
	"fmt"
	"go_learn/learn"
	"log"
	"os"
	"testing"
	"time"
)

var (
	loc *time.Location
	fe  *learn.FeatureExtractor
)

const (
	TIME_FMT = "2006-01-02 15:04:05"
	DATE_FMT = "20060102"
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
	fe = new(learn.FeatureExtractor)
	if err := fe.Init("/Users/ryomajia/Code/go_projects/go_learn/learn/feature_extractor.json"); err != nil {
		log.Fatal(err)
	}
}

func TestSerializeConf(t *testing.T) {
	conf1 := learn.FeatureConfig{
		Id:         1,
		Path:       "Seller.Name",
		Discretize: "uniq",
		Hash:       "murmur",
	}
	conf2 := learn.FeatureConfig{
		Id:         2,
		Path:       "Seller.Age",
		Discretize: "bin 18,25,30,35,40,50,60",
		Hash:       "farm",
	}
	confList := learn.FeatureConfigList{&conf1, &conf2}
	fout, err := os.OpenFile("/Users/ryomajia/Code/go_projects/go_learn/learn/test.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		defer fout.Close()
		bs, err := json.MarshalIndent(confList, "", " ")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fout.Write(bs)
		}
	}
}

func TestFeatureExtractor(t *testing.T) {
	tm, err := time.ParseInLocation(TIME_FMT, "2015-03-27 10:00:00", loc)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	product := &learn.Product{
		Id:       5346456,
		Name:     "iphone",
		Sales:    74575372,
		Feedback: 0.94,
		Seller: &learn.User1{
			Name:   "Apple",
			Age:    34,
			Gender: 1,
			Address: &learn.Location{
				Province: "北京",
				City:     "北京",
			},
		},
		OnShelfTime: tm,
		Tags:        []string{"电子产品", "手机"},
	}
	features := fe.Extract(product)
	fmt.Println("final result", features)
}

func BenchmarkGeatureExtractor(b *testing.B) {
	tm, err := time.ParseInLocation(TIME_FMT, "2015-03-27 10:00:00", loc)
	if err != nil {
		fmt.Println(err)
		b.Fail()
	}
	product := &learn.Product{
		Id:       5346456,
		Name:     "iphone",
		Sales:    74575372,
		Feedback: 0.94,
		Seller: &learn.User1{
			Name:   "Apple",
			Age:    34,
			Gender: 1,
			Address: &learn.Location{
				Province: "北京",
				City:     "北京",
			},
		},
		OnShelfTime: tm,
		Tags:        []string{"电子产品", "手机"},
	}
	for i := 0; i < b.N; i++ {
		fe.Extract(product)
	}
}

//go test -v learn/conf_ser_test.go -run=函数名字
