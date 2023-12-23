package token

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

var (
	_jwt *Token
)

func TestMain(m *testing.M) {
	_jwt = &Token{
		SecretKey: "",
		Timeout:   time.Second * 3600,
	}

	m.Run()
	os.Exit(0)
}

func TestToken(t *testing.T) {
	convey.Convey("SignToken", t, func(ctx convey.C) {
		token := _jwt.Sign("oid", "unionid")
		ctx.So(token, convey.ShouldNotBeEmpty)
		fmt.Printf("token: %s\n", token)

		fmt.Println("Wait 2 minutes")
		time.Sleep(time.Minute * 2)

		valid := _jwt.Verify(token)
		ctx.So(valid, convey.ShouldBeTrue)
		fmt.Println("valid token check pass")

		c := _jwt.GetClaimWithoutVerify(token)
		ctx.So(c, convey.ShouldNotBeNil)
		fmt.Println("token GetClaimWithoutVerify pass")
		fmt.Println(c)

		c, v := _jwt.GetClaimVerify(token)
		ctx.So(c, convey.ShouldNotBeNil)
		ctx.So(v, convey.ShouldBeTrue)
		fmt.Println("token GetClaimVerify pass")
		fmt.Println(c)

		token = ""
		valid = _jwt.Verify(token)
		ctx.So(valid, convey.ShouldBeFalse)
		fmt.Println("invalid token check pass")
	})
}
