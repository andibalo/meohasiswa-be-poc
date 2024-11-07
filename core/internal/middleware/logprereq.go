package middleware

//func LogPreReq(logger *zap.Logger) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//
//		// client app id
//		var clientID = ctx.Request.Header.Get(httpclient.XClientID)
//		payload, _ := io.ReadAll(ctx.Request.Body)
//
//		//traceID, spanID := observ.ReadTraceID(ctx.Request.Context())
//
//		compactPayload := &bytes.Buffer{}
//		err := json.Compact(compactPayload, payload)
//		if err != nil {
//			compactPayload = bytes.NewBuffer(payload)
//		}
//		// set client id
//		ctx.Set("x-client-id", clientID)
//		//ctx.Set("trace.id", traceID)
//		ctx.Set("path", ctx.Request.URL.Path)
//		ctx.Set("method", ctx.Request.Method)
//
//		// payload for log
//		logger.InfoWithContext(ctx, "Interceptor Log",
//			cilog.Field("payload", compactPayload),
//			cilog.Field("trace.id", traceID),
//			cilog.Field("span.id", spanID),
//		)
//
//		// payload for otel
//		ctx.Set("payload", string(payload))
//
//		// set req body again
//		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(payload))
//
//		ctx.Next()
//	}
//}
