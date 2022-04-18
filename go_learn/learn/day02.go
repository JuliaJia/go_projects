package learn

import (
	"bufio"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//time库
func Basic() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Unix(), t.UnixMilli(), t.UnixMicro(), t.UnixNano())
	fmt.Println()
	begin := time.Now()
	useTime := time.Since(begin)
	fmt.Println(useTime.Seconds())
	fmt.Println(useTime.Milliseconds())
	fmt.Println(useTime.Microseconds())
	fmt.Println(useTime.Nanoseconds())
	fmt.Println()

	end := time.Now()
	useTime = end.Sub(begin)
	fmt.Println(useTime.Seconds())
	fmt.Println(useTime.Milliseconds())
	fmt.Println(useTime.Microseconds())
	fmt.Println(useTime.Nanoseconds())
	fmt.Println()

	dua := time.Duration(8 * time.Hour)
	end = begin.Add(dua)
	useTime = end.Sub(begin)
	fmt.Println(useTime.Seconds())
	fmt.Println(useTime.Milliseconds())
	fmt.Println(useTime.Microseconds())
	fmt.Println(useTime.Nanoseconds())

}

func Basic2() {
	begin := time.Now()
	fmt.Println(begin.Weekday().String())
	fmt.Println(begin.Year())
	fmt.Println(begin.Month())
	fmt.Println(begin.Day())
	fmt.Println(begin.Hour())
	fmt.Println(begin.Minute())
	fmt.Println(begin.Second())
	fmt.Println(begin.YearDay())
	fmt.Println()

	end := time.Now()
	fmt.Println(end.Weekday().String())
	fmt.Println(end.Year())
	fmt.Println(end.Month())
	fmt.Println(end.Day())
	fmt.Println(end.Hour())
	fmt.Println(end.Minute())
	fmt.Println(end.Second())
	fmt.Println(end.YearDay())
	fmt.Println()

	dua := time.Duration(8 * time.Hour)
	end = begin.Add(dua)
	fmt.Println(end.Weekday().String())
	fmt.Println(end.Year())
	fmt.Println(end.Month())
	fmt.Println(end.Day())
	fmt.Println(end.Hour())
	fmt.Println(end.Minute())
	fmt.Println(end.Second())
	fmt.Println(end.YearDay())
	fmt.Println()

}

func Basic3() {
	begin := time.Now()
	fmt.Println(begin.Format("2006-01-02 15:04:05"))
	//不建议这样使用
	if t, err := time.Parse("2006-01-02", "2022-07-20"); err == nil {
		day := t.YearDay() - begin.YearDay()
		fmt.Println(day)
	}

	//建议这样使用
	var loc *time.Location
	loc, _ = time.LoadLocation("Asia/Shanghai")
	if t, err := time.ParseInLocation("2006-01-02", "2022-07-20", loc); err == nil {
		day := t.YearDay() - begin.YearDay()
		fmt.Println(day)
	}

}

func Ticker() {
	tc := time.NewTicker(10 * time.Second)
	defer tc.Stop()
	for i := 0; i < 6; i++ {
		<-tc.C
		fmt.Println(time.Now().Unix())
	}
}

func Timer() {
	fmt.Println(time.Now().Unix())
	tm := time.NewTimer(10 * time.Second)
	defer tm.Stop()
	<-tm.C
	fmt.Println(time.Now().Unix())
}

//数学计算库

func Constant() {
	fmt.Println(math.E)
	fmt.Println(math.Pi)
	fmt.Println(math.MaxInt)
	fmt.Println(uint(math.MaxUint64))
	fmt.Printf("%b\n", uint64(math.MaxUint64))
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)
}

func Nan(str string) float64 {
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return f
	} else {
		return math.NaN()
	}
}

func Ceil(f float64) {
	fmt.Println(math.Ceil(f))
}

func Floor(f float64) {
	fmt.Println(math.Floor(f))
}

func Trunc(f float64) {
	fmt.Println(math.Trunc(f))
}

func Modf(f float64) {
	fmt.Println(math.Modf(f))
}

func Abs(f float64) {
	fmt.Println(math.Abs(f))
}

func Max(a, b float64) {
	fmt.Println(math.Max(a, b))
}

func Min(a, b float64) {
	fmt.Println(math.Min(a, b))
}

func Dim(a, b float64) {
	fmt.Println(math.Dim(a, b))
}

func Sqrt(f float64) {
	fmt.Println(math.Sqrt(f))
}

func Pow(f, n float64) {
	fmt.Println(math.Pow(f, n))
}

func Pow10(f int) {
	fmt.Println(math.Pow10(f))
}

func MathELog(f float64) {
	fmt.Println(math.Log(f))
}

func Math2Log(f float64) {
	fmt.Println(math.Log2(f))
}

func Math10Log(f float64) {
	fmt.Println(math.Log10(f))
}

func Exp(f float64) {
	fmt.Println(math.Exp(f)) //e的f次方
}

func Sin(f float64) {
	fmt.Println(math.Sin(f))
}

func Cos(f float64) {
	fmt.Println(math.Cos(f))
}

func Tanh(f float64) {
	fmt.Println(math.Tanh(f))
}

func RanFloat32() {
	fmt.Println(rand.Float32())
}

func RanFloat64() {
	fmt.Println(rand.Float64())
}

