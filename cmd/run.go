/*
Copyright © 2025 Matt Krueger <mkrueger@rstms.net>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

 1. Redistributions of source code must retain the above copyright notice,
    this list of conditions and the following disclaimer.

 2. Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

 3. Neither the name of the copyright holder nor the names of its contributors
    may be used to endorse or promote products derived from this software
    without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package cmd

import (
	"fmt"
	"github.com/rstms/sysmenu/menu"
	"github.com/spf13/cobra"
	"log"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a process testing the tray menu",
	Long: `
test the tray menu
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		menu := menu.NewMenu("MenuTest", "a test program for the system menu", []byte{})
		/*
			foo := menu.AddItem("foo", "the foo item")
			menu.AddSeparator()
			bar := menu.AddItem("bar", "the bar item")
			bar.AddItem("bar-plough", "subitem under bar named plough")
			bar.AddItem("bar-plover", "subitem under bar named plover")
			menu.AddSeparator()
			bar.AddItem("bar-quux", "subitem under bar named quux")
			bar.AddItem("bar-quuux", "subitem under bar named quuux")
			menu.AddSeparator()
			baz := menu.AddItem("baz", "the goo item")

			log.Printf("foo=%v\n", foo)
			log.Printf("bar=%v\n", bar)
			log.Printf("baz=%v\n", baz)
		*/

		/*
			handlerExited := make(chan struct{}, 1)

			go func() {
				//defer log.Println("cmd.Run.handler exiting")
				//log.Println("cmd.Run.handler starting")
				for {
					select {
					case item := <-menu.Clicked:
						log.Printf("cmd.Run.handler.Clicked: %d %s\n", (*item).Id, (*item).Title)
					case <-menu.Exited:
						//log.Println("cmd.Run.handler.Exited; returning from handler")
						handlerExited <- struct{}{}
						return

					}
				}
			}()
		*/

		/*
			log.Println("cmd.Run - calling menu.Start()")
			err := menu.Start()
			cobra.CheckErr(err)
			log.Println("cmd.Run - after menu.Start()")

			<-handlerExited

			log.Println("cmd.Run - calling menu.Wait()")
			err = menu.Wait()
			cobra.CheckErr(err)
			log.Println("cmd.Run - after menu.Wait()")
		*/

		err := menu.Run(nil, nil)
		cobra.CheckErr(err)
		//<-handlerExited
		log.Println("menu.Run returned")

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
