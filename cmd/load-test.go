// MIT License
//
// Copyright (c) 2019 Top Free Games
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/topfreegames/eventsgateway/loadtest"
	logruswrapper "github.com/topfreegames/eventsgateway/logger/logrus"
)

// loadTest represents the testclient command
var loadTest = &cobra.Command{
	Use:   "load-test",
	Short: "runs a load test client",
	Long:  `runs a load test client`,
	Run: func(cmd *cobra.Command, args []string) {
		log := logrus.New()
		if debug {
			log.SetLevel(logrus.DebugLevel)
		}
		if json {
			log.Formatter = new(logrus.JSONFormatter)
		}
		tc, err := loadtest.NewLoadTest(logruswrapper.NewWithLogger(log), config)
		if err != nil {
			log.Panic(err)
		}
		tc.Run()
	},
}

func init() {
	RootCmd.AddCommand(loadTest)
}
