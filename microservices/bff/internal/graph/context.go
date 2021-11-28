package graph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hizzuu/grpc-example-bff/internal/graph/model"
)

func getClaimsFromCtx(ctx context.Context) (*model.JwtClaims, error) {
	b, err := json.Marshal(ctx.Value(model.CtxClaimsKey))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal")
	}

	var c *model.JwtClaims
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal")
	}

	return c, nil
}
