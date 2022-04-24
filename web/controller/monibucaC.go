package controller

//	ffmpeg推流到指定 ip
import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"syscall"
	"text/template"
)

func Ffmpeg(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm();	fmt.Println(r.PostForm)
	//
	//video_name := r.PostFormValue("video_name")
	//puth_ip := r.PostFormValue("puth_ip")
	//stream_name := r.PostFormValue("stream_name")
	//if() {
	//
	//	//video_name := "Integrated Camera"
	//	//puth_ip := "172.21.73.25"
	//	//stream_name := "yy"
	//	str := string("ffmpeg -f dshow -i video=\"" + video_name + "\"" +
	//		" -vcodec libx264  -preset:v  ultrafast  -tune:v zerolatency  " +
	//		"-f flv  rtmp://" + puth_ip + "/live/" + stream_name)
	//
	//	fmt.Println("str:", str)
	//
	//	in := bytes.Buffer{}
	//	outInfo := bytes.Buffer{}
	//	cmd := exec.Command("cmd")
	//
	//	cmd.Stdout = &outInfo
	//	cmd.Stdin = &in
	//
	//	in.WriteString("chcp 65001\n")
	//	in.WriteString(str + "\n")
	//	in.WriteString("chcp\n")
	//
	//	err := cmd.Start()
	//	if err != nil {
	//		fmt.Println("Ffmpeg推流"+video_name+"到"+puth_ip+"失败：", err.Error())
	//	}
	//	if err = cmd.Wait(); err != nil {
	//		fmt.Println(err.Error())
	//	} else {
	//		fmt.Println(cmd.ProcessState.Pid())
	//		//程序退出code
	//		fmt.Println(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
	//		//输出结果
	//		fmt.Println(outInfo.String())
	//	}
	//}
	t := template.Must(template.ParseFiles("web/views/pages/monibuca/ffmpeg.html"))
	t.Execute(w, "")
}

func FfmpegPuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.PostForm)
	video_name := r.PostFormValue("video_name")
	puth_ip := r.PostFormValue("puth_ip")
	stream_name := r.PostFormValue("stream_name")
	fmt.Println("FfmpegPuth:", video_name+"|", puth_ip, "|", stream_name)

	//video_name := "Integrated Camera"
	//puth_ip := "172.21.73.25"
	//stream_name := "yy"
	str := string("ffmpeg -f dshow -i video=\"" + video_name + "\"" +
		" -vcodec libx264  -preset:v  ultrafast  -tune:v zerolatency  " +
		"-f flv  rtmp://" + puth_ip + "/live/" + stream_name)

	fmt.Println("str:", str)

	in := bytes.Buffer{}
	outInfo := bytes.Buffer{}
	cmd := exec.Command("cmd")

	cmd.Stdout = &outInfo
	cmd.Stdin = &in

	in.WriteString("chcp 65001\n")
	in.WriteString(str + "\n")
	in.WriteString("chcp\n")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Ffmpeg推流"+video_name+"到"+puth_ip+"失败：", err.Error())
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(cmd.ProcessState.Pid())
		//程序退出code
		fmt.Println(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitCode)
		//输出结果
		fmt.Println(outInfo.String())
	}
}
