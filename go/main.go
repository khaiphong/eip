/* 
  Copyright (c) 2017, 2018, 2019, 2020, 2021 KhaiPhong

  Licensed under the Apache License, Version 2.0 (the "License"); you may not 
  use this file except in compliance with the License which can be copied at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  
  Mu EIP-Default() initiates EIP Free services to registered user: (1) Free 
  Chat, (2) Free Todos list, (3) Free accesses to PersonaDB, (4) Free access 
  to PersonaAI, (5) Free access to KP-Platform 
*/

package main

import (
  "fmt"
  "os"
)	
	
func main() {

//  var me [4]int me of 4 integers MyInstinct, MyMindfulness, MySamadhi, MyPrajna 

  switch os.Args[1] {
    case "run":
	run()
    default:
	panic("I am confused!")
    } 
}

// docker run <image> <cmd> <args>
// go     run main.go run <cmd> <args>

