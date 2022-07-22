package kuevs

import "github.com/golang/glog"

func Check(err error) {
	if err != nil {
		glog.Errorln(err)
	}
}
