package depended

import "go.uber.org/zap"

func Run() {
  logger, _ := zap.NewProduction()
  sugar := logger.Sugar()
  sugar.Infow("In depended")
  defer logger.Sync()
}