func RanShuffleInt(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	rect := make([]int, n)
	//rand.Shuffle(n, func(i, j int) {
	//	arr[i], arr[j] = arr[j], arr[i]
	//})
	for i := 0; i < 5*n; i++ {
		idx1 := rand.Intn(n)
		idx2 := rand.Intn(n)
		rect[idx1] = arr[idx2]
		rect[idx2] = arr[idx1]
	}
	return rect
}

func RanShuffleInt2(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	rect := arr
	rand.Shuffle(n, func(i, j int) {
		rect[i], rect[j] = rect[j], rect[i]
	})
	return rect
}

//io库
func Std(str string) {
	fmt.Println("Please input a string: ")
	fmt.Scan(&str)
	fmt.Printf("Your input %s\n", str)

}

func ReadFile(path string) {
	if file, err := os.Open(path); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		var buffer strings.Builder
		for {
			bs := make([]byte, 1024)
			if n, err := file.Read(bs); err != nil {
				if err == io.EOF {
					fmt.Println("文件已读完！")
					break
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Printf("从文件中读出%d个字节\n", n)
				buffer.WriteString(string(bs))
			}
		}
		fmt.Println(buffer.String())

	}
}

func ReadFile2(path string) {
	if file, err := os.Open(path); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		var i int
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					fmt.Println("文件已读完！")
					break
				}
			} else {
				fmt.Printf("第%d行\n", i)
				fmt.Println(strings.Trim(line, "\n"))
				i++
			}

		}
	}
}

func WriteFile(path, input string) {
	if file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		fmt.Println("请输入你需要写入的内容(输入quit退出)：")
		for {
			fmt.Scan(&input)
			if input != "quit" {
				input += "\n"
				file.Write([]byte(input))
			} else {
				break
			}

		}

	}
}

func WriteFile2(path, input string) {
	if file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)
		fmt.Println("请输入你需要写入的内容(输入quit退出)：")
		for {
			fmt.Scan(&input)
			if input != "quit" {
				input += "\n"
				writer.WriteString(input)
			} else {
				break
			}

		}
		writer.Flush()
	}
}

func FileOpRemoveFile(path string) {
	if err := os.Remove(path); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("删除文件%s成功\n", path)
	}
}

func FileOpMkdirAll(path string) {
	if err := os.MkdirAll(path, 0666); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("创建目录%s成功\n", path)
	}
}

func FileOpRemoveAll(path string) {
	if err := os.RemoveAll(path); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("删除文件%s成功\n", path)
	}
}

func FileOpRename(oldname, newname string) {
	if err := os.Rename(oldname, newname); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件改名成功")
	}
}

func FileOpInfo(path string) {
	file, _ := os.Open(path)
	info, _ := file.Stat()
	fmt.Println(info.IsDir())
	fmt.Println(info.ModTime())
	fmt.Println(info.Mode())
	fmt.Println(info.Size())
	fmt.Println(info.Name())
}

func FileOpInfo2(path string) error {
	info := ""
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			info = fileInfo.Name() + " " + strconv.FormatInt(fileInfo.Size(), 10) + " " + fileInfo.ModTime().Format("2006-01-02 15:04:05")
			fmt.Println(info)
			if fileInfo.IsDir() {
				if err := FileOpInfo2(path + "/" + fileInfo.Name()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func TempLogCreate(le, path string) {
	le += strconv.FormatInt(time.Now().UnixNano(), 10) + ".log"
	path += le
	if file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		logWriter := log.New(file, "[BIZ_NAME]", log.Ldate|log.Lmicroseconds)
		logWriter.Println("Hello")
		logWriter.Println("World")
	}
}

func SysCall(cmd string, args ...string) {
	if cmdPath, err := exec.LookPath(cmd); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cmdPath)
		cmd := exec.Command(cmd, args...)
		if output, err := cmd.Output(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(output))
		}
	}

}

func Compress(filepath, comfilepath string) {
	if file, err := os.Open(filepath); err != nil {
		fmt.Println(err)
	} else {
		if fout, err := os.OpenFile(comfilepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
			fmt.Println(err)
		} else {
			writer := zlib.NewWriter(fout)
			defer file.Close()
			defer fout.Close()
			for {
				bs := make([]byte, 10)
				if n, err := file.Read(bs); err != nil {
					if err == io.EOF {
						break
					} else {
						fmt.Println(err)
					}

				} else {
					writer.Write(bs[:n])
				}

			}
			writer.Flush()
		}
	}
}

func CompressReader(filepath string) {
	if file, err := os.Open(filepath); err != nil {
		fmt.Println(err)
	} else {

		if reader, err := zlib.NewReader(file); err != nil {
			fmt.Println(err)
		} else {
			io.Copy(os.Stdout, reader)
			defer file.Close()
			defer reader.Close()
		}
	}

}

//编码相关库
func JsonFunc() {
	user := User{"Ryomajia", "天才", 18, 1, 1}
	if bs, err := json.Marshal(user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}
}

func Base64Func(path string, size int) string {
	if fin, err := os.Open(path); err != nil {
		fmt.Println(err)
	} else {
		bs := make([]byte, 1024*size)
		if n, err := fin.Read(bs); err != nil {
			fmt.Println(err)
		} else {
			str := base64.StdEncoding.EncodeToString(bs[:n])
			fmt.Println(str)
			return str
		}

	}
	return ""
}

func Base64Decode(path, str string) {
	if bd, err := base64.StdEncoding.DecodeString(str); err != nil {
		fmt.Println(err)
	} else {
		if fout, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
			fmt.Println(err)
		} else {
			fout.Write(bd)
			fout.Close()
		}
	}

}