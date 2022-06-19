package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const SWD = "../OOP_course_work_modules/transformer/"

func useModule(id string, data string) {
	log.Println("executing test with", data)
	cmd := exec.Command("./venv/Scripts/python.exe", "script.py", data)
	cmd.Env = os.Environ()
	cmd.Dir = SWD + id
	cmd.Start()
}

func newModule(mtype string, settings string) string {
	if mtype == "script" {
		return newScript(settings)
	}
	return "0"
}

func newScript(settings string) string {
	var set map[string]interface{}
	json.Unmarshal([]byte(settings), &set)
	content := set["script"].(string)
	id := RandStringRunes(16)
	os.Mkdir(SWD+id, 777)
	cmdvenv := exec.Command(SWD+"script/venv/Scripts/python.exe", "-m", "venv", "./venv")
	cmdvenv.Env = os.Environ()
	cmdvenv.Dir = SWD + id
	cmdvenv.Run()
	scriptbts, _ := os.ReadFile(SWD + "script/script.py")
	script := string(scriptbts)
	script = strings.Replace(script, "$", content, 1)
	script = strings.Replace(script, ">", id, 1)
	f, _ := os.Create(SWD + id + "/script.py")
	f.WriteString(script)
	f.Close()
	cmdvenv.Wait()
	cmdpip := exec.Command(SWD+id+"/venv/Scripts/pip.exe", "install", "requests")
	cmdpip.Env = os.Environ()
	cmdpip.Dir = SWD + id
	cmdpip.Run()
	return id
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
