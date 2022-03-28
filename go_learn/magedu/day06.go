package magedu

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type task3 struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	Status    int
	User      string
}

func FilePath(filename string) string {
	fp, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return fp
}

func FileName(fp string) string {
	fn := filepath.Base(fp)

	return fn
}

func FileExt(fp string) string {
	fe := filepath.Ext(fp)

	return fe
}

func PathClean(fp string) string {
	pc := filepath.Clean(fp)
	return pc
}

func Builder1(str string) string {
	var builder strings.Builder
	builder.WriteString(str)
	new_str := builder.String()
	builder.Reset()
	return new_str

}

func Builder2(str string) string {
	reader := strings.NewReader(str)
	ctx := make([]byte, len(str))
	_, err := reader.Read(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return string(ctx)
}
func Builder3(str string) {
	reader := strings.NewReader(str)
	reader.WriteTo(os.Stdout)
}

func Buffer1(str string, str2 string) {
	bb := bytes.NewBuffer([]byte(str))
	bb.Write([]byte(str2))
	strl := len(str) + len(str2)
	ctx := make([]byte, strl)
	_, err := bb.Read(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(ctx))

}

func Buffer2(str string, str2 string) {
	bs := bytes.NewBufferString(str)
	bs.WriteString(str2)
	strl := len(str) + len(str2)
	ctx := make([]byte, strl)
	_, err := bs.Read(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(ctx))

}

func Buffer3(str string, str2 string) {
	bb := bytes.NewBuffer([]byte(str))
	bb.Write([]byte(str2))
	bb.WriteTo(os.Stdout)

}

func Buffer4(str string, str2 string) {
	bs := bytes.NewBufferString(str)
	bs.WriteString(str2)
	bs.WriteTo(os.Stdout)

}

func CopyFile1(src, dst string, size int) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	buffer := make([]byte, size)
	for {
		n, err := srcFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		n, err = dstFile.Write(buffer[:n])
		if err != nil {
			return err
		}

	}
	return nil
}

func CopyFile2(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

func CopyFile3(src, dst string, size []byte) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.CopyBuffer(dstFile, srcFile, size)
	return err
}

func CopyFile4(src, dst string, size int64) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.CopyN(dstFile, srcFile, size)
	return err
}

func MultiWriter1(log1, log2, str string) {
	logfile1, _ := os.Create(log1)
	logfile2, _ := os.Create(log2)
	writer := io.MultiWriter(logfile1, logfile2, os.Stdout)
	writer.Write([]byte(str))
	defer logfile1.Close()
	defer logfile2.Close()
}
func MultiReader1(log1, log2 string, buffer []byte) {
	logfile1, _ := os.Open(log1)
	logfile2, _ := os.Open(log2)
	reader := io.MultiReader(logfile1, logfile2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	defer logfile1.Close()
	defer logfile2.Close()
}

func MultiReader2(log1, log2 string) {
	logfile1, _ := os.Open(log1)
	logfile2, _ := os.Open(log2)
	buffer := bytes.NewBuffer([]byte(""))
	io.Copy(buffer, logfile1)
	io.Copy(buffer, logfile2)

	fmt.Println(buffer.String())

	defer logfile1.Close()
	defer logfile2.Close()
}

func Ioutil1(filePath string, buffer []byte) {
	ioutil.WriteFile(filePath, buffer, os.ModePerm)
}

func FormatSize(num int64) string {
	if num < 1024 {
		return strconv.Itoa(int(num)) + "B"
	} else if 1024 <= num && num < 1024*1024 {
		str := strconv.FormatFloat(float64(num)/float64(1024), 'f', 2, 64) + "KB"
		return str
	} else if 1024*1024 <= num && num < 1024*1024*1024 {
		str := strconv.FormatFloat(float64(num)/float64(1024*1024), 'f', 2, 64) + "MB"
		return str
	} else if 1024*1024*1024 <= num && num < 1024*1024*1024*1024 {
		str := strconv.FormatFloat(float64(num)/float64(1024*1024*1024), 'f', 2, 64) + "GB"
		return str
	} else if 1024*1024*1024*1024 <= num && num < 1024*1024*1024*1024*1024 {
		str := strconv.FormatFloat(float64(num)/float64(1024*1024*1024*1024), 'f', 2, 64) + "TB"
		return str
	} else if 1024*1024*1024*1024*1024 <= num && num < 1024*1024*1024*1024*1024*1024 {
		str := strconv.FormatFloat(float64(num)/float64(1024*1024*1024*1024*1024), 'f', 2, 64) + "PB"
		return str
	}
	return ""
}

func FormatSize2(num int64) string {
	units := map[int]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
	}

	index := 0
	fsize := float64(num)
	unit := float64(1024)
	for fsize >= unit {
		fsize /= unit
		index++
	}

	str := strconv.FormatFloat(fsize, 'f', 2, 64) + units[index]
	return str
}

