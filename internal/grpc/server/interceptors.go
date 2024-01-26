package productsAPI

//func authorizationInterceptor(logger *slog.Logger, userService userService.User) grpc.UnaryServerInterceptor {
//	return func(ctx context.Context,
//		req interface{},
//		info *grpc.UnaryServerInfo,
//		handler grpc.UnaryHandler) (interface{}, error) {
//		const op = "grpcServer.authorizationInterceptor"
//
//		md, ok := metadata.FromIncomingContext(ctx)
//		if !ok {
//			return nil, fmt.Errorf("%s: %w", op, appError.InvalidMetadata)
//		}
//		authHeader := md.Get("authorization")
//
//		if len(authHeader) == 0 {
//			return nil, fmt.Errorf("%s: %w", op, appError.Unauthorized)
//		}
//
//		_, err :=
//
//		return handler(ctx, req)
//
//	}
//}
