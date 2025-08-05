package test

import (
	"errors"
	"rasp-cloud/conf"
	"rasp-cloud/models"
	"rasp-cloud/tests/inits"
	"rasp-cloud/tests/start"
	"reflect"
	"testing"

	"github.com/astaxie/beego/context"
	"github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIastConnection(t *testing.T) {
	Convey("Subject: Test iast conn Api\n", t, func() {
		monkey.PatchInstanceMethod(reflect.TypeOf(&context.BeegoInput{}), "Header",
			func(input *context.BeegoInput, key string) string {
				return start.TestApp.Id
			},
		)
		defer monkey.UnpatchInstanceMethod(reflect.TypeOf(&context.BeegoInput{}), "Header")

		Convey("when the param is valid", func() {
			rasp := start.TestRasp
			rasp.Environ["JAVA_HOME"] = "/home/java/jdk-7.0.25"
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("when the mongodb has errors", func() {
			monkey.Patch(models.UpsertRaspById,
				func(id string, rasp *models.Rasp) error {
					return errors.New("")
				},
			)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(start.TestRasp))
			So(r.Status, ShouldBeGreaterThan, 0)
			monkey.Unpatch(models.UpsertRaspById)
		})

		Convey("when the rasp_id is empty", func() {
			rasp := *start.TestRasp
			rasp.Id = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of rasp_id is less than 16", func() {
			rasp := *start.TestRasp
			rasp.Id = "123456789"
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of version is greater than 50", func() {
			rasp := *start.TestRasp
			rasp.Version = inits.GetLongString(51)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of version is 0", func() {
			rasp := *start.TestRasp
			rasp.Version = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of hostname is greater than 1024", func() {
			rasp := *start.TestRasp
			rasp.HostName = inits.GetLongString(1025)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the hostname is empty", func() {
			rasp := *start.TestRasp
			rasp.HostName = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the language is empty", func() {
			rasp := *start.TestRasp
			rasp.Language = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of language is greater than 50", func() {
			rasp := *start.TestRasp
			rasp.Language = inits.GetLongString(51)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the language version is empty", func() {
			rasp := *start.TestRasp
			rasp.LanguageVersion = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of language version is greater than 50", func() {
			rasp := *start.TestRasp
			rasp.LanguageVersion = inits.GetLongString(51)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of server type is greater than 256", func() {
			rasp := *start.TestRasp
			rasp.ServerType = inits.GetLongString(257)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of server version is greater than 50", func() {
			rasp := *start.TestRasp
			rasp.ServerVersion = inits.GetLongString(51)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the register ip is invalid ip address", func() {
			rasp := *start.TestRasp
			rasp.RegisterIp = "123456.1223"
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the heartbeat interval is less than 0", func() {
			rasp := *start.TestRasp
			rasp.HeartbeatInterval = -10
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the environ is nil", func() {
			rasp := *start.TestRasp
			rasp.Environ = nil
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("when the register callback url is invalid", func() {
			conf.AppConfig.RegisterCallbackUrl = "xxxxx.xxxx.xxxx.xxxx"
			rasp := *start.TestRasp
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("when the server type is empty", func() {
			rasp := *start.TestRasp
			rasp.ServerType = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the server version is empty", func() {
			rasp := *start.TestRasp
			rasp.ServerVersion = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of description is greater than 1024", func() {
			rasp := *start.TestRasp
			rasp.Description = inits.GetLongString(1025)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the length of host type is greater than 256", func() {
			rasp := *start.TestRasp
			rasp.HostType = inits.GetLongString(257)
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the register ip is empty", func() {
			rasp := *start.TestRasp
			rasp.RegisterIp = ""
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("the total length of environment variable is greater than 100000", func() {
			rasp := *start.TestRasp
			rasp.Environ = map[string]string{
				inits.GetLongString(50001): inits.GetLongString(50001),
			}
			r := inits.GetResponse("POST", "/v1/agent/rasp", inits.GetJson(rasp))
			So(r.Status, ShouldEqual, 0)
		})
	})
}

func TestSearchRaspIast(t *testing.T) {
	Convey("Subject: Test Rasp Search Api\n", t, func() {
		Convey("when the param is valid", func() {
			r := inits.GetResponse("POST", "/v1/api/rasp/search", inits.GetJson(
				map[string]interface{}{
					"data":    start.TestRasp,
					"page":    1,
					"perpage": 1,
				},
			))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("when the data is nil", func() {
			r := inits.GetResponse("POST", "/v1/api/rasp/search", inits.GetJson(
				map[string]interface{}{
					"data":    nil,
					"page":    1,
					"perpage": 1,
				},
			))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the mongodb has errors", func() {
			monkey.Patch(models.FindRasp, func(*models.Rasp, int, int) (int, []*models.Rasp, error) {
				return 0, nil, errors.New("")
			})
			r := inits.GetResponse("POST", "/v1/api/rasp/search", inits.GetJson(
				map[string]interface{}{
					"data":    start.TestRasp,
					"page":    1,
					"perpage": 1,
				},
			))
			So(r.Status, ShouldBeGreaterThan, 0)
			monkey.Unpatch(models.FindRasp)
		})
	})
}

func TestDeleteRaspIast(t *testing.T) {
	Convey("Subject: Test Rasp Delete Api\n", t, func() {
		Convey("delete the rasp with id", func() {
			rasp := &models.Rasp{
				Id:                "123456789987654321654987654312",
				AppId:             start.TestApp.Id,
				Language:          "java",
				Version:           "1.0",
				HostName:          "ubuntu",
				RegisterIp:        "10.23.25.36",
				LanguageVersion:   "1.8",
				ServerType:        "tomcat",
				RaspHome:          "/home/work/tomcat8",
				PluginVersion:     "2019-03-10-10000",
				HeartbeatInterval: 180,
				LastHeartbeatTime: 1551781949000,
				RegisterTime:      1551781949000,
				Environ:           map[string]string{},
			}
			monkey.Patch(models.FindRasp, func(*models.Rasp, int, int) (int, []*models.Rasp, error) {
				return 1, nil, errors.New("")
			})
			r := inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"id":     rasp.Id,
				"app_id": rasp.AppId,
			}))
			monkey.Unpatch(models.FindRasp)
			So(r.Status, ShouldBeGreaterThan, 0)

			monkey.Patch(models.RemoveRaspById, func(id string) error {
				return errors.New("")
			})
			r = inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"id":     rasp.Id,
				"app_id": rasp.AppId,
			}))
			monkey.Unpatch(models.RemoveRaspById)
			So(r.Status, ShouldBeGreaterThan, 0)

			models.UpsertRaspById(rasp.Id, rasp)
			monkey.Patch(models.RemoveRaspById, func(id string) error {
				return nil
			})
			defer monkey.Unpatch(models.RemoveRaspById)
			r = inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"id":     rasp.Id,
				"app_id": rasp.AppId,
			}))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("delete the rasp with register_ip", func() {
			rasp := &models.Rasp{
				Id:                "123456789987654321654987654312",
				AppId:             start.TestApp.Id,
				Language:          "java",
				Version:           "1.0",
				HostName:          "ubuntu",
				RegisterIp:        "123.23.23.23",
				LanguageVersion:   "1.8",
				ServerType:        "tomcat",
				RaspHome:          "/home/work/tomcat8",
				PluginVersion:     "2019-03-10-10000",
				HeartbeatInterval: 180,
				LastHeartbeatTime: 1551781949000,
				RegisterTime:      1551781949000,
				Environ:           map[string]string{},
			}
			models.UpsertRaspById(rasp.Id, rasp)
			r := inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"register_ip": rasp.RegisterIp,
				"app_id":      rasp.AppId,
				"expire_time": 10,
			}))
			So(r.Status, ShouldEqual, 0)
		})

		Convey("when the rasp_id doesn't exist", func() {
			r := inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"id":     "123456789",
				"app_id": start.TestApp.Id,
			}))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the app_id is empty", func() {
			r := inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"register_ip": "123465789",
				"app_id":      "",
			}))
			So(r.Status, ShouldBeGreaterThan, 0)
		})

		Convey("when the param is invalid", func() {
			r := inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"app_id":      start.TestApp.Id,
				"register_ip": "173.2323",
			}))
			So(r.Status, ShouldBeGreaterThan, 0)

			r = inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"app_id":      start.TestApp.Id,
				"expire_time": -100,
			}))
			So(r.Status, ShouldBeGreaterThan, 0)

			monkey.Patch(models.RemoveRaspBySelector, func(selector map[string]interface{}, appId string) (int, error) {
				return 0, errors.New("")
			})
			r = inits.GetResponse("POST", "/v1/api/rasp/delete", inits.GetJson(map[string]interface{}{
				"app_id":      start.TestApp.Id,
				"register_ip": "173.23.0.0",
			}))
			monkey.Unpatch(models.RemoveRaspBySelector)
			So(r.Status, ShouldBeGreaterThan, 0)
		})
	})
}
