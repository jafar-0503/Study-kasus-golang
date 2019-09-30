package main

import "fmt"
import "encoding/json"
import "net/http"
import "bytes"
import "net/url"
var baseURL = "http://localhost:8080"

type materi_golang struct{
  ID int
  Nama string
  Jurusan string
  Alamat string
}

func ambil_api(mhs string)(materi_golang, error){
  var err error
  var client = &http.Client{}
  var data materi_golang

  var param = url.Values{}
  param.Set("Nama", mhs)
  var payload = bytes.NewBufferString(param.Encode())
  request, err := http.NewRequest("POST", baseURL + "/cari_mhs", payload)

  if err != nil{
    return data, err
  }
  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)
  if err != nil{
    return data, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil{
    return data, err
  }
  return data, nil
}

func main(){
  var mhs, err = ambil_api("Andika Lesmana")
  if err != nil{
    fmt.Println("Data Tidak Tersedia!", err.Error())
    return
  }


    fmt.Println("ID : ", mhs.ID, " Nama : ", mhs.Nama, " Jurusan : ", mhs.Jurusan, " Alamat : ", mhs.Alamat,)

}
