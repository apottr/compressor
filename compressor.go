package main

import (
  "io/ioutil"
  "log"
  "encoding/csv"
  "time"
  "strings"
  "strconv"
  "os"
)

func processFile (filename string) string {
  var name = strings.Split(filename,".")
  var day,_ = strconv.Atoi(strings.Split(name[0],"-")[1])
  var month,_ = strconv.Atoi(name[1])
  var year,_ = strconv.Atoi("20"+name[2])
  var hour,_ = strconv.Atoi(name[3])
  var minute,_ = strconv.Atoi(name[4])
  var t = time.Date(year,time.Month(month),day,hour,minute,0,0,time.Local)
  return t.Format(time.RFC822Z)
}

func pull_module(module string) {
  files, err := ioutil.ReadDir("/home/Datasets/panopticon_collector/"+module+"/raw")
  if err != nil {
    log.Fatal(err)
  }
  out := [][]string {{"filename","date"}}
  for _, file := range files {
    out = append(out,[]string {file.Name(), processFile(file.Name())})
  }
  f,err := os.OpenFile("/home/Datasets/panopticon_collector/"+module+"/files.csv",
                        os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
  w := csv.NewWriter(f)
  if err != nil {
    log.Fatal(err)
  }
  w.WriteAll(out)

}

func main(){
  for _,mod := range []string {"rss", "cameras", "scanner"} {
    pull_module(mod)
  }
}
