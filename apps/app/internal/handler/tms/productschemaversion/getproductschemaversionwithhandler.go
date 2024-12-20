package productschemaversion

import (
	"github.com/iscosmos/neurium/pkg/response"
	"github.com/iscosmos/neurium/pkg/validator"
	"net/http"

	"github.com/iscosmos/neurium/apps/app/internal/logic/tms/productschemaversion"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductSchemaVersionWithHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductSchemaVersionWithReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := validator.Validate(&req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := productschemaversion.NewGetProductSchemaVersionWithLogic(r.Context(), svcCtx)
		resp, err := l.GetProductSchemaVersionWith(&req)

		response.Response(w, resp, err)
	}
}
