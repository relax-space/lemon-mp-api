package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/relax-space/go-kit/model"

	"github.com/labstack/echo"
	"github.com/relax-space/lemon-wxmp-sdk/mpAuth"
)

func OpenIdGate(c echo.Context) error {
	reqDto := new(ReqDto)
	if err := c.Bind(reqDto); err != nil {
		return c.JSON(http.StatusOK, model.Result{Error: model.Error{Message: err.Error()}})
	}
	reqUrl := mpAuth.GetUrlForAccessToken(reqDto.ReqDto)
	return c.Redirect(http.StatusFound, reqUrl)
}

func OpenIdRequest(c echo.Context) error {
	appId := c.Param("appId")
	reUrl := c.QueryParam("reurl")
	code := c.QueryParam("code")
	//get secret by appId
	secret := os.Getenv("WXMP_SECRET")
	respDto, err := mpAuth.GetAccessTokenAndOpenId(code, appId, secret)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Error: model.Error{Message: err.Error()}})
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("%v?openId=%", reUrl, respDto.OpenId))

}
