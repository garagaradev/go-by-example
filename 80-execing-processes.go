/*package main
import (
  "os"
  "os/exec"
  "syscall"
)

func main(){
//  binary, lookErr := exec.LookPath("ls") on unix system
  binary, lookErr := exec.LookPath("cmd") //on windows

  if lookErr != nil {
    panic(lookErr)
  }

//  args := []string{"ls","-a","-l","-h"} on unix system
  args := []string{"cmd","/C","dir","/a","/q","/s"} //on windows


  env := os.Environ()

  execErr := syscall.Exec(binary, args, env)
  if execErr != nil {
    panic(execErr)
  }
}
*/

package main

import (
    "os"
    "os/exec"
    "log"
)

func main() {
    // Menggunakan cmd.exe untuk menjalankan perintah
    cmd := exec.Command("cmd", "/C", "dir", "/a", "/q", "/s")

    // Menyiapkan output agar bisa langsung ditampilkan
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Menjalankan perintah
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Command failed with error: %v", err)
    }
}

