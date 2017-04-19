package main

import (
	"fmt"

	"net/http"

	"os"

	"github.com/go-martini/martini"

	"io/ioutil"
)

func main() {

	fmt.Println("SERVER IS RUNNING....")

	m := martini.Classic()

	m.Get("/get", func(res http.ResponseWriter, r *http.Request) {

		//Директория с файлом
		//Для продакшн сервера
		path := "~/home/kek/MKS_AR_server/src/STEMserver/schedule.xml"
		//Для теста
		//path := "C:/gopath/src/STEMserver/schedule.xml"

		reqFile, err := os.Open(path)
		if err != nil {
			res.WriteHeader(1488)
			res.Write([]byte(err.Error()))
		} else {
			//Считываем файл в массив байт
			fi, err := reqFile.Stat()
			if err != nil {
				res.WriteHeader(1488)
				res.Write([]byte(err.Error()))
			} else {
				size := fi.Size()
				var bytes = make([]byte, size)
				reqFile.Read(bytes)

				//Отправляем
				res.WriteHeader(200)
				res.Write(bytes)
			}
			defer reqFile.Close()
		}
	})

	m.Post("/upload", func(res http.ResponseWriter, r *http.Request) {
		file, err := os.Create("schedule.xml")
		if err != nil {
			res.WriteHeader(1488)
			res.Write([]byte(err.Error()))
			defer file.Close()
		} else {
			defer file.Close()
			tmp, err := ioutil.ReadAll(r.Body)
			if err != nil {
				res.WriteHeader(1488)
				res.Write([]byte(err.Error()))
			} else {
				file.Write(tmp)
				defer file.Close()
				res.WriteHeader(200)
			}
		}
	})

	m.Run()
}
