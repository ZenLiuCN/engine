package engine

import "log/slog"

func init() {
	RegisterMod(NewConsole(slog.Default()))
	RegisterMod(Require{})
	RegisterMod(TextEncoders{})

	RegisterModule(GoModule{})
	RegisterModule(BufferModule{})
	RegisterModule(Os{})

	RegisterModule(IoModule{})
	RegisterModule(ContextModule{})
	RegisterModule(EngineModule{})
	RegisterModule(CryptoModule{})
	RegisterModule(EsBuild{})
	RegisterModule(HashModule{})
	RegisterModule(CodecModule{})
	RegisterModule(Compiler{})
	RegisterModule(TimeModule{})
	RegisterModule(HttpModule{})
	RegisterModule(BigModule{})

}
