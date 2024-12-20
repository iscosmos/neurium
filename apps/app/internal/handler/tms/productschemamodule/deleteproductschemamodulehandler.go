package productschemamodule

import (
	"github.com/iscosmos/neurium/pkg/response"
	"github.com/iscosmos/neurium/pkg/validator"
	"net/http"

	"github.com/iscosmos/neurium/apps/app/internal/logic/tms/productschemamodule"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductSchemaModuleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductSchemaModuleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := validator.Validate(&req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := productschemamodule.NewDeleteProductSchemaModuleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductSchemaModule(&req)

		response.Response(w, resp, err)
	}
}
