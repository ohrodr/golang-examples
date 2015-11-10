package main

import "bytes"
import "flag"
import "fmt"
import "io/ioutil"

func main() {
  var out_buffer bytes.Buffer
  flag.Parse()
  //for in_buffer.Scan() {
  //  out_buffer.WriteString(in_buffer.Text())
  //}
  for _, file_arg := range flag.Args() {
    dat, err := ioutil.ReadFile(file_arg)
    if err != nil {
      panic(err)
    }
    out_buffer.WriteString(string(dat))
  }
  s := out_buffer.String()
  fmt.Println(s)
}