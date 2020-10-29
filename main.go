/*
 * Copyright (C) 2020  Dmitry Minko
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 3.
 *
 * MEGA is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"./ui"
	"github.com/nanu-c/qml-go"
	"log"
)

func main() {
	err := qml.Run(run)
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ui.SetEngine()

	ui.InitModels()
	err := ui.SetComponent()
	if err != nil {
		return err
	}
	ui.Win.Show()
	ui.Win.Wait()

	return nil
}

func showProgress(ch chan int) {
	bytesread := 0
	for {
		b := 0
		ok := false
		select {
		case b, ok = <-ch:
			if ok == false {
				return
			}
		}
		bytesread += b
	}
}
