package common

import (
	"fmt"
	"net/http"

	"allaboutapps.dev/aw/go-starter/internal/api"
	"github.com/labstack/echo/v4"
)

func GetSumRoute(s *api.Server) *echo.Route {
	//  CodeChanges: changed the path from /sum/:count to /sum and the SumHandler will check the count in the http query params
	return s.Router.Management.GET("/sum", getSumHandler(s))
}

// :count "1" =>   1      => "1"
// :count "2" =>   1+2    => "3"
// :count "3" =>   1+2+3  => "6"
func getSumHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {

		var cnt uint64
		//  CodeChanges: changed the Int64 buinding to MustUInt64 and all the int64 to uint64
		//  CodeChanges: changed the PathParamsBinder to QueryParamsBinder to check the count in the http query params
		if err := echo.QueryParamsBinder(c).MustUint64("count", &cnt).BindError(); err != nil {
			return c.String(http.StatusBadRequest, "Please provide an integer.\n")
		}
		//  CodeChanges: removed the cnt == 0 check, since it is not needed anymore with the Gaussian Sum
		//  CodeChanges: removed the for and replaced it with the Gaussian Sum
		sum := cnt * (cnt + 1) / 2
		return c.String(http.StatusOK, fmt.Sprint(sum, "\n"))
	}
}