func Ioutil2(filePath string) string {
	files, err := ioutil.ReadDir(filePath)
	str := ""
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			str += ("D " + file.Name() + " " + file.ModTime().Format("2006-01-02 15:03:04") + "\n")
		} else {
			str += ("F " + file.Name() + " " + FormatSize2(file.Size()) + " " + file.ModTime().Format("2006-01-02 15:03:04") + "\n")
		}

	}
	return str
}

func Log1(filename, input string) {
	dir, _ := ioutil.TempDir("./log", "log")
	file, _ := os.Create(filepath.Join(dir, filename))
	defer file.Close()
	file.WriteString(input)
}

func FileList(path string) {
	filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
		fmt.Println(path, file.Name(), file.IsDir())
		return nil
	})
}

func OsFile() {
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.Executable())
	path, _ := os.Executable()
	fmt.Println(filepath.Dir(path))
}

func BufioScanner(log1 string) {
	logfile1, _ := os.Open(log1)

	scanner := bufio.NewScanner(logfile1)
	//fmt.Println("请输入内容：")
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	defer logfile1.Close()
}
func BufioScanner2() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入内容：")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if "quit" == scanner.Text() {
			break
		}
	}

}

func BufioWriter1(filepath, input string) {
	file, _ := os.Create(filepath)
	writer := bufio.NewWriter(file)
	writer.WriteString(input)
	defer writer.Flush()
	defer file.Close()
}

func gobinit() {
	gob.Register(&task3{})
}

func GobEncode(filepath string) {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	tasks := []*task3{
		{
			Id:        1,
			Name:      "test",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "天才哥哥",
		},
		{
			Id:        2,
			Name:      "test2",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "小猪头",
		},
	}
	gobinit()
	file, _ := os.Create(filepath)
	in := gob.NewEncoder(file)

	in.Encode(tasks)
	defer file.Close()

}

func GobDecode(input string) {
	var tasks []*task3
	inf, _ := os.Open(input)
	out := gob.NewDecoder(inf)
	out.Decode(&tasks)
	for _, task := range tasks {
		fmt.Printf("%#v\n", task)
	}
	defer inf.Close()
}

func JsonEncode() []byte {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	tasks := []*task3{
		{
			Id:        1,
			Name:      "test",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "天才哥哥",
		},
		{
			Id:        2,
			Name:      "test2",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "小猪头",
		},
	}
	ctx, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}
	return ctx
}

func JsonDecode(input []byte, filepath string) {
	var buffer bytes.Buffer
	file, _ := os.Create(filepath)
	json.Indent(&buffer, input, "", "\t")
	buffer.WriteTo(file)
	defer file.Close()
}

func JsonCheck(filepath string) {
	jsontxt, _ := ioutil.ReadFile(filepath)
	fmt.Println(json.Valid(jsontxt))
}

func JsonEncodeFile(filepath string) {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	tasks := []*task3{
		{
			Id:        1,
			Name:      "test",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "天才哥哥",
		},
		{
			Id:        2,
			Name:      "test2",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "小猪头",
		},
	}
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(tasks)
}
