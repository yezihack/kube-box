package internal

import (
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func TestCheck_IPNetPool(t *testing.T) {
	Convey("TestCheck_IPNetPool", t, func() {
		Convey("ipnet", func() {
			ck := Check{
				FilePath:        "../data/ip.data",
				Timeout:         time.Second * 3,
				NetworkCardName: "WLAN",
			}
			result, err := ck.IPNetPool()
			So(err, ShouldBeNil)
			fmt.Println("len:", len(result.List))
			for _, info := range result.List {
				fmt.Println(info)
			}
		})
	})
}

func TestCheck_IPHealthzPool(t *testing.T) {
	Convey("Check_IPHealthzPool", t, func() {
		Convey("IPHealthz", func() {
			ck := Check{
				FilePath:        "../data/ip.data",
				Timeout:         time.Second * 3,
				NetworkCardName: "WLAN",
			}
			result, err := ck.IPHealthzPool(80, "healthz")
			So(err, ShouldBeNil)
			for _, info := range result.List {
				fmt.Println(info)
			}
		})
	})
}
