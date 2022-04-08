package internal

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPingIPNet(t *testing.T) {
	Convey("127.0.0.1", t, func() {
		ok, stats, err := PingNetwork("127.0.0.1")
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)
		fmt.Printf("最大耗时:%s\n", stats.MaxRtt)
		fmt.Printf("%+v\n", stats)
	})
	Convey("www.baidu.com", t, func() {
		ok, stats, err := PingNetwork("www.baidu.com")
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)
		fmt.Printf("最大耗时:%d 毫秒\n", stats.MaxRtt.Milliseconds())
		fmt.Printf("%+v\n", stats)
	})
	Convey("www.google.com", t, func() {
		ok, stats, err := PingNetwork("www.google.com")
		So(err, ShouldBeNil)
		So(ok, ShouldBeFalse)
		fmt.Printf("最大耗时:%d 毫秒\n", stats.MaxRtt.Milliseconds())
		fmt.Printf("%+v\n", stats)
	})
}

// go test -v -run TestReadTextLine
func TestReadTextLine(t *testing.T) {
	Convey("read line by text", t, func() {
		Convey("a", func() {
			filePath := "../data/ip.data"
			result, err := ReadTextLine(filePath)
			So(err, ShouldBeNil)
			fmt.Println(result)
		})
		Convey("b", func() {
			filePath := "../data/ipb.data"
			result, err := ReadTextLine(filePath)
			So(err, ShouldNotBeNil)
			fmt.Println(result)
		})
	})
}

func TestGetEnvValue(t *testing.T) {
	Convey("GetEnvValue", t, func() {
		Convey("get default val", func() {
			val := GetEnvValue(EnvVersion, DefaultVersion)
			So(val, ShouldEqual, DefaultVersion)
		})
		Convey("set env", func() {
			os.Setenv("PORT", "8080")
			val := GetEnvValueToInteger(EnvPort, DefaultPort)
			So(val, ShouldEqual, 8080)
		})
	})
}

func TestCurlGet(t *testing.T) {
	Convey("Curl GET", t, func() {
		Convey("site", func() {
			uri := "https://sgfoot.com/curl.html"
			result, statusCode, err := CurlGet(uri, time.Second*2)
			So(err, ShouldBeNil)
			So(statusCode, ShouldEqual, 200)
			fmt.Println(string(result))
			So(len(result), ShouldBeGreaterThan, 100)
		})
	})
}

func TestIsExistPathLastPole(t *testing.T) {
	Convey("IsExistPathLastPole", t, func() {
		Convey("None", func() {
			So(IsExistPathLastPole(""), ShouldBeFalse)
		})
		Convey("last one", func() {
			So(IsExistPathLastPole("1"), ShouldBeFalse)
			So(IsExistPathLastPole("/"), ShouldBeTrue)
		})
		Convey("gt one", func() {
			So(IsExistPathLastPole("foot/"), ShouldBeTrue)
			So(IsExistPathLastPole("foo/aoo/"), ShouldBeTrue)
			So(IsExistPathLastPole("foo/aoo"), ShouldBeFalse)
		})
	})
}

func TestLocalIPv4s(t *testing.T) {
	Convey("Get local ip", t, func() {
		Convey("ip1", func() {
			fmt.Println(LocalIPv4s())
		})
		Convey("ip2", func() {
			fmt.Println(GetLocalIpV4())
		})
		Convey("ip3", func() {
			fmt.Println(GetInterfaceIpv4Addr("WLAN"))
		})
	})
}
