作业描述：[CLI 命令行实用程序开发](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic)
作业中提供了两种语言，一种语言是c的，一种语言是python。这两种语言我都略懂一点，但却深感自己学的不深，这里尝试看C语言版。
参考博客：
[刘一 思否博客](https://segmentfault.com/a/1190000016648238)
[CSDN有关博客](https://blog.csdn.net/C486C/article/details/82990187)
#### 前期准备
##### CLI
CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。Linux提供了cat、ls、copy等命令与操作系统交互；go语言提供一组实用程序完成从编码、编译、库管理、产品发布全过程支持；容器服务如docker、k8s提供了大量实用程序支撑云服务的开发、部署、监控、访问等管理任务；git、npm等都是大家比较熟悉的工具。尽管操作系统与应用系统服务可视化、图形化，但在开发领域，CLI在编程、调试、运维、管理中提供了图形化程序不可替代的灵活性与效率。

阅读：[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)


系统:centos7
#### 开发实践
使用 golang 开发 [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 中的 selpg

提示：

1、请按文档 使用 selpg 章节要求测试你的程序
2、请使用 pflag 替代 goflag 以满足 Unix 命令行规范， 参考：[Golang之使用Flag和Pflag](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)
3、golang 文件读写、读环境变量，请自己查 os 包
4、“-dXXX” 实现，请自己查 os/exec 库，例如案例 [Command](https://godoc.org/os/exec#example-Command)，管理子进程的标准输入和输出通常使用 io.Pipe，具体案例见 [Pipe](https://godoc.org/io#Pipe)


具体实现：
1、根据文档，此程序需要满足以下几个要求：
命令输入
参数处理
参数规范检测
读取文件
处理文件
输出文件

其中过程中若存在一些情况可以抛出错误。

###### 代码实现：
1、最开始需要做的是设置参数结构体：

```

type selpg_args struct {
    startPage  int
    endPage    int
    inFileName string
    pageLen    int
    pageType   bool
    printDest  string
}
```
输入参数使用包中提供的pflag进行处理。

```
fun getselpg(args *selpg_args){
    pflag.IntVarP(&(args.startPage), "startPage", "s", -1, "Define startPage")
    pflag.IntVarP(&(args.endPage), "endPage", "e", -1, "Define endPage")
    pflag.IntVarP(&(args.pageLen), "pageLength", "l", 72, "Define pageLength")
    pflag.StringVarP(&(args.printDest), "printDest", "d", "", "Define printDest")
    pflag.BoolVarP(&(args.pageType), "pageType", "f", false, "Define pageType")
    pflag.Parse()
	argLeft := pflag.args()
	if len(argLeft)>0{
		args.inFileName = string(argLeft[0])
	} else {
		args.inFileName = ""
	}
}
```
2、命令行参数获取之后，首先要进行参数检查以尽量避免参数谬误。出现错误时输出问题并正常结束程序。参数正确则将各个参数值输出到屏幕上。

```
func examselpg(args *selpg_args) {

	if (args.startPage == -1) || (args.endPage == -1) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be empty! Please check your command!\n")
		os.Exit(2)
	} else if (args.startPage <= 0) || (args.endPage <= 0) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be negative! Please check your command!\n")
		os.Exit(3)
	} else if args.startPage > args.endPage {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can't be bigger than the endPage! Please check your command!\n")
		os.Exit(4)
	} else if (args.pageType == true) && (args.pageLen != 72) {
		fmt.Fprintf(os.Stderr, "\n[Error]The command -l and -f are exclusive, you can't use them together!\n")
		os.Exit(5)
	} else if args.pageLen <= 0 {
		fmt.Fprintf(os.Stderr, "\n[Error]The pageLen can't be less than 1 ! Please check your command!\n")
		os.Exit(6)
	} else {
		pageType := "page length."
		if args.pageType == true {
			pageType = "The end sign /f."
		}
		fmt.Printf("\n[ArgsStart]\n")
		fmt.Printf("startPage: %d\nendPage: %d\ninputFile: %s\npageLength: %d\npageType: %s\nprintDestation: %s\n[ArgsEnd]", args.startPage, args.endPage, args.inFileName, args.pageLen, pageType, args.printDest)
	}

}

func ErrorCheck(err error, object string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n[Error]%s:", object)
		panic(err)
	}
}
```

3、参数检查结束之后，程序开始调用excutecmd函数执行命令。

```
func excutecmd(args *selpg_args) {
	var fin *os.File
	if args.inFileName == "" {
		fin = os.Stdin
	} else {
		checkFileAccess(args.inFileName)
		var err error
		fin, err = os.Open(args.inFileName)
		ErrorCheck(err, "File input")
	}

	if len(args.printDest) == 0 {
		output2Des(os.Stdout, fin, args.startPage, args.endPage, args.pageLen, args.pageType)
	} else {
		output2Des(cmdExec(args.printDest), fin, args.startPage, args.endPage, args.pageLen, args.pageType)
	}
}
```
4、在-d参数存在时，涉及到了os/exec包的使用。

```
func cmdExec(printDest string) (*exec.Cmd, io.WriteCloser) {
	cmd := exec.Command("lp", "-d"+printDest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fout, err := cmd.StdinPipe()
	ErrorCheck(err, "Input pipe open\n")
	return cmd, fout
}
```
5、输出函数output2Des将输入的文件，按页码要求读取并输出到fout中。

```
func output2Des(printDest string, fin *os.File, pageStart int, pageEnd int, pageLen int, pageType bool) {

	lineCount := 0
	pageCount := 1
	buf := bufio.NewReader(fin)

	var cmd *exec.Cmd
	var fout io.WriteCloser
	if len(printDest) > 0 {
		cmd, fout = cmdExec(printDest)
	}

	for true {

		var line string
		var err error
		if pageType {
			//If the command argument is -f
			line, err = buf.ReadString('\f')
			pageCount++
		} else {
			//If the command argument is -lnumber
			line, err = buf.ReadString('\n')
			lineCount++
			if lineCount > pageLen {
				pageCount++
				lineCount = 1
			}
		}

		if err == io.EOF {
			break
		}
		ErrorCheck(err, "file read in\n")

		if (pageCount >= pageStart) && (pageCount <= pageEnd) {
			var outputErr error
			if len(printDest) == 0 {
				_, outputErr = fmt.Fprintf(os.Stdout, "%s", line)
			} else {
				_, outputErr = fout.Write([]byte(line))
				ErrorCheck(outputErr, "pipe input")
			}
			ErrorCheck(outputErr, "Error happend when output the pages.")
		}
	}

	if len(printDest) > 0 {
		fout.Close()
		errStart := cmd.Run()
		ErrorCheck(errStart, "CMD Run")
	}

	if pageCount < pageStart {
		fmt.Fprintf(os.Stderr, "\n[Error]: startPage (%d) greater than total pages (%d), no output written\n", pageStart, pageCount)
		os.Exit(9)
	} else if pageCount < pageEnd {
		fmt.Fprintf(os.Stderr, "\n[Error]: endPage (%d) greater than total pages (%d), less output than expected\n", pageEnd, pageCount)
		os.Exit(10)
	}
}

```

